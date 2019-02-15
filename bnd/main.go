package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/andlabs/ui"
	"github.com/iikira/BaiduPCS-Go/internal0/pcscommand"
	"github.com/iikira/BaiduPCS-Go/internal0/pcsconfig"
	"github.com/iikira/BaiduPCS-Go/pcscache"
	"github.com/iikira/BaiduPCS-Go/requester"
	"github.com/pkg/errors"
)

const (
	Width  = 800
	Height = 150
)

var selectedFile = ""
var activeBaiduUser *pcsconfig.Baidu

func init() {
	pcscommand.HomeDir = pcsconfig.UserHome()
	pcscommand.BndDir = filepath.Join(pcscommand.HomeDir, ".bnd")
	os.MkdirAll(pcscommand.BndDir, os.ModePerm)

	go unpackAria2()

	needUpgrade, tip, newVerURL := pcscommand.CheckUpgrade()
	if needUpgrade {
		ui.Main(func() {
			win := ui.NewWindow(pcscommand.WindowTitle, 400, 40, false)
			win.SetMargined(true)
			label := ui.NewLabel(tip)
			entry := ui.NewEntry()
			entry.SetText(newVerURL)
			entry.SetReadOnly(true)
			box := ui.NewVerticalBox()
			box.SetPadded(true)
			box.Append(label, true)
			box.Append(entry, true)
			win.SetChild(box)
			win.OnClosing(onClose)
			win.Show()
		})

		os.Exit(1)
	}

	pcsconfig.Config.Init()
	pcsconfig.Config.SetAppID(pcscommand.AppID)
	pcsconfig.Config.SetUserAgent("BND/" + pcscommand.Ver)

	pcscommand.SaveDir = filepath.Join(pcscommand.HomeDir, "Downloads")
	os.MkdirAll(pcscommand.SaveDir, os.ModePerm)
	pcsconfig.Config.SetSaveDir(pcscommand.SaveDir)
	pcsconfig.Config.Save()

	pcscache.DirCache.GC()
	requester.TCPAddrCache.GC()
}

func main() {
	activeBaiduUser = pcsconfig.Config.ActiveUser()
	if 0 < activeBaiduUser.UID {
		_, err := pcsconfig.Config.SetupUserByBDUSS(activeBaiduUser.BDUSS, "", "")
		if nil == err {
			pcscommand.BDUSS = activeBaiduUser.BDUSS
		}
	}

	if err := ui.Main(func() { showBDLogin() }); nil != err {
		panic(err)
	}
}

