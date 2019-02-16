package command

import (
	"os"

	"github.com/b3log/bnd2/log"
)

var logger = log.NewLogger(os.Stdout)

type Cmd interface {
	Name() string
	Exec(map[string]interface{})
}

var Commands = map[string]Cmd{}

var (
	lsCmd             = &ls{}
	traverseCmd       = &traverse{}
	canceltraverseCmd = &canceltraverse{}
	downloaddirCmd    = &downloaddir{}
	downloadfileCmd   = &downloadfile{}
	pauseCmd          = &pause{}
	unpauseCmd        = &unpause{}
	pauseallCmd       = &pauseall{}
	unpauseallCmd     = &unpauseall{}
	deldownloadCmd    = &deldownload{}
	deldownloadallCmd = &deldownloadall{}
	lstasksCmd        = &lstasks{}
	stoplstasksCmd    = &stoplstasks{}
	lsctasksCmd       = &lsctasks{}
	stoplsctasksCmd   = &stoplsctasks{}
	statisticCmd      = &statistic{}
	stopstatisticCmd  = &stopstatistic{}
	delctaskCmd       = &delctask{}
	delctaskallCmd    = &delctaskall{}
	counttasksCmd     = &counttasks{}
)

func init() {
	Commands[lsCmd.Name()] = lsCmd
	Commands[traverseCmd.Name()] = traverseCmd
	Commands[canceltraverseCmd.Name()] = canceltraverseCmd
	Commands[downloaddirCmd.Name()] = downloaddirCmd
	Commands[downloadfileCmd.Name()] = downloadfileCmd
	Commands[pauseCmd.Name()] = pauseCmd
	Commands[unpauseCmd.Name()] = unpauseCmd
	Commands[pauseallCmd.Name()] = pauseallCmd
	Commands[unpauseallCmd.Name()] = unpauseallCmd
	Commands[deldownloadCmd.Name()] = deldownloadCmd
	Commands[deldownloadallCmd.Name()] = deldownloadallCmd
	Commands[lstasksCmd.Name()] = lstasksCmd
	Commands[stoplstasksCmd.Name()] = stoplstasksCmd
	Commands[lsctasksCmd.Name()] = lsctasksCmd
	Commands[stoplsctasksCmd.Name()] = stoplsctasksCmd
	Commands[statisticCmd.Name()] = statisticCmd
	Commands[stopstatisticCmd.Name()] = stopstatisticCmd
	Commands[delctaskCmd.Name()] = delctaskCmd
	Commands[delctaskallCmd.Name()] = delctaskallCmd
	Commands[counttasksCmd.Name()] = counttasksCmd
}
