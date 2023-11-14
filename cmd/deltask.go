package cmd

import (
	"errors"
	"flag"

	"github.com/maxime-filippini/gott-cli/common"
	"github.com/maxime-filippini/gott-cli/db"
)

type DeleteTaskCommand struct {
	fs          *flag.FlagSet
	commonFlags *CommonFlags
	taskId      string
}

func NewDeleteTaskCommand() *DeleteTaskCommand {
	c := &DeleteTaskCommand{
		fs:          flag.NewFlagSet("deltask", flag.ContinueOnError),
		commonFlags: NewCommonFlags(),
	}

	c.fs.StringVar(&c.taskId, "id", common.StringSentinel, "task ID")
	c.fs.StringVar(&c.taskId, "i", common.StringSentinel, "task ID")

	ParseCommonFlags(c.fs, c.commonFlags)

	return c

}

func (c *DeleteTaskCommand) Init(args []string) error {
	err := c.fs.Parse(args)

	if c.taskId == common.StringSentinel {
		return errors.New("taskId cannot be left empty")
	}

	return err
}

func (c *DeleteTaskCommand) Name() string {
	return "deltask"
}

func (c *DeleteTaskCommand) Run() error {
	db := db.NewDatabase(c.commonFlags.DbPath)
	db.DeleteTask(c.taskId)
	db.Save(c.commonFlags.DbPath)
	return nil
}
