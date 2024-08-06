package repository

import (
	"errors"
	"task-management/model"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskRepository struct {
	tasks  map[int]model.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

func (repo *TaskRepository) CreateTask(task model.Task) (model.Task, error) {
	task.ID = repo.nextID
	repo.nextID++
	repo.tasks[task.ID] = task
	return task, nil
}

func (repo *TaskRepository) GetTask(id int) (model.Task, error) {
	task, exists := repo.tasks[id]
	if !exists {
		return model.Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (repo *TaskRepository) UpdateTask(task model.Task) (model.Task, error) {
	if _, exists := repo.tasks[task.ID]; !exists {
		return model.Task{}, ErrTaskNotFound
	}
	repo.tasks[task.ID] = task
	return task, nil
}

func (repo *TaskRepository) DeleteTask(id int) error {
	if _, exists := repo.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(repo.tasks, id)
	return nil
}

func (repo *TaskRepository) ListTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	for _, task := range repo.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
