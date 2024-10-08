package image_processor

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/jpeg"
	"log/slog"
	"os"
)

const (

	// The amount by which the scanning batch size will be reduced.
	// Used to exclude borders of adjacent cells.
	quadSideDelta = 12
	// Cell count
	cellCount = 9
	threshold = 62000
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
	logger *slog.Logger
}

func NewImageProcessorV1(logger *slog.Logger) *ImageProcessorV1 {
	return &ImageProcessorV1{
		logger: logger,
	}
}

func (processor *ImageProcessorV1) ProcessImage(path string) [][]int {
	f, err := os.Open(path)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("failed to open file: %s", path))
		return nil
	}
	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		processor.logger.Error(fmt.Sprintf("failed to decode file: %s", path))
		return nil
	}

	data := processor.GetBattlefield(img)

	fmt.Println(data)
	return data
}

func (processor *ImageProcessorV1) GetBattlefield(img image.Image) [][]int {
	rect := processor.findGameArea(img)
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

func (processor *ImageProcessorV1) findGameArea(img image.Image) image.Rectangle {
	gray := rgbaToGray(img)
	width := gray.Bounds().Dx()
	height := gray.Bounds().Dy()
	processor.logger.Info(fmt.Sprintf("process image with size %dx%d px", width, height))

	quadSideLength := 0
	quadSide := 0
	startPoint := image.Point{}

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

	endPoint := image.Point{
		startPoint.X + quadSide,
		startPoint.Y + quadSide,
	}

	processor.logger.Info(fmt.Sprintf("calculated game area with side %d px", quadSide))

	return image.Rectangle{
		Min: startPoint,
		Max: endPoint,
	}
}

// Converts the image to black and white. Used to clearly define the playing area.
func rgbaToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

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
