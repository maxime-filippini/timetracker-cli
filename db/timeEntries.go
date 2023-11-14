package db

import (
	"log"
	"time"
)

func (db *Database) GetTimes() []Time {
	var out []Time
	for _, time := range db.Times {
		if time.Id != sentinelTime.Id {
			out = append(out, time)
		}
	}
	return out
}

func (db *Database) DeleteTime(id int) {
	for i, time := range db.Times {
		if time.Id == id {
			db.Times[i] = sentinelTime
		}
	}
}

func (db *Database) AddTimeEntry(taskId string, date string, timeSpent int) {
	lastId := db.LastTimeId

	if db.getTaskById(taskId) == nil {
		log.Printf("Task id [%s] does not exist! Time entry was not added.\n", taskId)
		return
	}

	newEntry := Time{
		Id:        lastId + 1,
		TaskId:    taskId,
		Date:      date,
		TimeSpent: timeSpent,
	}

	db.Times = append(db.Times, newEntry)
	log.Println("Time entry added!")

	db.LastTimeId = lastId + 1

}

func (db *Database) FilterEntriesByDate(start time.Time, end time.Time) []Time {
	var out []Time

	for _, timeEntry := range db.GetTimes() {
		check, _ := time.Parse("2006-01-02", timeEntry.Date)

		fromStart := check.Sub(start)
		toEnd := end.Sub(check)

		if fromStart >= 0 && toEnd >= 0 {
			out = append(out, timeEntry)
		}
	}

	return out
}

func (db *Database) AggregateAcrossTasks(entries []Time, dates []time.Time) map[string]map[time.Time]int {
	taskTimeMap := make(map[string]map[time.Time]int)

	for _, entry := range entries {

		currentDate, _ := time.Parse("2006-01-02", entry.Date)

		if _, ok := taskTimeMap[entry.TaskId]; !ok {
			taskTimeMap[entry.TaskId] = make(map[time.Time]int)

			for _, date := range dates {
				taskTimeMap[entry.TaskId][date] = 0
			}

		}

		subMap := taskTimeMap[entry.TaskId]

		if _, ok := subMap[currentDate]; !ok {
			continue
		}

		taskTimeMap[entry.TaskId][currentDate] += entry.TimeSpent

	}

	return taskTimeMap

}

func GetUniqueTasks(entries []Time) []string {
	var uniqueEntries []string
	var visited bool

	for _, entry := range entries {
		visited = false

		for _, check := range uniqueEntries {
			if entry.TaskId == check {
				visited = true
				break
			}
		}

		if !visited {
			uniqueEntries = append(uniqueEntries, entry.TaskId)
		}
	}

	return uniqueEntries
}

func GetDateMinMax(entries []Time) (time.Time, time.Time) {
	minDate := time.Date(2199, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)

	var date time.Time

	for _, entry := range entries {

		date, _ = time.Parse("2006-01-02", entry.Date)

		if date.Sub(minDate) < 0 {
			minDate = date
		} else if date.Sub(maxDate) > 0 {
			maxDate = date
		}

	}

	return minDate, maxDate

}
