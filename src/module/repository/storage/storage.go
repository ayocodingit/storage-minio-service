package repository

import (
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	minio *minio.Client
}

func New(minio *minio.Client) domain.Repository {
	return &repository{minio}
}
