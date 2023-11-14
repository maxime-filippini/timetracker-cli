package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/maxime-filippini/gott-cli/cmd"
)

func rootHandler(args []string) error {

	if len(args) < 1 {
		return errors.New("no sub-command passed")
	}

	cmds := []cmd.Command{
		cmd.NewAddTaskCommand(),
		cmd.NewDeleteTaskCommand(),
		cmd.NewAddTimeCommand(),
		cmd.NewDeleteTimeCommand(),
		cmd.NewReportCommand(),
	}

	subcommand := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err := cmd.Init(args[1:])

			if err != nil {
				return err
			}

			return cmd.Run()

		}
	}

	return fmt.Errorf("unknown subcommand [%s]", subcommand)

}

func main() {

	if err := rootHandler(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// dbb := db.NewDatabase("data/db.json")

	// x := dbb.FilterEntriesByDate(
	// 	time.Date(2023, 11, 12, 0, 0, 0, 0, time.UTC),
	// 	time.Date(2023, 11, 13, 0, 0, 0, 0, time.UTC),
	// )

	// out := dbb.AggregateAcrossTasks(x)

	// fmt.Println(out)

	// dbb.AddTask(
	// 	"TEST", "Perform some test", "Perform some test",
	// )

	// dbb.AddTimeEntry("TEST", "2023-11-12", 10)
	// dbb.AddTimeEntry("TEST", "2023-11-12", 10)
	// dbb.AddTimeEntry("TEST", "2023-11-12", 10)
	// dbb.AddTimeEntry("TEST", "2023-11-12", 10)

	// // dbb.DeleteTask("TEST")

	// // dbb.DeleteTime(1)
	// // dbb.AddTask("task", "a new task")
	// // dbb.AddTask("task", "a new task")
	// // dbb.AddTask("task", "a new task")
	// dbb.Save("data/out.json")

}
