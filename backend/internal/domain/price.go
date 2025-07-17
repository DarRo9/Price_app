package domain

type PriceInfo struct {
	OriginalText string  `json:"original_text"`
	PricePerKg   float64 `json:"price_per_kg"`
	Weight       float64 `json:"weight"`
	Unit         string  `json:"unit"`
	Price        float64 `json:"price"`
	Confidence   float64 `json:"confidence"`
}

type ImageRecognitionRequest struct {
	ImagePath string `json:"image_path"`
}

type ImageRecognitionResponse struct {
	Text      string     `json:"text"`
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
	Error     string     `json:"error,omitempty"`
}
