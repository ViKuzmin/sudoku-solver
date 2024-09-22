package image_processor

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const (
	// Величина на которую уменьшится батч. Нужна для того, чтобы в рамку
	// сканирования не входили границы соседних ячеек
	quadSideDelta = 12
	// Количество ячеек на поле
	cellCount = 9
)

var possibleNumber = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

type ImageProcessorV1 struct {
}

func NewImageProcessorV1() *ImageProcessorV1 {
	return &ImageProcessorV1{}
}

func (processor *ImageProcessorV1) ProcessImage(path string) [][]int {
	data := getBattlefield(path)

	fmt.Println(data)
	return data
}

func getBattlefield(path string) [][]int {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		panic(err)
	}

	// Находим поле
	rect := findBlackSquare(img)
	croppedImage := imaging.Crop(img, rect)
	croppedImage = imaging.AdjustContrast(croppedImage, 80)
	croppedImage = imaging.Grayscale(croppedImage)

	return cutImageToParts(croppedImage, rect)
}

func cutImageToParts(croppedImage *image.NRGBA, rect image.Rectangle) [][]int {
	var data = make([][]int, cellCount)
	stepX := float32(rect.Dx()) / float32(cellCount)
	stepY := float32(rect.Dy()) / float32(cellCount)

	client := gosseract.NewClient()
	client.SetPageSegMode(gosseract.PSM_SINGLE_LINE)
	client.SetVariable("tessedit_char_whitelist", "0123456789")

	defer client.Close()

	counter := 1
	coordinateX := 0
	coordinateY := 0

	for y := 0; coordinateY < cellCount; y += int(stepY) {
		data[coordinateY] = make([]int, cellCount)
		for x := 0; coordinateX < cellCount; x += int(stepX) {
			imagePart := imaging.Crop(croppedImage, image.Rectangle{
				Min: image.Point{
					X: x + quadSideDelta/2,
					Y: y + quadSideDelta/2,
				},
				Max: image.Point{
					X: x + int(stepX) - quadSideDelta/2,
					Y: y + int(stepY) - quadSideDelta/2,
				},
			})

			buf := new(bytes.Buffer)
			jpeg.Encode(buf, imagePart, nil)

			client.SetImageFromBytes(buf.Bytes())
			out, _ := client.Text()

			val, ok := possibleNumber[out]
			if ok {
				data[coordinateY][coordinateX] = val
			} else {
				data[coordinateY][coordinateX] = 0
			}

			counter++
			coordinateX++
		}
		coordinateX = 0
		coordinateY++
	}

	return data
}

func findBlackSquare(img image.Image) image.Rectangle {
	gray := rgbaToGray(img)
	width := gray.Bounds().Dx()
	height := gray.Bounds().Dy()

	quadSideLength := 0
	quadSide := 0
	startPoint := image.Point{}
	endPoint := image.Point{}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r, g, b, _ := gray.At(j, i).RGBA()
			if r == 0 && g == 0 && b == 0 {
				quadSideLength++
			}
		}

		if float32(quadSideLength) > 0.85*float32(width) {
			quadSide = quadSideLength
			startPoint.Y = i
			startPoint.X = int((float32(width) - float32(quadSide)) / 2)
			break
		}
		quadSideLength = 0
	}

	endPoint.X = startPoint.X + quadSide
	endPoint.Y = startPoint.Y + quadSide

	fmt.Println(quadSide)
	fmt.Println(startPoint)
	fmt.Println(endPoint)

	return image.Rectangle{
		Min: startPoint,
		Max: endPoint,
	}
}

// Пережатие изображения в черно-белое. Используется для четкого определения игровой области.
func rgbaToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	threshold := 62000
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			avg := (r + g + b) / 3
			if avg > uint32(threshold) {
				gray.Set(x, y, color.White)
			} else {
				gray.Set(x, y, color.Black)
			}
		}
	}

	return gray
}
