package command

import (
	"github.com/b3log/bnd2/util"
)

type ls struct {
}

func (cmd *ls) Exec(param map[string]interface{}) {
	ret := util.NewCmdResult(cmd.Name())
	path := param["path"].(string)
	by := param["by"].(string)
	order := param["order"].(string)

	ret.Data = util.Ls(path, by, order)

	util.Push(ret.Bytes())
}

func (cmd *ls) Name() string {
	return "ls"
}
