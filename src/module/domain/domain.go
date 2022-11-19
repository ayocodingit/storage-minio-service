package domain

import (
	"context"
)

type File struct {
	Name        string `json:"name"`
	Dest        string `json:"dest,omitempty"`
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
}

type UploadResponse struct {
	Data File `json:"data"`
}

type Repository interface {
	Upload(context.Context, *File) error
}

type Usecase interface {
	Upload(context.Context, File) (UploadResponse, error)
}
