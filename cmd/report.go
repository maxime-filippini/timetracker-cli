package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/maxime-filippini/gott-cli/common"
	"github.com/maxime-filippini/gott-cli/db"
)

type ReportCommand struct {
	fs          *flag.FlagSet
	commonFlags *CommonFlags
	startDate   string
	endDate     string
}

func NewReportCommand() *ReportCommand {
	c := &ReportCommand{
		fs:          flag.NewFlagSet("report", flag.ContinueOnError),
		commonFlags: NewCommonFlags(),
	}

	c.fs.StringVar(&c.startDate, "start", common.StringSentinel, "start date")
	c.fs.StringVar(&c.startDate, "s", common.StringSentinel, "start date")

	c.fs.StringVar(&c.endDate, "end", common.StringSentinel, "end date")
	c.fs.StringVar(&c.endDate, "e", common.StringSentinel, "end date")

	c.fs.StringVar(&c.endDate, "output", common.StringSentinel, "path to output")
	c.fs.StringVar(&c.endDate, "o", common.StringSentinel, "path to output")

	ParseCommonFlags(c.fs, c.commonFlags)

	return c

}

func (c *ReportCommand) Init(args []string) error {
	errParse := c.fs.Parse(args)

	if c.startDate == common.StringSentinel {
		c.startDate = "1999-01-01"
	}

	if c.endDate == common.StringSentinel {
		c.endDate = "2199-12-31"
	}

	if _, err := time.Parse("2006-01-02", c.startDate); err != nil {
		return errors.New("invalid date provided as start date")
	}

	if _, err := time.Parse("2006-01-02", c.endDate); err != nil {
		return errors.New("invalid date provided as end date")
	}

	return errParse
}

func (c *ReportCommand) Name() string {
	return "report"
}

func (c *ReportCommand) Run() error {
	db_ := db.NewDatabase(c.commonFlags.DbPath)

	start, _ := time.Parse("2006-01-02", c.startDate)
	end, _ := time.Parse("2006-01-02", c.endDate)

	entries := db_.FilterEntriesByDate(start, end)

	// Get unique tasks
	uniqueTasks := db.GetUniqueTasks(entries)

	// Get min, max from db
	minDate, maxDate := db.GetDateMinMax(entries)

	// Generate range
	var dateRange []time.Time
	currentDate := minDate

	for maxDate.Sub(currentDate) >= 0 {
		dateRange = append(dateRange, currentDate)
		currentDate = currentDate.Add(time.Hour * 24)
	}

	// Aggregate over the dates and tasks
	outMap := db_.AggregateAcrossTasks(entries, dateRange)

	header := "\t\t"

	for _, date := range dateRange {
		header += fmt.Sprintf("%s\t", date.Format("2006-01-02"))
	}

	divider := strings.Repeat("-", len(header)*3)

	fmt.Println(divider)
	fmt.Println("Your report")
	fmt.Println(divider)

	fmt.Println(header)

	for _, task := range uniqueTasks {
		fmt.Printf("%s\t", task)

		for _, date := range dateRange {
			fmt.Printf("%d\t\t", outMap[task][date])
		}

		fmt.Print("\n")
	}

	fmt.Println(divider)

	return nil
}
