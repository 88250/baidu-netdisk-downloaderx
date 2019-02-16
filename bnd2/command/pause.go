package command

type pause struct {
}

func (cmd *pause) Name() string {
	return "pause"
}

func (cmd *pause) Exec(param map[string]interface{}) {
	gid := param["gid"].(string)
	if t := getTask(gid); nil != t {
		t.pause()
	}
}
