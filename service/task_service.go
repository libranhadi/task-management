package service

import (
	"errors"
	"task-management/model"
	"task-management/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(taskRequest model.TaskRequest) (model.Task, error) {
	task, err := taskRequest.ToTask()
	if err != nil {
		return task, err
	}
	task.Status = "Pending"
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask(id int) (model.Task, error) {
	return s.repo.GetTask(id)
}

func (s *TaskService) UpdateTask(taskRequest model.TaskRequest, id int) (model.Task, error) {
	task, err := taskRequest.ToTask()
	if err != nil {
		return task, err
	}

	if id <= 0 {
		return model.Task{}, errors.New("invalid task ID")
	}

	task.ID = id
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) ListTasks() ([]model.Task, error) {
	return s.repo.ListTasks()
}
