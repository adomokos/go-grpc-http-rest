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
	db, err := gorm.Open("sqlite3", "db/todo-db.sqlt")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Enable Logger, show detailed log
	db.LogMode(true)

	var todo = Todo{Title: "My Todo", Description: "Some interesting todo", Reminder: time.Now()}
	db.Create(&todo)
	fmt.Println(todo.ID)

	var todoToLookup = Todo{ID: 1}
	db.First(&todoToLookup)

	fmt.Println(todoToLookup)

	todoToLookup.Title = "My Todo - updated"
	db.Save(&todoToLookup)

	var todos []Todo
	db.Find(&todos)
	fmt.Print(todos)
}
