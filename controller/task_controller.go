package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-management/model"
	"task-management/service"
	"task-management/utils"

	"github.com/gorilla/mux"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest model.TaskRequest

	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	createdTask, err := tc.taskService.CreateTask(taskRequest)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Task created successfully",
		"task":    createdTask,
	})
}

func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := tc.taskService.GetTask(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Task Not Found")
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Task retrieved successfully",
		"task":    task,
	})
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest model.TaskRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	updatedTask, err := tc.taskService.UpdateTask(taskRequest, id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Task updated successfully",
		"task":    updatedTask,
	})
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := tc.taskService.DeleteTask(id); err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Task Not Found")
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Task deleted successfully",
	})
}

func (tc *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := tc.taskService.ListTasks()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve tasks")
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Task retrieved successfully",
		"task":    tasks,
	})
}
