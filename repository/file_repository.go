package repository

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"task-management/model"
)

const fileDirectory = "./resources"

func init() {
	if err := os.MkdirAll(fileDirectory, os.ModePerm); err != nil {
		panic("Failed to create file directory: " + err.Error())
	}
}

type FileRepository struct {
	fileStore map[int]model.File
	mutex     sync.Mutex
	nextID    int
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		fileStore: make(map[int]model.File),
		nextID:    1,
	}
}

func (fr *FileRepository) SaveFile(originalName, pathName string) (model.File, error) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	file := model.File{
		ID:           fr.nextID,
		PathName:     pathName,
		OriginalName: originalName,
	}

	fr.fileStore[fr.nextID] = file
	fr.nextID++

	return file, nil
}
func (fr *FileRepository) GetFile(id int) (model.File, bool) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	file, exists := fr.fileStore[id]
	return file, exists
}

func (fr *FileRepository) ListFile() ([]model.File, error) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	var files []model.File
	for _, file := range fr.fileStore {
		files = append(files, file)
	}

	return files, nil
}

func (fr *FileRepository) SaveFileToDisk(filename string, file io.Reader) error {
	filePath := filepath.Join(fileDirectory, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}

func (fr *FileRepository) GetFilePath(filename string) string {
	return filepath.Join(fileDirectory, filename)
}
