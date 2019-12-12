package command

import (
	"sort"
	"sync"
	"time"

	"github.com/88250/bnd2/util"
)

var tasks = tasksSorted{}
var tasksMutex = sync.Mutex{}

var listing = false

type lstasks struct {
}

func (cmd *lstasks) Name() string {
	return "lstasks"
}

type tasksSorted []*dtask

func (c tasksSorted) Len() int {
	return len(c)
}
func (c tasksSorted) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c tasksSorted) Less(i, j int) bool {
	if c[i].State < c[j].State {
		return true
	}

	if c[i].State == c[j].State {
		return c[i].Created <= c[j].Created
	}

	return false
}

func (cmd *lstasks) Exec(param map[string]interface{}) {
	listing = true

	for {
		result := util.NewCmdResult(cmd.Name())
		sort.Sort(tasks)
		result.Data = tasks
		util.Push(result.Bytes())

		if 0 == len(tasks) {
			break
		}

		if !listing {
			break
		}

		time.Sleep(time.Second)
	}

	result := util.NewCmdResult(cmd.Name())
	sort.Sort(tasks)
	result.Data = tasks
	util.Push(result.Bytes())
}

type stoplstasks struct {
}

func (cmd *stoplstasks) Name() string {
	return "stoplstasks"
}

func (cmd *stoplstasks) Exec(param map[string]interface{}) {
	listing = false
}

func containsTask(path string) bool {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	for _, task := range tasks {
		if task.Path == path {
			return true
		}
	}

	return false
}

func countDownloadingTask() (ret int) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	for _, task := range tasks {
		if stateDownloading == task.State {
			ret++
		}
	}

	return
}

func addTask(t *dtask) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tasks = append(tasks, t)
}

func getTask(gid string) *dtask {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	for _, task := range tasks {
		if task.Gid == gid {
			return task
		}
	}

	return nil
}

func deleteAllTasks() {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tasks = tasks[:0]
}

func deleteTask(gid string) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tmp := tasks[:0]
	for _, task := range tasks {
		if task.Gid != gid {
			tmp = append(tmp, task)
		}
	}
	tasks = tmp
}
