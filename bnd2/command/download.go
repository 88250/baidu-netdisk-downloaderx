package command

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/88250/bnd2/util"
	"github.com/dustin/go-humanize"
)

type downloaddir struct {
}

type downloadfile struct {
}

type dtask struct {
	Path      string  `json:"path"`
	Name      string  `json:"name"`
	PCSURL    string  `json:"pcsURL"`
	SaveDir   string  `json:"saveDir"`
	SavePath  string  `json:"savePath"`
	State     int     `json:"state"`
	HState    string  `json:"hState"`
	Speed     string  `json:"speed"`
	Size      uint64  `json:"size"`
	HSize     string  `json:"hSize"`
	CSize     uint64  `json:"csize"`
	HCSize    string  `json:"hCSize"`
	Progress  float64 `json:"progress"`
	ETA       string  `json:"eta"`
	Pieces    string  `json:"pieces"`
	Conns     string  `json:"conns"`
	Created   int64   `json:"created"`
	Completed int64   `json:"completed"`
	Gid       string  `json:"gid"`
}

const (
	stateDownloading = iota
	stateCompleted
	stateFailed
	statePaused
)

func (cmd *downloadfile) Exec(param map[string]interface{}) {
	path := param["path"].(string)
	saveDir := param["saveDir"].(string)
	util.User.SaveDir = saveDir
	size := uint64(param["size"].(float64))

	if !containsTask(path) {
		task := &dtask{
			Path:     path,
			Name:     filepath.Base(path),
			PCSURL:   util.DowanloadURL(path),
			SaveDir:  saveDir,
			SavePath: filepath.Base(path),
			State:    statePaused,
			HState:   "暂停",
			Size:     size,
			HSize:    humanize.Bytes(size),
			Speed:    "0 B",
			ETA:      "--",
			Pieces:   "0",
			Conns:    "0",
			Created:  time.Now().UnixNano(),
		}

		task.doDownload()
	}
}

func (cmd *downloaddir) Exec(param map[string]interface{}) {
	path := param["path"].(string)
	saveDir := param["saveDir"].(string)
	util.User.SaveDir = saveDir

	// 拷贝出来，因为遍历命令随时可能会被重置
	var files []util.File
	for i := 0; i < len(traverseCmd.files); i++ {
		files = append(files, *traverseCmd.files[i])
	}

	for i := 0; i < len(files); i++ {
		travFile := files[i]

		if containsTask(travFile.Path) {
			continue
		}

		fullPath := travFile.Path
		dirPath := filepath.ToSlash(filepath.Dir(path))
		savePath := strings.Replace(fullPath, dirPath, "", 1)

		task := &dtask{
			Path:     fullPath,
			Name:     filepath.Base(travFile.Name),
			PCSURL:   util.DowanloadURL(travFile.Path),
			SaveDir:  saveDir,
			SavePath: savePath,
			State:    statePaused,
			HState:   "暂停",
			Size:     travFile.Size,
			HSize:    humanize.Bytes(travFile.Size),
			Speed:    "0 B",
			ETA:      "--",
			Pieces:   "0",
			Conns:    "0",
			Created:  time.Now().UnixNano(),
		}

		task.doDownload()
	}
}

func (cmd *downloadfile) Name() string {
	return "downloadfile"
}

func (cmd *downloaddir) Name() string {
	return "downloaddir"
}

func (task *dtask) pause() {
	g, err := util.R.Pause(task.Gid)
	if nil != err {
		logger.Errorf("pause task [path=%s] failed [%s]", task.Path, err)

		return
	}
	task.HState = "暂停"
	task.State = statePaused
	task.Speed = "0 B"

	logger.Infof("paused task [path=%s, g=%s]", task.Path, g)
}

func (task *dtask) unpause() {
	g, err := util.R.Unpause(task.Gid)
	if nil != err {
		msg := fmt.Sprintf("unpause task [path=%s, g=%s] failed [%s]", task.Path, g, err)
		if !strings.Contains(msg, "cannot be unpaused now") {
			logger.Errorf(msg)
		}

		return
	}
	task.HState = "下载中"
	task.State = stateDownloading
	logger.Infof("unpaused task [path=%s, g=%s]", task.Path, g)
}

func (task *dtask) doDownload() {
	time.Sleep(5 * time.Millisecond)

	options := map[string]interface{}{"dir": task.SaveDir, "out": task.SavePath}
	downloadingCount := countDownloadingTask()
	if 5 <= downloadingCount {
		options["pause"] = true
		task.HState = "暂停"
		task.State = statePaused
	} else {
		task.HState = "下载中"
		task.State = stateDownloading
	}

	gid, err := util.R.AddURI(task.PCSURL, options)
	if nil != err {
		logger.Errorf("add task [path=%s] failed [%s]", task.Path, err)

		return
	}
	task.Gid = gid
	addTask(task)

	go func() {
		for {
			time.Sleep(time.Second)

			info, err := util.R.TellStatus(gid)
			if nil != err {
				//logger.Errorf("query task [path=%s] state failed [%s]", task.Path, err)

				break
			}

			cl, _ := strconv.ParseInt(info.CompletedLength, 10, 64)
			tl, _ := strconv.ParseInt(info.TotalLength, 10, 64)
			if 1 > tl {
				continue
			}

			task.CSize = uint64(cl)
			task.HCSize = humanize.Bytes(task.CSize)
			s, _ := strconv.ParseInt(info.DownloadSpeed, 10, 64)
			ss := humanize.Bytes(uint64(s))
			eta := float64(tl-cl) / float64(s)
			leftStr := (time.Duration(eta) * time.Second).String()
			p := float64(cl) / float64(tl) * 100
			task.Progress = p
			task.Speed = ss
			task.ETA = leftStr
			task.Pieces = info.NumPieces
			task.Conns = info.Connections

			if "removed" == info.Status {
				break
			}

			if "error" == info.Status {
				task.HState = "失败"
				task.State = stateFailed

				break
			}

			if "complete" == info.Status {
				task.HState = "完成"
				task.State = stateCompleted
				task.Completed = time.Now().UnixNano()
				deleteTask(task.Gid)
				addCTask(task)

				elapsed := (task.Completed - task.Created) / int64(time.Second)
				averageSpeed := task.Size / uint64(elapsed)

				go util.Stat(&util.StatData{Size: fmt.Sprintf("%d", task.Size), Speed: fmt.Sprintf("%d", averageSpeed)})

				break
			}

			if "active" == info.Status {
				task.HState = "下载中"
				task.State = stateDownloading
			}

			if "paused" == info.Status {
				result := util.NewCmdResult(lstasksCmd.Name())
				result.Data = tasks
				util.Push(result.Bytes())

				continue
			}

			logger.Tracef(task.Path + " " + task.HCSize + "/" + task.HSize + " " + strconv.FormatFloat(p, 'f', 2, 64) + "% " + ss + "/s ETA " + leftStr +
				" " + task.Pieces + "p " + task.Conns + "c")

			result := util.NewCmdResult(lstasksCmd.Name())
			result.Data = tasks
			util.Push(result.Bytes())
		}

		result := util.NewCmdResult(lstasksCmd.Name())
		result.Data = tasks
		util.Push(result.Bytes())
	}()
}
