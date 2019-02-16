package command

type unpauseall struct {
}

func (cmd *unpauseall) Name() string {
	return "unpauseall"
}

func (cmd *unpauseall) Exec(param map[string]interface{}) {
	for i, task := range tasks {
		if 5 > i {
			task.unpause()
		}
	}

	logger.Infof("unpaused all tasks")

	go lstasksCmd.Exec(map[string]interface{}{})
}
