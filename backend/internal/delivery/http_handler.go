package delivery

import (
	"net/http"
	"os"
	"path/filepath"

	"price_app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ImageRecognitionHandler struct {
	useCase *usecase.ImageRecognitionUseCase
}

func NewImageRecognitionHandler(useCase *usecase.ImageRecognitionUseCase) *ImageRecognitionHandler {
	return &ImageRecognitionHandler{
		useCase: useCase,
	}
}

func (h *ImageRecognitionHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		return
	}

	tempFile := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл"})
		return
	}
	defer os.Remove(tempFile)

	response, err := h.useCase.ProcessImage(tempFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке изображения"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ImageRecognitionHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
