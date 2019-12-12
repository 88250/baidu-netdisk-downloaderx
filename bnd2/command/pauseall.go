package command

import (
	"time"

	"github.com/88250/bnd2/util"
)

type pauseall struct {
}

func (cmd *pauseall) Name() string {
	return "pauseall"
}

func (cmd *pauseall) Exec(param map[string]interface{}) {
	ok, err := util.R.PauseAll()
	if nil != err {
		logger.Errorf("pause all tasks failed [%s]", err)

		return
	}

	for _, task := range tasks {
		task.HState = "暂停"
		task.State = statePaused
		task.Speed = "0 B"
	}

	logger.Infof("paused all tasks [%s]", ok)
	listing = false

	time.Sleep(time.Second)
	result := util.NewCmdResult(lstasksCmd.Name())
	result.Data = tasks
	util.Push(result.Bytes())
}
