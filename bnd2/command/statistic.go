package command

import (
	"time"

	"strconv"

	"github.com/b3log/bnd2/util"
	"github.com/dustin/go-humanize"
)

var statlisting = false
var speeds []uint64
var hSpeeds []string

type statistic struct {
}

func (cmd *statistic) Name() string {
	return "statistic"
}

type tasksStat struct {
	TotalSize    uint64  `json:"totalSize"`
	CurrentSize  uint64  `json:"currentSize"`
	HTotalSize   string  `json:"hTotalSize"`
	HCurrentSize string  `json:"hCurrentSize"`
	Progress     float64 `json:"progress"`
	Speed        uint64  `json:"speed"`
	HSpeed       string  `json:"hSpeed"`
	TaskCount    int     `json:"taskCount"`
}

type ctasksStat struct {
	TotalSize  uint64 `json:"totalSize"`
	HTotalSize string `json:"hTotalSize"`
	CTaskCount int    `json:"ctaskCount"`
}

func (cmd *statistic) Exec(param map[string]interface{}) {
	statlisting = true

	for {
		if !statlisting {
			break
		}

		statTasks := statTasks()

		result := util.NewCmdResult(cmd.Name())
		speeds = append(speeds, statTasks.Speed)
		hSpeeds = append(hSpeeds, humanize.Bytes(statTasks.Speed))
		if 60 < len(speeds) {
			speeds = speeds[1:]
			hSpeeds = hSpeeds[1:]
		}

		result.Data = map[string]interface{}{
			"tasks":   statTasks,
			"ctasks":  statCTasks(),
			"speeds":  speeds,
			"hSpeeds": hSpeeds,
			"user":    util.User,
		}

		util.Push(result.Bytes())
		time.Sleep(time.Second)
	}
}

type stopstatistic struct {
}

func (cmd *stopstatistic) Name() string {
	return "stopstatistic"
}

func (cmd *stopstatistic) Exec(param map[string]interface{}) {
	statlisting = false
}

func statTasks() *tasksStat {
	ret := &tasksStat{
		HSpeed: "0 B",
	}
	for _, task := range tasks {
		ret.TotalSize += task.Size
		ret.CurrentSize += task.CSize
	}

	ret.HTotalSize = humanize.Bytes(ret.TotalSize)
	ret.HCurrentSize = humanize.Bytes(ret.CurrentSize)

	if 0 < ret.TotalSize {
		ret.Progress = float64(ret.CurrentSize) / float64(ret.TotalSize) * 100
	}
	if nil != util.R {
		info, err := util.R.GetGlobalStat()
		if nil != err {
			logger.Errorf("get global stat failed [%s]", err)
		} else {
			s, _ := strconv.ParseInt(info.DownloadSpeed, 10, 64)
			ret.Speed = uint64(s)
			ret.HSpeed = humanize.Bytes(ret.Speed)
		}
	}
	ret.TaskCount = len(tasks)

	return ret
}

func statCTasks() *ctasksStat {
	ret := &ctasksStat{}
	for _, ctask := range ctasks {
		ret.TotalSize += ctask.Size
	}

	ret.HTotalSize = humanize.Bytes(ret.TotalSize)
	ret.CTaskCount = len(ctasks)

	return ret
}
