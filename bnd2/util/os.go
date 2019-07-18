package util

import (
	"os"
	"time"

	"github.com/mitchellh/go-ps"
)

var ppid = os.Getppid()

func ParentExited() {
	for range time.Tick(2 * time.Second) {
		process, e := ps.FindProcess(ppid)
		if nil == process || nil != e {
			logger.Info("can't find parent process, exited")

			os.Exit(0)
		}
	}
}
