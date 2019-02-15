// +build !windows

package pcscommand

import (
	"github.com/iikira/BaiduPCS-Go/baidupcs"
)

func GetAria2DownloadFunc(uiProgress *UIProgress, path, savePath, s string) baidupcs.DownloadFunc {
	return nil
}
