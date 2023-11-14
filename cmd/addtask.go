package cmd

import (
	"errors"
	"flag"

	"github.com/maxime-filippini/gott-cli/common"
	"github.com/maxime-filippini/gott-cli/db"
)

type AddTaskCommand struct {
	fs              *flag.FlagSet
	commonFlags     *CommonFlags
	taskId          string
	taskName        string
	taskDescription string
}

func NewAddTaskCommand() *AddTaskCommand {
	c := &AddTaskCommand{
		fs:          flag.NewFlagSet("addtask", flag.ContinueOnError),
		commonFlags: NewCommonFlags(),
	}

	c.fs.StringVar(&c.taskId, "id", common.StringSentinel, "task ID")
	c.fs.StringVar(&c.taskId, "i", common.StringSentinel, "task ID")

	c.fs.StringVar(&c.taskName, "name", common.StringSentinel, "task name")
	c.fs.StringVar(&c.taskName, "n", common.StringSentinel, "task name")

	c.fs.StringVar(&c.taskDescription, "desc", common.StringSentinel, "task description")
	c.fs.StringVar(&c.taskDescription, "d", common.StringSentinel, "task description")

	ParseCommonFlags(c.fs, c.commonFlags)

	return c

}

func (c *AddTaskCommand) Init(args []string) error {
	err := c.fs.Parse(args)

	if c.taskName == common.StringSentinel {
		return errors.New("taskName cannot be left empty")
	}

	return err
}

func (c *AddTaskCommand) Name() string {
	return "addtask"
}

func (c *AddTaskCommand) Run() error {
	db := db.NewDatabase(c.commonFlags.DbPath)
	db.AddTask(c.taskId, c.taskName, c.taskDescription)
	db.Save(c.commonFlags.DbPath)
	return nil
}
