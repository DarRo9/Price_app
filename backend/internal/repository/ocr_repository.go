package repository

import (
	"sync"

	"github.com/otiai10/gosseract/v2"
)

type tesseractOCRRepository struct {
	client      *gosseract.Client
	clientMutex sync.Mutex
}

func NewTesseractOCRRepository() (OCRRepository, error) {
	client := gosseract.NewClient()
	if err := client.SetLanguage("rus"); err != nil {
		return nil, err
	}

	return &tesseractOCRRepository{
		client: client,
	}, nil
}

func (r *tesseractOCRRepository) ExtractText(imagePath string) (string, error) {
	r.clientMutex.Lock()
	defer r.clientMutex.Unlock()

	r.client.SetImage(imagePath)
	return r.client.Text()
}

func (r *tesseractOCRRepository) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
