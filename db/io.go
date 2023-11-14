package db

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func (db *Database) readFile() {
	file, err := os.Open(db.path)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	byt, err := io.ReadAll(file)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(byt, db)

}

func (db *Database) cleanUp() {
	db.Times = db.GetTimes()
	db.Tasks = db.GetTasks()
}

func (db *Database) Save(path string) {

	db.cleanUp()

	byt, err := json.MarshalIndent(db, "", "  ")

	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(path, byt, 0644)
	log.Printf("Saved database to file [%s]\n", path)
}
