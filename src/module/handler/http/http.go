package http

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/module/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	usecase domain.Usecase
	cfg     config.Config
}

func New(usecase domain.Usecase, cfg config.Config) handler {
	return handler{usecase, cfg}
}

func (h handler) Handler(r *gin.Engine) {
	r.POST("upload", h.upload)
}

func (h handler) upload(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	_ = file.Header["Content-Type"][0]

	dst := h.cfg.Dst + "/" + filename
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", filename))

	return
}
