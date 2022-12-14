package module

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/module/handler/http"
	repoMinio "github.com/ayocodingit/storage-minio-service/src/module/repository/minio"
	"github.com/ayocodingit/storage-minio-service/src/module/usecase"
	"github.com/ayocodingit/storage-minio-service/src/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Http(r *gin.Engine, storage storage.Storage, cfg config.Config) {
	repository := repoMinio.New(storage.Minio, cfg)
	usecase := usecase.New(repository)
	http := http.New(usecase, cfg)

	http.Handler(r)
}
