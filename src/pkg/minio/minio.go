package minio

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	minioClient "github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClientMinio(cfg *config.Config) *minioClient.Client {
	// Initialize minio client object.
	client, err := minioClient.New(cfg.Minio.Endpoint, &minioClient.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.SecretKey, ""),
		Secure: cfg.Minio.Ssl,
	})

	if err != nil {
		panic(err)
	}

	return client
}
