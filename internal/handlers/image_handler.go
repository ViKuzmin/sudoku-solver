package handlers

import (
	"fmt"
	"image"
	"net/http"
)

type ImageProcessor struct {
}

func (processor *ImageProcessor) ProcessImage(w http.ResponseWriter, r *http.Request) {
	// Парсим изображение из тела HTTP запроса
	img, _, err := image.Decode(r.Body)
	if err != nil {
		http.Error(w, "Ошибка при чтении изображения", http.StatusInternalServerError)
		return
	}

	// Производим необходимую обработку изображения здесь
	// Например, можно изменить размер, применить фильтры и т.д.

	// Пример: отображаем размеры изображения
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	fmt.Fprintf(w, "Ширина: %d, Высота: %d", width, height)
}
