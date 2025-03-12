package taskservice

import (
	"test_ByMyself/internal/database"
)

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

func GetTaskByID(id uint) (Task, error) {
	var task Task
	err := database.DB.First(&task, id).Error
	return task, err
}

func CreateTask(task *Task) error {
	return database.DB.Create(task).Error
}

func UpdateTaskByID(task *Task) error {
	return database.DB.Save(task).Error
}

func DeleteTaskByID(id uint) error {
	return database.DB.Delete(&Task{}, id).Error
}