func showBDLogin() {
	loginWin := ui.NewWindow(pcscommand.WindowTitle, Width, Height, false)
	vBox := ui.NewVerticalBox()
	vBox.SetPadded(true)

	bdussEntry := ui.NewEntry()
	if "" != pcscommand.BDUSS {
		bdussEntry.SetText(pcscommand.BDUSS)
	} else {
		bdussEntry.SetText("")
	}
	usageButton := ui.NewButton("点我查看使用指南")
	usageButton.OnClicked(func(button *ui.Button) {
		OpenBrowser(loginWin, "https://hacpai.com/article/1524460877352")
	})
	vBox.Append(ui.NewHorizontalSeparator(), false)
	vBox.Append(usageButton, true)
	vBox.Append(ui.NewHorizontalSeparator(), false)
	bdussBox := ui.NewHorizontalBox()
	bdussBox.Append(ui.NewLabel("BDUSS："), false)
	bdussBox.Append(bdussEntry, true)
	vBox.Append(bdussBox, true)
	loginButton := ui.NewButton("登录")
	loginButton.OnClicked(func(button *ui.Button) {
		pcscommand.BDUSS = strings.TrimSpace(bdussEntry.Text())
		_, err := pcsconfig.Config.SetupUserByBDUSS(pcscommand.BDUSS, "", "")
		if err != nil {
			ui.MsgBoxError(loginWin, "登录失败", "请重新粘贴 BDUSS 后再试")
			bdussEntry.SetText("")
			return
		}

		pcsconfig.Config.Save()
		pcsconfig.Config.Reload()

		loginWin.Destroy()

		selectDownloadsDirWin := ui.NewWindow(pcscommand.WindowTitle, Width, Height, false)
		selectDownloadsDirWin.SetMargined(true)
		selectBox := ui.NewVerticalBox()
		selectBox.SetPadded(true)
		selectLabel := ui.NewLabel("请选择数据存放目录（Mac、Linux 版只能默认 ~/Downloads/）")
		selectBox.Append(selectLabel, true)
		selectCombobox := ui.NewCombobox()
		var saveDirs []string
		saveDirs = append(saveDirs, pcscommand.SaveDir)
		if pcsconfig.IsWindows() {
			var partition []string
			for _, drive := range "ABDEFGHIJKLMNOPQRSTUVWXYZ" {
				if _, err := os.Open(string(drive) + ":\\"); nil == err {
					saveDir := filepath.Join(string(drive)+":\\", "Downloads")
					partition = append(partition, saveDir)
				}
			}
			saveDirs = append(saveDirs, partition...)
		}
		for _, saveDir := range saveDirs {
			selectCombobox.Append(saveDir)
		}
		selectCombobox.SetSelected(0)
		selectBox.Append(selectCombobox, true)
		selectCombobox.OnSelected(func(combobox *ui.Combobox) {
			saveDir := saveDirs[combobox.Selected()]
			pcscommand.SaveDir = saveDir
			if err := os.MkdirAll(saveDir, os.ModePerm); nil != err {
				ui.MsgBoxError(selectDownloadsDirWin, "错误", "设置保存目录失败，请到论坛反馈问题，谢谢！")
				return
			}
			pcsconfig.Config.SetSaveDir(saveDir)
			fmt.Println("保存目录：" + pcsconfig.Config.SaveDir())
			pcsconfig.Config.Save()
		})
		selectButton := ui.NewButton("确定")
		selectButton.OnClicked(func(button *ui.Button) {
			selectDownloadsDirWin.Destroy()
			showMain()
		})
		selectBox.Append(selectButton, true)
		selectDownloadsDirWin.OnClosing(onClose)
		selectDownloadsDirWin.SetChild(selectBox)
		selectDownloadsDirWin.Show()
	})

	vBox.Append(loginButton, true)
	loginWin.SetMargined(true)
	loginWin.SetChild(vBox)
	loginWin.OnClosing(onClose)

	loginWin.Show()

	pcscommand.MainWin = ui.NewWindow(pcscommand.WindowTitle, Width, Height, true)
	pcscommand.MainWin.SetMargined(true)
}

