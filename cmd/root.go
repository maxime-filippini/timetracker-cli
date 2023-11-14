package cmd

import (
	"flag"

	"github.com/maxime-filippini/gott-cli/common"
)

type Command interface {
	Init([]string) error
	Run() error
	Name() string
}

type CommonFlags struct {
	DbPath string
}

func NewCommonFlags() *CommonFlags {
	c := &CommonFlags{}
	return c
}

func ParseCommonFlags(fs *flag.FlagSet, flags *CommonFlags) {
	fs.StringVar(&flags.DbPath, "db", common.StringSentinel, "path to database")
}
