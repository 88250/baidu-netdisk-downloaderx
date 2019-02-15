package pcscommand

import (
	"container/list"
	"fmt"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/andlabs/ui"
	"github.com/iikira/BaiduPCS-Go/baidupcs"
	"github.com/iikira/BaiduPCS-Go/internal0/pcsconfig"
	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
	"github.com/iikira/BaiduPCS-Go/requester"
	"github.com/iikira/BaiduPCS-Go/requester/downloader"
	"github.com/iikira/BaiduPCS-Go/requester/rio"
)

var ProgressFilenamePlaceholder = labelPlaceholder(56)

func labelPlaceholder(n int) string {
	builder := strings.Builder{}

	for i := 0; i < n; i++ {
		builder.WriteString(" ")
	}

	return builder.String()
}

const (
	//DownloadSuffix 文件下载后缀
	DownloadSuffix = ".加速下载中"
	//StrDownloadInitError 初始化下载发生错误
	StrDownloadInitError = "初始化下载发生错误"
)

// dtask 下载任务
type dtask struct {
	ListTask
	path         string                  // 下载的路径
	savePath     string                  // 保存的路径
	downloadInfo *baidupcs.FileDirectory // 文件或目录详情
}

//DownloadOption 下载可选参数
type DownloadOption struct {
	IsTest               bool
	IsPrintStatus        bool
	IsExecutedPermission bool
	IsOverwrite          bool
	SaveTo               string
	Parallel             int
}

func getDownloadFunc(id int, path, savePath string, cfg *downloader.Config, isPrintStatus, isExecutedPermission bool, uiProgress *UIProgress) baidupcs.DownloadFunc {
	if cfg == nil {
		cfg = downloader.NewConfig()
	}

	return func(downloadURL string, jar *cookiejar.Jar) (speed int64, err error) {
		h := requester.NewHTTPClient()
		h.SetCookiejar(jar)
		h.SetKeepAlive(true)
		h.SetTimeout(10 * time.Minute)
		setupHTTPClient(h)

		var (
			file          *os.File
			writeCloserAt rio.WriteCloserAt
			exitChan      chan struct{}
		)

		if !cfg.IsTest {
			cfg.InstanceStatePath = savePath + DownloadSuffix

			// 创建下载的目录
			dir := filepath.Dir(savePath)
			var fileInfo os.FileInfo
			fileInfo, err = os.Stat(dir)
			if err != nil {
				err = os.MkdirAll(dir, 0777)
				if err != nil {
					return
				}
			} else if !fileInfo.IsDir() {
				err = fmt.Errorf("%s, path %s: not a directory", StrDownloadInitError, dir)
				return
			}

			file, err = os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0666)
			if file != nil {
				defer file.Close()
			}
			if err != nil {
				err = fmt.Errorf("%s, %s", StrDownloadInitError, err)
				return
			}

			// 空指针和空接口不等价
			if file != nil {
				writeCloserAt = file
			}
		}

		download := downloader.NewDownloader(downloadURL, writeCloserAt, cfg)
		download.SetClient(h)

		exitChan = make(chan struct{})
		download.OnExecute(func() {
			if cfg.IsTest {
				fmt.Printf("[%d] 测试下载开始\n\n", id)
			}

			var (
				ds                            = download.GetDownloadStatusChan()
				downloaded, totalSize, speeds int64
				leftStr                       string
			)
			for {
				select {
				case <-exitChan:
					return
				case v, ok := <-ds:
					if !ok { // channel 已经关闭
						return
					}

					downloaded, totalSize, speeds = v.Downloaded(), v.TotalSize(), v.SpeedsPerSecond()
					if speeds <= 0 {
						leftStr = "-"
					} else {
						leftStr = (time.Duration((totalSize-downloaded)/(speeds)) * time.Second).String()
					}

					percentage := float64(v.Downloaded()) / float64(v.TotalSize()) * 100
					downloaded := converter.ConvertFileSize(v.Downloaded(), 2)
					total := converter.ConvertFileSize(v.TotalSize(), 2)
					speedLable := converter.ConvertFileSize(v.SpeedsPerSecond(), 2)
					ui.QueueMain(func() {
						defer recover()
						uiProgress.PB.SetValue(int(percentage))
						uiProgress.PL.SetText(downloaded + "/" + total + " 速度 " + speedLable + " 估计剩余 " + leftStr)
					})

					totalSize = v.TotalSize()

					s := int64(float64(totalSize) / v.TimeElapsed().Seconds())
					speed = s
				}
			}
		})

		if isPrintStatus {
			go func() {
				for {
					time.Sleep(1 * time.Second)
					select {
					case <-exitChan:
						return
					default:
						download.PrintAllWorkers()
					}
				}
			}()
		}
		err = download.Execute()
		close(exitChan)
		if err != nil {
			return
		}

		if isExecutedPermission {
			file.Chmod(0766)
			//if err != nil {
			//	fmt.Printf("\n\n[%d] 警告, 加执行权限错误: %s\n\n", id, err)
			//}
		}

		if !cfg.IsTest {
			//fmt.Printf("\n\n[%d] 下载完成, 保存位置: %s\n\n", id, savePath)
			ui.QueueMain(func() {
				uiProgress.PF.SetText(ShortLabel("下载完成 " + path))
			})
		} else {
			fmt.Printf("\n\n[%d] 测试下载结束\n\n", id)
		}

		return
	}
}

