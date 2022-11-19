package main

import (
	"fmt"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
	"github.com/ayocodingit/storage-minio-service/src/transport/http"
)

func main() {
	cfg := config.LoadConfig()
	_ = storage.New(cfg)
	r := http.NewTransportHttp()

	fmt.Println(fmt.Sprintf("listen app on http://0.0.0.0:%s", cfg.Port))

	r.Run(":" + cfg.Port)
}
