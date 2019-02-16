package command

import (
	"github.com/b3log/bnd2/util"
	"github.com/dustin/go-humanize"
)

type traverse struct {
	all, dirs, files []*util.File
	size             uint64
	traversing       bool
}

func (cmd *traverse) reset() {
	cmd.all = []*util.File{}
	cmd.dirs = []*util.File{}
	cmd.files = []*util.File{}
	cmd.size = 0
	cmd.traversing = false
}

func (cmd *traverse) Exec(param map[string]interface{}) {
	path := param["path"].(string)
	cmd.reset()

	cmd.traverse0(path)
}

func (cmd *traverse) Name() string {
	return "traverse"
}

func (cmd *traverse) traverse0(path string) {
	cmd.traversing = true
	defer func() { cmd.traversing = false }()

	ret := util.NewCmdResult(cmd.Name())
	ret.Data = map[string]interface{}{
		"all":      0,
		"dirs":     0,
		"files":    0,
		"hSize":    0,
		"finished": false,
	}

	rootFiles := util.Ls(path, "name", "asc")
	for i := 0; i < len(rootFiles); i++ {
		if !cmd.traversing {
			return
		}

		file := rootFiles[i]
		cmd.all = append(cmd.all, file)

		result := util.NewCmdResult(cmd.Name())

		if 0 == file.IsDir {
			cmd.files = append(cmd.files, file)
			cmd.size += file.Size

			result.Data = map[string]interface{}{
				"all":      len(cmd.all),
				"dirs":     len(cmd.dirs),
				"files":    len(cmd.files),
				"hSize":    humanize.Bytes(cmd.size),
				"finished": false,
			}
			util.Push(result.Bytes())
		} else {
			cmd.dirs = append(cmd.dirs, file)
			result.Data = map[string]interface{}{
				"all":      len(cmd.all),
				"dirs":     len(cmd.dirs),
				"files":    len(cmd.files),
				"hSize":    humanize.Bytes(cmd.size),
				"finished": false,
			}
			util.Push(result.Bytes())

			cmd.traverse1(file.Path)
		}
	}

	ret.Data = map[string]interface{}{
		"all":      len(cmd.all),
		"dirs":     len(cmd.dirs),
		"files":    len(cmd.files),
		"hSize":    humanize.Bytes(cmd.size),
		"finished": true,
	}

	util.Push(ret.Bytes())
}

func (cmd *traverse) traverse1(path string) {
	if !cmd.traversing {
		return
	}

	curFiles := util.Ls(path, "name", "asc")
	for i := 0; i < len(curFiles); i++ {
		if !cmd.traversing {
			return
		}

		file := curFiles[i]
		cmd.all = append(cmd.all, file)

		result := util.NewCmdResult(cmd.Name())

		if 0 == file.IsDir {
			cmd.files = append(cmd.files, file)
			cmd.size += file.Size

			result.Data = map[string]interface{}{
				"all":      len(cmd.all),
				"dirs":     len(cmd.dirs),
				"files":    len(cmd.files),
				"hSize":    humanize.Bytes(cmd.size),
				"finished": false,
			}
			util.Push(result.Bytes())
		} else {
			cmd.dirs = append(cmd.dirs, file)
			result.Data = map[string]interface{}{
				"all":      len(cmd.all),
				"dirs":     len(cmd.dirs),
				"files":    len(cmd.files),
				"hSize":    humanize.Bytes(cmd.size),
				"finished": false,
			}
			util.Push(result.Bytes())

			cmd.traverse1(file.Path)
		}
	}
}
