package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Todo struct {
	ID          uint
	Title       string
	Description string
	Reminder    time.Time
}

func main() {
	db, _ := gorm.Open("sqlite3", "db/todo-db.sqlt")
	defer db.Close()

	// Enable Logger, show detailed log
	db.LogMode(true)

	// var todo = Todo{Title: "My Todo", Description: "Some interesting todo", Reminder: time.Now()}
	// db.Create(&todo)

	var todos []Todo
	db.Find(&todos)
	fmt.Print(todos)
}
