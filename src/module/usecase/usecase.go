package usecase

import (
	"context"
	"io/ioutil"
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

	file.Dest = ""

	return domain.UploadResponse{Data: file}, nil
}

func (uc usecase) Download(ctx context.Context, file domain.File) (fileBytes []byte, err error) {
	if err = uc.Repository.Download(ctx, file); err != nil {
		return
	}

	fileBytes, err = ioutil.ReadFile(file.Dest)
	if err != nil {
		return nil, err
	}

	os.Remove(file.Dest)

	return fileBytes, nil
}

func (uc usecase) Delete(ctx context.Context, filename string) error {
	if err := uc.Repository.Delete(ctx, filename); err != nil {
		return err
	}

	return nil
}