func showMain() {
	panelBox := ui.NewVerticalBox()
	panelBox.SetPadded(true)
	progressBox := ui.NewHorizontalBox()
	pb := ui.NewProgressBar()
	progressLabelBox := ui.NewVerticalBox()
	pf := ui.NewLabel("尚未开始下载文件")
	pl := ui.NewLabel(pcscommand.ProgressFilenamePlaceholder)
	progressLabelBox.Append(pf, true)
	progressLabelBox.Append(pl, true)
	progressBox.Append(progressLabelBox, true)
	progressBox.Append(pb, true)
	refreshButton := ui.NewButton("1. 刷新文件列表")
	refreshButton.OnClicked(func(button *ui.Button) {
		filesBox, err := newFilesBox(pcscommand.MainWin)
		if nil != err {
			return
		}
		panelBox.Delete(2)
		panelBox.Append(filesBox, true)
	})
	downloadButton := ui.NewButton("3. 开始下载")
	downloadButton.OnClicked(func(button *ui.Button) {
		if "" == selectedFile || pcscommand.PlsSelectFile == selectedFile {
			ui.MsgBoxError(pcscommand.MainWin, "输入有误", pcscommand.PlsSelectFile)

			return
		}

		if pcscommand.Downloading {
			ui.MsgBoxError(pcscommand.MainWin, "禁止操作", "已经有文件正在下载，请等待下载完成或关闭程序重新打开")

			return
		}

		pl.SetText("正在进行数据缓冲，请稍后....")
		uiProgress := &pcscommand.UIProgress{PB: pb, PL: pl, PF: pf, Win: pcscommand.MainWin}
		pcscommand.Downloading = true
		go pcscommand.RunDownload([]string{selectedFile}, pcscommand.DownloadOption{
			IsTest:               false,
			IsPrintStatus:        false,
			IsExecutedPermission: !pcsconfig.IsWindows(),
			IsOverwrite:          false,
			SaveTo:               "",
			Parallel:             256,
		}, uiProgress)
	})
	buttonBox := ui.NewHorizontalBox()
	buttonBox.Append(refreshButton, true)
	buttonBox.Append(downloadButton, true)
	pauseButton := ui.NewButton("暂停下载")
	pauseButton.OnClicked(func(button *ui.Button) {
		ui.MsgBox(pcscommand.MainWin, "关于暂停", "如果你需要暂停下载，请直接关闭程序，下次重启后选择相同的文件进行下载即可！")
	})
	buttonBox.Append(pauseButton, true)
	bnd2Button := ui.NewButton("试试 BND2！")
	bnd2Button.OnClicked(func(button *ui.Button) {
		OpenBrowser(pcscommand.MainWin, pcscommand.BND2)
	})
	buttonBox.Append(bnd2Button, true)
	panelBox.Append(buttonBox, true)
	panelBox.Append(progressBox, true)
	filesBox, err := newFilesBox(pcscommand.MainWin)
	if nil == err {
		panelBox.Append(filesBox, true)
	} else {
		panelBox.Append(ui.NewVerticalBox(), true)
	}
	pcscommand.MainWin.SetChild(panelBox)
	pcscommand.MainWin.OnClosing(onClose)
	pcscommand.MainWin.Show()
}

func onClose(window *ui.Window) bool {
	ui.Quit()

	return true
}

type Files []map[string]interface{}

func (files Files) Len() int {
	return len(files)
}
func (files Files) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}
func (files Files) Less(i, j int) bool {
	return files[i]["Mtime"].(float64) > files[j]["Mtime"].(float64)
}

func newFilesBox(win *ui.Window) (box *ui.Box, err error) {
	output, err := pcscommand.RunLs("/")

	if nil != err {
		return nil, errors.New("数据解析失败")
	}

	var files Files
	if err = json.Unmarshal([]byte(output), &files); nil != err {
		msg := "数据解析失败"
		ui.MsgBoxError(win, "错误", msg)

		return nil, errors.New(msg)
	}

	sort.Sort(files)
	if 10 < len(files) {
		files = files[:10]
	}

	box = ui.NewVerticalBox()
	fileComboBox := ui.NewCombobox()
	fileComboBox.OnSelected(func(combobox *ui.Combobox) {
		idx := combobox.Selected()
		if idx < 1 {
			selectedFile = pcscommand.PlsSelectFile
		} else {
			selectedFile = files[idx-1]["Filename"].(string)
		}
	})
	fileComboBox.Append(pcscommand.PlsSelectFile)
	fileComboBox.SetSelected(0)
	for _, file := range files {
		fileComboBox.Append(file["Filename"].(string))
	}

	box.Append(fileComboBox, true)

	return
}

func unpackAria2() {
	if !pcsconfig.IsWindows() {
		return
	}

	aria2 := filepath.Join(pcscommand.BndDir, "aria2c.exe")
	_, err := os.Stat(aria2)
	if err == nil || os.IsExist(err) {
		return
	}

	data, err := ioutil.ReadFile("aria2c.zip")
	if nil != err {
		os.Exit(-1)
	}

	packZipPath := filepath.Join(pcscommand.BndDir, "aria2.zip")
	if err := ioutil.WriteFile(packZipPath, data, os.ModePerm); nil != err {
		fmt.Println("w", err)
		return
	}

	if err := Unzip(packZipPath, pcscommand.BndDir); nil != err {
		return
	}

	os.Remove(packZipPath)
}
