package main

import (
	"fmt"
	"log"

	"price_app/configs"
	"price_app/internal/delivery"
	"price_app/internal/repository"
	"price_app/internal/usecase"
)

func main() {
	config := configs.LoadConfig()

	ocrRepo, err := repository.NewTesseractOCRRepository()
	if err != nil {
		log.Fatalf("Ошибка инициализации OCR: %v", err)
	}

	priceExtractorRepo := repository.NewPriceExtractorRepository()

	useCase := usecase.NewImageRecognitionUseCase(ocrRepo, priceExtractorRepo)

	handler := delivery.NewImageRecognitionHandler(useCase)

	router := delivery.SetupRouter(handler)

	fmt.Printf("Сервер запущен на порту %s\n", config.ServerPort)
	if err := router.Run(":" + config.ServerPort); err != nil {
		log.Fatal(err)
	}
}
