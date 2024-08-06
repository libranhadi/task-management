package router

import "task-management/controller"

type RouterDependencies struct {
	UserController *controller.UserController
	TaskController *controller.TaskController
	FileController *controller.FileController
}
