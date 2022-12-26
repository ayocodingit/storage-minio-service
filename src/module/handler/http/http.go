package http

import (
	"net/http"
	"path/filepath"

	"github.com/ayocodingit/storage-minio-service/src/config"
	"github.com/ayocodingit/storage-minio-service/src/domain"
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
	r.GET("download/:filename", h.download)
	r.DELETE("delete/:filename", h.delete)
}

func (h handler) upload(c *gin.Context) {
	f, _ := c.FormFile("file")

	if f == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No upload image",
		})
		return
	}

	filename := uuid.New().String() + filepath.Ext(f.Filename)

	file := domain.File{
		Name:        filename,
		ContentType: f.Header["Content-Type"][0],
		Dest:        h.cfg.Dst + "/" + filename,
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

func (h handler) download(c *gin.Context) {
	filename := c.Param("filename")

	file := domain.File{
		Name: filename,
		Dest: h.cfg.Dst + "/" + filename,
	}

	fileBytes, err := h.usecase.Download(c.Request.Context(), file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
	c.Writer.Write(fileBytes)

	return
}

func (h handler) delete(c *gin.Context) {
	filename := c.Param("filename")

	if err := h.usecase.Delete(c.Request.Context(), filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Remove file with filename is " + filename,
	})
	return
}
