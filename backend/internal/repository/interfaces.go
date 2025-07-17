package repository

import "price_app/internal/domain"

type OCRRepository interface {
	ExtractText(imagePath string) (string, error)
}

type PriceExtractorRepository interface {
	ExtractPriceAndWeight(text string) (*domain.PriceInfo, error)
}
