// +build !windows

package main

import (
	"os/exec"
	"runtime"

	"github.com/andlabs/ui"
)

func OpenBrowser(win *ui.Window, url string) {
	cmd := exec.Command(map[string]string{"darwin": "open", "linux": "xdg-open"}[runtime.GOOS], url)
	if err := cmd.Start(); nil != err {
		ui.MsgBoxError(win, "错误", "打开浏览器失败，请到论坛反馈问题，感谢！")
	}
}
