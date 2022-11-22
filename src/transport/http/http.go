package http

import (
	"net/http"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/gin-contrib/cors"
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

	// middleware global
	r.Use(verify(cfg))

	// enabled cors
	r.Use(cors.Default())

	r.NoRoute(routeNotFound)

	return r
}

func routeNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error": "Not Found!",
	})
}

func verify(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Api-Key"]
		if len(header) == 0 {
			verifyError(c, "Unauthorized, Please set header Api-Key before send request")
		}

		if header[0] != cfg.Secret {
			verifyError(c, "Unauthorized, Api-Key not match with app secret")
		}
	}
}

func verifyError(c *gin.Context, error string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": error,
	})
	c.Abort()
}
