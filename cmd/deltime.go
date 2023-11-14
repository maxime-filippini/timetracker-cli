package cmd

import (
	"errors"
	"flag"

	"github.com/maxime-filippini/gott-cli/common"
	"github.com/maxime-filippini/gott-cli/db"
)

type DeleteTimeCommand struct {
	fs          *flag.FlagSet
	commonFlags *CommonFlags
	timeId      int
}

func NewDeleteTimeCommand() *DeleteTimeCommand {
	c := &DeleteTimeCommand{
		fs:          flag.NewFlagSet("deltime", flag.ContinueOnError),
		commonFlags: NewCommonFlags(),
	}

	c.fs.IntVar(&c.timeId, "id", common.IntSentinel, "task ID")
	c.fs.IntVar(&c.timeId, "i", common.IntSentinel, "task ID")

	ParseCommonFlags(c.fs, c.commonFlags)

	return c

}

func (c *DeleteTimeCommand) Init(args []string) error {
	err := c.fs.Parse(args)

	if c.timeId == common.IntSentinel {
		return errors.New("time ID cannot be left empty")
	}

	return err
}

func (c *DeleteTimeCommand) Name() string {
	return "deltime"
}

func (c *DeleteTimeCommand) Run() error {
	db := db.NewDatabase(c.commonFlags.DbPath)
	db.DeleteTime(c.timeId)
	db.Save(c.commonFlags.DbPath)
	return nil
}
