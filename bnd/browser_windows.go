package main

import (
	"os/exec"
	"syscall"

	"github.com/andlabs/ui"
)

func OpenBrowser(win *ui.Window, url string) {
	params := []string{"/C", "start", url}
	cmd := exec.Command("cmd", params...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); nil != err {
		ui.MsgBoxError(win, "错误", "打开浏览器失败，请到论坛反馈问题，感谢！")
	}
}
