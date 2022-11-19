package storage

import (
	"context"
	"fmt"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(cfg config.Config) *minio.Client {
	// Initialize minio client object.
	client, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.SecretKey, ""),
		Secure: cfg.Minio.Ssl,
	})

	if err != nil {
		panic(err)
	}

	exists, err := client.BucketExists(context.TODO(), cfg.Minio.Bucket)
	if err != nil {
		panic(err)
	}

	if !exists {
		err = fmt.Errorf("bucket not already")
		panic(err)
	}

	return client
}
