package command

import (
	"sort"
	"sync"
	"time"

	"github.com/b3log/bnd2/util"
)

var ctasks = tasksSorted{}
var ctasksMutex = sync.Mutex{}

var clisting = false

type lsctasks struct {
}

func (cmd *lsctasks) Name() string {
	return "lsctasks"
}

type ctasksSorted []*dtask

func (c ctasksSorted) Len() int {
	return len(c)
}
func (c ctasksSorted) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c ctasksSorted) Less(i, j int) bool {
	return c[i].Completed > c[j].Completed
}

func (cmd *lsctasks) Exec(param map[string]interface{}) {
	clisting = true

	for {
		result := util.NewCmdResult(cmd.Name())
		sort.Sort(ctasks)
		result.Data = ctasks
		util.Push(result.Bytes())

		if 0 == len(ctasks) {
			break
		}

		if !clisting {
			break
		}

		time.Sleep(time.Second)
	}

	result := util.NewCmdResult(cmd.Name())
	sort.Sort(ctasks)
	result.Data = ctasks
	util.Push(result.Bytes())
}

type stoplsctasks struct {
}

func (cmd *stoplsctasks) Name() string {
	return "stoplsctasks"
}

func (cmd *stoplsctasks) Exec(param map[string]interface{}) {
	clisting = false
}

func addCTask(t *dtask) {
	ctasksMutex.Lock()
	defer ctasksMutex.Unlock()

	ctasks = append(ctasks, t)
}

func getCTask(gid string) *dtask {
	ctasksMutex.Lock()
	defer ctasksMutex.Unlock()

	for _, ctask := range ctasks {
		if ctask.Gid == gid {
			return ctask
		}
	}

	return nil
}

func deleteAllCTasks() {
	ctasksMutex.Lock()
	defer ctasksMutex.Unlock()

	ctasks = ctasks[:0]
}

func deleteCTask(gid string) {
	ctasksMutex.Lock()
	defer ctasksMutex.Unlock()

	tmp := ctasks[:0]
	for _, ctask := range ctasks {
		if ctask.Gid != gid {
			tmp = append(tmp, ctask)
		}
	}
	ctasks = tmp
}
