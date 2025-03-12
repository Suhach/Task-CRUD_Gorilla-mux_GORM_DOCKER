package main

import (
	"log"
	"test_ByMyself/internal/database"
	taskservice "test_ByMyself/internal/taskService"
)

func main() {
	database.InitDB()

	err := database.DB.AutoMigrate(&taskservice.Task{})
	if err != nil {
		log.Fatal("migration error", err)
	}
}
