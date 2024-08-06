package controller

import (
	"encoding/json"
	"net/http"
	"task-management/service"
	"task-management/utils"

	"github.com/gorilla/mux"
)

type FileController struct {
	fileService *service.FileService
}

func NewFileController(fileService *service.FileService) *FileController {
	return &FileController{fileService: fileService}
}

func (fc *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file_upload")
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid file upload")
		return
	}
	defer file.Close()

	originalName := fileHeader.Filename

	if originalName == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Original filename is required")
		return
	}

	uploaded, err := fc.fileService.UploadFile(originalName, file)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "File uploaded successfully",
		"file":    uploaded,
	})
}

func (fc *FileController) ListFile(w http.ResponseWriter, r *http.Request) {
	files, err := fc.fileService.ListFile()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to list files")
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Files listed successfully",
		"files":   files,
	})
}

func (fc *FileController) DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	filePath := fc.fileService.GetFilePath(filename)
	http.ServeFile(w, r, filePath)
}
