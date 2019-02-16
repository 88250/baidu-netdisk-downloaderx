package command

type delctask struct {
}

func (cmd *delctask) Name() string {
	return "delctask"
}

func (cmd *delctask) Exec(param map[string]interface{}) {
	gid := param["gid"].(string)
	deleteCTask(gid)
}
