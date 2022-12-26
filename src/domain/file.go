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
	Download(context.Context, File) error
	Delete(context.Context, string) error
}

type Usecase interface {
	Upload(context.Context, File) (UploadResponse, error)
	Download(context.Context, File) ([]byte, error)
	Delete(context.Context, string) error
}
