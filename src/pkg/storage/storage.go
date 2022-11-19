package storage

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	minio *minio.Client
}

func New(cfg config.Config) Storage {
	return Storage{
		minio: NewMinioClient(cfg),
	}
}
