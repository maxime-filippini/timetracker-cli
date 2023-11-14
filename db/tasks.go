package db

import (
	"log"
)

// GetTasks retrieves all of the tasks in the database which do not match
// the sentinel.
func (db *Database) GetTasks() []Task {
	var out []Task
	for _, task := range db.Tasks {
		if task != sentinelTask {
			out = append(out, task)
		}
	}
	return out
}

// getTaskById finds the first occurence of a task that matches the
// provided id.
func (db *Database) getTaskById(id string) *Task {
	for _, task := range db.Tasks {
		if task.Id == id {
			return &task
		}
	}
	return nil
}

// AddTask adds a task to the database, if the ID does not exist.
func (db *Database) AddTask(id string, name string, desc string) {

	if db.getTaskById(id) != nil {
		log.Printf("Id [%s] already exists! No addition was performed.\n", id)
		return
	}

	task := Task{
		Id: id, Name: name, Desc: desc,
	}

	db.Tasks = append(db.Tasks, task)
	log.Printf("Task [%s] added!", task.Id)

}

// DeleteTask deletes the task which matches the ID provided.
func (db *Database) DeleteTask(id string) {
	for i, task := range db.Tasks {
		if task.Id == id {
			db.Tasks[i] = sentinelTask
			break
		}
	}

	for i, time := range db.Times {
		if time.TaskId == id {
			db.Times[i] = sentinelTime
		}
	}

}
