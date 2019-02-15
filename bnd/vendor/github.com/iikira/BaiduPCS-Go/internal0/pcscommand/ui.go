package pcscommand

import "github.com/andlabs/ui"

var Downloading = false
var MainWin *ui.Window

type UIProgress struct {
	PB  *ui.ProgressBar // 进度条
	PL  *ui.Label       // 进度文本
	PF  *ui.Label       // 文件名
	Win *ui.Window      // 主窗口
}

func ShortLabel(path string) string {
	ret := path
	runes := []rune(path)
	const max = 56
	if max < len(runes) {
		ret = string(runes[:max]) + "..."
	}

	return ret
}
