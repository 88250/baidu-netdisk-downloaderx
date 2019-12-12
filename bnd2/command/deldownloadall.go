package command

import (
	"github.com/88250/bnd2/util"
)

type deldownloadall struct {
}

func (cmd *deldownloadall) Name() string {
	return "deldownloadall"
}

func (cmd *deldownloadall) Exec(param map[string]interface{}) {
	for _, task := range tasks {
		util.R.Remove(task.Gid)
	}
	deleteAllTasks()

	listing = false
	result := util.NewCmdResult(lstasksCmd.Name())
	result.Data = tasks
	util.Push(result.Bytes())
}
