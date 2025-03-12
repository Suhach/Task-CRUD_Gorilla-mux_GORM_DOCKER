package taskservice

import (
	"errors"
)

func CreateTaskService(task *Task) error {
	if task.Task == "" {
		return errors.New("incorrect data")
	}
	return CreateTask(task)
}

func GetAllTasksService() ([]Task, error) {
	return GetAllTasks()
}

func GetTaskByIDService(id uint) (Task, error) {
	return GetTaskByID(id)
}

func UpdateTaskByIDService(task *Task) error {
	return UpdateTaskByID(task)
}

func DeleteTaskByIDService(id uint) error {
	return DeleteTaskByID(id)
}
