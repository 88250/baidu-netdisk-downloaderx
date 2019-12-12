package command

import (
	"time"

	"github.com/88250/bnd2/util"
)

type counttasks struct {
}

func (cmd *counttasks) Name() string {
	return "counttasks"
}

var count = 0

func (cmd *counttasks) Exec(param map[string]interface{}) {
	for {
		time.Sleep(time.Second)

		curCnt := len(tasks)
		if curCnt == count {
			continue
		}

		count = curCnt

		result := util.NewCmdResult(cmd.Name())
		result.Data = map[string]interface{}{
			"taskCount": curCnt,
		}

		util.Push(result.Bytes())
	}
}
