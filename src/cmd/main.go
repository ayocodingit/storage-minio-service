package main

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/pkg/minio"
	"github.com/ayocodingit/storage-minio-service/src/transport/http"
)

func main() {
	cfg := config.LoadConfig()
	_ = minio.NewClientMinio(cfg)
	r := http.NewTransportHttp()

	r.Run(":" + cfg.Port)
}
