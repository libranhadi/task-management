package router

import (
	"net/http"
	"task-management/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(deps *RouterDependencies) *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/register", deps.UserController.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", deps.UserController.Login).Methods(http.MethodPost)

	authenticated := r.PathPrefix("/").Subrouter()
	authenticated.Use(middleware.AuthMiddleware)
	authenticated.HandleFunc("/profile", deps.UserController.GetProfile).Methods(http.MethodGet)

	taskRoutes := authenticated.PathPrefix("/tasks").Subrouter()

	taskRoutes.HandleFunc("", deps.TaskController.CreateTask).Methods(http.MethodPost)
	taskRoutes.HandleFunc("/{id:[0-9]+}", deps.TaskController.GetTask).Methods(http.MethodGet)
	taskRoutes.HandleFunc("/{id:[0-9]+}", deps.TaskController.UpdateTask).Methods(http.MethodPut)
	taskRoutes.HandleFunc("/{id:[0-9]+}", deps.TaskController.DeleteTask).Methods(http.MethodDelete)
	taskRoutes.HandleFunc("", deps.TaskController.ListTasks).Methods(http.MethodGet)

	fileManagementRoute := authenticated.PathPrefix("/file-management").Subrouter()
	fileManagementRoute.HandleFunc("/upload", deps.FileController.UploadFile).Methods(http.MethodPost)
	fileManagementRoute.HandleFunc("/files", deps.FileController.ListFile).Methods(http.MethodGet)
	fileManagementRoute.HandleFunc("/files/{filename}", deps.FileController.DownloadFile).Methods(http.MethodGet)

	return r
}
