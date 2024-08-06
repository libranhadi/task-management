package service

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"task-management/model"
	"task-management/repository"
	"time"
)

var allowedExtensions = map[string]bool{
	".png":  true,
	".jpeg": true,
	".jpg":  true,
}

type FileService struct {
	repo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
	return &FileService{repo: repo}
}

func isValidFileExtension(filename string) bool {
	ext := filepath.Ext(filename)
	_, valid := allowedExtensions[strings.ToLower(ext)]
	return valid
}

func (fs *FileService) UploadFile(originalName string, file io.Reader) (model.File, error) {
	if !isValidFileExtension(originalName) {
		return model.File{}, errors.New("invalid file format; only .png, .jpeg, and .jpg are allowed")
	}

	timestamp := time.Now().Format("20060102150405")
	ext := filepath.Ext(originalName)
	baseName := filepath.Base(originalName[:len(originalName)-len(ext)])
	filename := fmt.Sprintf("%s_%s%s", baseName, timestamp, ext)

	err := fs.repo.SaveFileToDisk(filename, file)
	if err != nil {
		return model.File{}, err
	}

	return fs.repo.SaveFile(originalName, filename)
}

func (fs *FileService) ListFile() ([]model.File, error) {
	return fs.repo.ListFile()
}

func (fs *FileService) GetFilePath(filename string) string {
	return fs.repo.GetFilePath(filename)
}
