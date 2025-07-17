package usecase

import (
	"os"
	"path/filepath"

	"price_app/internal/domain"
	"price_app/internal/repository"
)

type ImageRecognitionUseCase struct {
	ocrRepo            repository.OCRRepository
	priceExtractorRepo repository.PriceExtractorRepository
}

func NewImageRecognitionUseCase(ocrRepo repository.OCRRepository, priceExtractorRepo repository.PriceExtractorRepository) *ImageRecognitionUseCase {
	return &ImageRecognitionUseCase{
		ocrRepo:            ocrRepo,
		priceExtractorRepo: priceExtractorRepo,
	}
}

func (uc *ImageRecognitionUseCase) ProcessImage(imagePath string) (*domain.ImageRecognitionResponse, error) {
	text, err := uc.ocrRepo.ExtractText(imagePath)
	if err != nil {
		return &domain.ImageRecognitionResponse{
			Text:  "",
			Error: "Ошибка при распознавании текста: " + err.Error(),
		}, nil
	}

	priceInfo, err := uc.priceExtractorRepo.ExtractPriceAndWeight(text)
	if err != nil {
		return &domain.ImageRecognitionResponse{
			Text:  text,
			Error: err.Error(),
		}, nil
	}

	return &domain.ImageRecognitionResponse{
		Text:      text,
		PriceInfo: priceInfo,
	}, nil
}

func (uc *ImageRecognitionUseCase) SaveUploadedFile(filePath, fileName string) (string, error) {
	tempFile := filepath.Join(os.TempDir(), fileName)
	return tempFile, nil
}
