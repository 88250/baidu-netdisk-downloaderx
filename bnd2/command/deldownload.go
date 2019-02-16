package command

import "github.com/b3log/bnd2/util"

type deldownload struct {
}

func (cmd *deldownload) Name() string {
	return "deldownload"
}

func (cmd *deldownload) Exec(param map[string]interface{}) {
	gid := param["gid"].(string)
	util.R.Remove(gid)
	deleteTask(gid)
}
