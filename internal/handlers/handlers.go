package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	taskservice "test_ByMyself/internal/taskService"

	"github.com/gorilla/mux"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task taskservice.Task
	_ = json.NewDecoder(r.Body).Decode(&task)

	err := taskservice.CreateTaskService(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := taskservice.GetAllTasksService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	task, err := taskservice.GetTaskByIDService(uint(id))
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task taskservice.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uint(id)

	err := taskservice.UpdateTaskByIDService(&task)
	if err != nil {
		http.Error(w, "error to update task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := taskservice.DeleteTaskByIDService(uint(id))
	if err != nil {
		http.Error(w, "error to delete task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "task successfully deleted"})
}

func RegisterTaskRoutes(r *mux.Router) {
	r.HandleFunc("/api/task", GetAllTasksHandler).Methods("GET")
	r.HandleFunc("/api/task/{id}", GetTaskByIDHandler).Methods("GET")
	r.HandleFunc("/api/task", CreateTaskHandler).Methods("POST")
	r.HandleFunc("/api/task/{id}", UpdateTaskByIDHandler).Methods("PATCH")
	r.HandleFunc("/api/task/{id}", DeleteTaskByIDHandler).Methods("DELETE")
}
