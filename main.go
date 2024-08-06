package main

import (
	"net/http"
	"task-management/controller"
	"task-management/repository"
	"task-management/router"
	"task-management/service"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)

	fileRepo := repository.NewFileRepository()
	fileService := service.NewFileService(fileRepo)
	fileController := controller.NewFileController(fileService)

	deps := &router.RouterDependencies{
		UserController: userController,
		TaskController: taskController,
		FileController: fileController,
	}

	r := router.NewRouter(deps)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
