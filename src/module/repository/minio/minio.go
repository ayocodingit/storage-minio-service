package minio

import (
	"context"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	minio *minio.Client
	cfg   config.Config
}

func New(minio *minio.Client, config config.Config) domain.Repository {
	return repository{minio, config}
}

func (r repository) Upload(ctx context.Context, file domain.File) error {
	if _, err := r.minio.FPutObject(ctx, r.cfg.Minio.Bucket, file.Name, file.Dest, minio.PutObjectOptions{ContentType: file.ContentType}); err != nil {
		return err
	}

	return nil
}

func (r repository) Download(ctx context.Context, file domain.File) error {
	if err := r.minio.FGetObject(ctx, r.cfg.Minio.Bucket, file.Name, file.Dest, minio.GetObjectOptions{}); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, filename string) error {
	if err := r.minio.RemoveObject(ctx, r.cfg.Minio.Bucket, filename, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}
