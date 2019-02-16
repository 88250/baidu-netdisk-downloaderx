package command

import (
	"github.com/b3log/bnd2/util"
)

type delctaskall struct {
}

func (cmd *delctaskall) Name() string {
	return "delctaskall"
}

func (cmd *delctaskall) Exec(param map[string]interface{}) {
	deleteAllCTasks()

	clisting = false
	result := util.NewCmdResult(lsctasksCmd.Name())
	result.Data = ctasks
	util.Push(result.Bytes())
}
