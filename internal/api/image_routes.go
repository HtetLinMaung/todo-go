package api

import (
	"fmt"
	"net/http"

	"github.com/HtetLinMaung/todo/internal/utils"
	"github.com/gin-gonic/gin"
)

type ImageRoute struct{}

func NewImageRoute() *ImageRoute {
	return &ImageRoute{}
}

func (ir *ImageRoute) ImageRoutes(r *gin.Engine) {
	imageGroup := r.Group("/api/image")
	imageGroup.POST("/upload", ir.UploadImage)
}

func (ir *ImageRoute) UploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var images []string
	files := form.File["files[]"]
	for _, file := range files {
		fileName := utils.GenerateFileName(file.Filename)
		images = append(images, fmt.Sprintf("/images/%s", fileName))
		dst := "./images/" + fileName
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Error uploading file!",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Image uploaded successfully.",
		"urls":    images,
	})
}
