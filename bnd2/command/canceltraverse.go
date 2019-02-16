package command

type canceltraverse struct {
}

func (cmd *canceltraverse) Exec(param map[string]interface{}) {
	traverseCmd.reset()
}

func (cmd *canceltraverse) Name() string {
	return "canceltraverse"
}
