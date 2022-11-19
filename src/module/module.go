package module

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/module/handler/http"
	repository "github.com/ayocodingit/storage-minio-service/src/module/repository/storage"
	"github.com/ayocodingit/storage-minio-service/src/module/usecase"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine, storage storage.Storage, cfg config.Config) {
	repository := repository.New(storage)
	usecase := usecase.New(repository)
	http := http.New(usecase, cfg)

	http.Handler(r)
}
