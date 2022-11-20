package usecase

import (
	"context"
	"os"

	"github.com/ayocodingit/storage-minio-service/src/module/domain"
)

type usecase struct {
	Repository domain.Repository
}

func New(repository domain.Repository) domain.Usecase {
	return usecase{repository}
}

func (uc usecase) Upload(ctx context.Context, file domain.File) (domain.UploadResponse, error) {
	if err := uc.Repository.Upload(ctx, file); err != nil {
		return domain.UploadResponse{}, err
	}

	os.Remove(file.Dest)

	return domain.UploadResponse{Data: file}, nil
}
