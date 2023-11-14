package db

import (
	"github.com/maxime-filippini/gott-cli/common"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Time struct {
	Id        int    `json:"id"`
	TaskId    string `json:"taskId"`
	Date      string `json:"date"`
	TimeSpent int    `json:"time_spent"`
}

type Database struct {
	path       string
	LastTimeId int    `json:"last_time_id"`
	Tasks      []Task `json:"tasks"`
	Times      []Time `json:"times"`
}

var sentinelTask = Task{
	Id:   common.StringSentinel,
	Name: common.StringSentinel,
	Desc: common.StringSentinel,
}

var sentinelTime = Time{
	Id:        common.IntSentinel,
	TaskId:    common.StringSentinel,
	Date:      common.StringSentinel,
	TimeSpent: common.IntSentinel,
}

func NewDatabase(path string) *Database {
	db := Database{path: path}
	db.readFile()
	return &db
}