// RunDownload 执行下载网盘内文件
func RunDownload(paths []string, option DownloadOption, uiProgress *UIProgress) {
	// 设置下载配置
	cfg := &downloader.Config{
		IsTest:    option.IsTest,
		CacheSize: pcsconfig.Config.CacheSize(),
	}

	// 设置下载最大并发量
	if option.Parallel == 0 {
		option.Parallel = pcsconfig.Config.MaxParallel()
	}
	cfg.MaxParallel = option.Parallel

	paths, err := getAllAbsPaths(paths...)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("\n")
	//fmt.Printf("[0] 提示: 当前下载最大并发量为: %d, 下载缓存为: %d\n", cfg.MaxParallel, cfg.CacheSize)

	var (
		pcs    = GetBaiduPCS()
		dlist  = list.New()
		lastID = 0
	)

	for k := range paths {
		lastID++
		ptask := &dtask{
			ListTask: ListTask{
				ID:       lastID,
				MaxRetry: 3,
			},
			path: paths[k],
		}
		if option.SaveTo != "" {
			ptask.savePath = filepath.Join(option.SaveTo, filepath.Base(paths[k]))
		} else {
			ptask.savePath = GetActiveUser().GetSavePath(paths[k])
		}
		dlist.PushBack(ptask)
		//fmt.Printf("[%d] 加入下载队列: %s\n", lastID, paths[k])
	}

	var (
		handleTaskErr = func(task *dtask, errManifest string, err error) {
			if task == nil {
				panic("task is nil")
			}

			if err == nil {
				return
			}

			// 不重试的情况
			switch {
			case strings.Compare(errManifest, "下载文件错误") == 0 && strings.Contains(err.Error(), StrDownloadInitError):
				fmt.Printf("[%d] %s, %s\n", task.ID, errManifest, err)
				return
			}

			//fmt.Printf("[%d] %s, %s, 重试 %d/%d\n", task.ID, errManifest, err, task.retry, task.MaxRetry)

			// 未达到失败重试最大次数, 将任务推送到队列末尾
			if task.retry < task.MaxRetry {
				task.retry++
				dlist.PushBack(task)
			}
			time.Sleep(3 * time.Duration(task.retry) * time.Second)
		}
		totalSize int64
	)

	fileCnt := 0
	for {
		e := dlist.Front()
		if e == nil { // 结束
			break
		}

		dlist.Remove(e) // 载入任务后, 移除队列

		task := e.Value.(*dtask)
		if task == nil {
			continue
		}

		if task.downloadInfo == nil {
			task.downloadInfo, err = pcs.FilesDirectoriesMeta(task.path)
			if err != nil {
				// 不重试
				fmt.Printf("[%d] 获取路径信息错误, %s\n", task.ID, err)
				continue
			}
		}

		//fmt.Printf("\n")
		//fmt.Printf("[%d] ----\n%s\n", task.ID, task.downloadInfo.String())

		// 如果是一个目录, 将子文件和子目录加入队列
		if task.downloadInfo.Isdir {
			if !option.IsTest { // 测试下载, 不建立空目录
				os.MkdirAll(task.savePath, 0777) // 首先在本地创建目录, 保证空目录也能被保存
			}

			fileList, err := pcs.FilesDirectoriesList(task.path)
			if err != nil {
				// 不重试
				fmt.Printf("[%d] 获取目录信息错误, %s\n", task.ID, err)
				continue
			}

			for k := range fileList {
				lastID++
				subTask := &dtask{
					ListTask: ListTask{
						ID:       lastID,
						MaxRetry: 3,
					},
					path:         fileList[k].Path,
					downloadInfo: fileList[k],
				}

				if option.SaveTo != "" {
					subTask.savePath = filepath.Join(task.savePath, fileList[k].Filename)
				} else {
					subTask.savePath = GetActiveUser().GetSavePath(subTask.path)
				}

				dlist.PushBack(subTask)
				//fmt.Printf("[%d] 加入下载队列: %s\n", lastID, fileList[k].Path)
			}
			continue
		}

		//fmt.Printf("[%d] 准备下载: %s\n", task.ID, task.path)

		useAria2 := pcsconfig.IsWindows() && task.downloadInfo.Size > 1024*1024*512

		if !option.IsTest && !option.IsOverwrite && fileExist(task.savePath) && !useAria2 {
			//fmt.Printf("[%d] 文件已经存在: %s, 跳过...\n", task.ID, task.savePath)
			ui.QueueMain(func() {
				uiProgress.PF.SetText(ShortLabel("文件已存在 " + task.path))
			})

			continue
		}

		//if !option.IsTest {
		//	fmt.Printf("[%d] 将会下载到路径: %s\n\n", task.ID, task.savePath)
		//}

		ui.QueueMain(func() {
			uiProgress.PF.SetText(ShortLabel(task.path))
		})

		var downloadFunc baidupcs.DownloadFunc
		if useAria2 {
			parts := int(task.downloadInfo.Size / 1024 / 1024 / 2)
			if parts > 512 {
				parts = 512
			}
			if 32 > parts {
				parts = 32
			}
			s := strconv.Itoa(parts)

			downloadFunc = GetAria2DownloadFunc(uiProgress, task.path, task.savePath, s)
		} else {
			downloadFunc = getDownloadFunc(task.ID, task.path, task.savePath, cfg, option.IsPrintStatus, option.IsExecutedPermission, uiProgress)
		}

		speed, err := pcs.DownloadFile(task.path, downloadFunc)
		if err != nil {
			handleTaskErr(task, "下载文件错误", err)
			continue
		}

		totalSize += task.downloadInfo.Size

		//consumeG := int(math.Ceil(float64(task.downloadInfo.Size) / 1024 / 1024 / 1024))
		fileCnt++
		downloadSize := strconv.FormatInt(task.downloadInfo.Size, 10)
		go Stat(&StatData{Size: downloadSize, Speed: strconv.FormatInt(speed, 10)})
	}

	//fmt.Printf("任务结束, 数据总量: %s\n", converter.ConvertFileSize(totalSize))
	ui.QueueMain(func() {
		uiProgress.PB.SetValue(0)
		uiProgress.PF.SetText(ShortLabel("下载完成 " + paths[0]))
		uiProgress.PL.SetText(ShortLabel("本地目录 " + pcsconfig.Config.GetSaveDir()))
		if 0 < totalSize {
			ui.MsgBox(uiProgress.Win, "下载完成", "本次共下载 "+strconv.Itoa(fileCnt)+" 个文件，数据总量 "+
				converter.ConvertFileSize(totalSize, 2))
		}
	})
	Downloading = false
}

// fileExist 检查文件是否存在,
// 只有当文件存在, 断点续传文件不存在时, 才判断为存在
func fileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		if _, err = os.Stat(path + DownloadSuffix); err != nil {
			return true
		}
	}

	return false
}
