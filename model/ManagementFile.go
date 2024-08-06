package model

type File struct {
	ID           int `json:"id"`
	PathName     string
	OriginalName string
}
