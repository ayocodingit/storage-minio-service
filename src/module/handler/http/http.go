package http

import (
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
	f, _ := c.FormFile("file")

	filename := uuid.New().String() + filepath.Ext(f.Filename)

	file := domain.File{
		Name:        filename,
		ContentType: f.Header["Content-Type"][0],
		Dest:        h.cfg.Dst + "/" + filename,
		Url:         h.cfg.Minio.Url + h.cfg.Minio.Bucket + "/" + filename,
	}

	c.SaveUploadedFile(f, file.Dest)

	res, err := h.usecase.Upload(c.Request.Context(), file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)

	return
}
