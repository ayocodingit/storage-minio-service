package main

import (
	"fmt"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/lang"
	"github.com/ayocodingit/storage-minio-service/src/module"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
	"github.com/ayocodingit/storage-minio-service/src/transport/http"
	"github.com/ayocodingit/storage-minio-service/src/validator"
)

func main() {
	lang := lang.NewLang()
	validator := validator.New(lang)
	cfg := config.LoadConfig(validator)
	storage := storage.New(cfg)
	r := http.NewTransportHttp(cfg)

	module.Load(r, storage, cfg)

	fmt.Println(fmt.Sprintf("listen app on http://0.0.0.0:%s", cfg.Port))

	r.Run(":" + cfg.Port)
}
