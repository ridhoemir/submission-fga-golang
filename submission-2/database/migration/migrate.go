package main

import (
	"submission-2/core"
	"submission-2/database"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(core.Order{}, core.Item{})
}
