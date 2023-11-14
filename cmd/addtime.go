package cmd

import (
	"errors"
	"flag"
	"time"

	"github.com/maxime-filippini/gott-cli/common"
	"github.com/maxime-filippini/gott-cli/db"
)

type AddTimeCommand struct {
	fs          *flag.FlagSet
	commonFlags *CommonFlags
	taskId      string
	timeSpent   int
}

func NewAddTimeCommand() *AddTimeCommand {
	c := &AddTimeCommand{
		fs:          flag.NewFlagSet("addtime", flag.ContinueOnError),
		commonFlags: NewCommonFlags(),
	}

	c.fs.StringVar(&c.taskId, "task", common.StringSentinel, "task ID")
	c.fs.StringVar(&c.taskId, "t", common.StringSentinel, "task ID")

	c.fs.IntVar(&c.timeSpent, "time", common.IntSentinel, "time spent")
	c.fs.IntVar(&c.timeSpent, "s", common.IntSentinel, "time spent")

	ParseCommonFlags(c.fs, c.commonFlags)

	return c

}

func (c *AddTimeCommand) Init(args []string) error {
	err := c.fs.Parse(args)

	if c.taskId == common.StringSentinel {
		return errors.New("taskId cannot be left empty")
	}

	return err
}

func (c *AddTimeCommand) Name() string {
	return "addtime"
}

func (c *AddTimeCommand) Run() error {
	db := db.NewDatabase(c.commonFlags.DbPath)

	date := time.Now().Format("2006-01-02")

	db.AddTimeEntry(c.taskId, date, c.timeSpent)
	db.Save(c.commonFlags.DbPath)
	return nil
}
