// Package pcscommand 命令包
package pcscommand

import (
	"github.com/iikira/BaiduPCS-Go/baidupcs"
	"github.com/iikira/BaiduPCS-Go/internal0/pcsconfig"
)

// GetActiveUser 获取当前登录的百度帐号
func GetActiveUser() *pcsconfig.Baidu {
	return pcsconfig.Config.ActiveUser()
}

// GetBaiduPCS 从配置读取BaiduPCS
func GetBaiduPCS() *baidupcs.BaiduPCS {
	return pcsconfig.Config.ActiveUserBaiduPCS()
}
