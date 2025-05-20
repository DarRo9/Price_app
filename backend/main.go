package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	r := gin.Default()

	r.StaticFile("/", "./index.html")

	r.POST("/upload", func(c *gin.Context) {
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

		client.SetImage(tempFile)
		text, err := client.Text()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при распознавании текста"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"text": text,
		})
	})

	fmt.Println("Сервер запущен на порту 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
