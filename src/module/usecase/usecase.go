package usecase

import (
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
)

type usecase struct {
	Repository domain.Repository
}

func New(repository domain.Repository) domain.Usecase {
	return usecase{repository}
}
