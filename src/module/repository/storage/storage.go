package repository

import (
	"context"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	storage storage.Storage
	cfg     config.Config
}

func New(storage storage.Storage, config config.Config) domain.Repository {
	return &repository{storage, config}
}

func (r repository) Upload(ctx context.Context, file *domain.File) error {
	// Upload the zip file with FPutObject
	_, err := r.storage.Minio.FPutObject(ctx, r.cfg.Minio.Bucket, file.Name, file.Dest, minio.PutObjectOptions{ContentType: file.ContentType})
	if err != nil {
		return err
	}

	return nil
}
