package http

import (
	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/gin-gonic/gin"
)

func NewTransportHttp(cfg config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	if cfg.IsPublicAccess {
		r.Static(cfg.Dst, cfg.Dst)
	}

	return r
}
