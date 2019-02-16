package command

type unpause struct {
}

func (cmd *unpause) Name() string {
	return "unpause"
}

func (cmd *unpause) Exec(param map[string]interface{}) {
	gid := param["gid"].(string)
	if t := getTask(gid); nil != t {
		t.unpause()
	}

	go lstasksCmd.Exec(map[string]interface{}{})
}
