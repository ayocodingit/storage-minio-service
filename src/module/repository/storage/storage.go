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
	return repository{storage, config}
}

func (r repository) Upload(ctx context.Context, file domain.File) error {
	if _, err := r.storage.Minio.FPutObject(ctx, r.cfg.Minio.Bucket, file.Name, file.Dest, minio.PutObjectOptions{ContentType: file.ContentType}); err != nil {
		return err
	}

	return nil
}

func (r repository) Download(ctx context.Context, file domain.File) error {
	if err := r.storage.Minio.FGetObject(ctx, r.cfg.Minio.Bucket, file.Name, file.Dest, minio.GetObjectOptions{}); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, filename string) error {
	if err := r.storage.Minio.RemoveObject(ctx, r.cfg.Minio.Bucket, filename, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}
