package repository

import (
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
)

type repository struct {
	storage storage.Storage
}

func New(storage storage.Storage) domain.Repository {
	return &repository{storage}
}
