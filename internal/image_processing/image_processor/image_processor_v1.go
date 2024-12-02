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
	"net/http"
	"os"
	"strings"
	"time"
)

const (

	// The amount by which the scanning batch size will be reduced.
	// Used to exclude borders of adjacent cells.
	quadSideDelta = 12
	// Cell count
	cellCount = 9
	threshold = 62000
	fileKey   = "file"
)

var possibleNumberString = map[string]string{
	"1": "0",
	"2": "1",
	"3": "2",
	"4": "3",
	"5": "4",
	"6": "5",
	"7": "6",
	"8": "7",
	"9": "8",
}

type ImageProcessorV1 struct {
	Logger *slog.Logger
}

func NewImageProcessorV1(logger *slog.Logger) *ImageProcessorV1 {
	return &ImageProcessorV1{
		Logger: logger,
	}
}

func (processor *ImageProcessorV1) GetBattlefieldFromFile(file *os.File) string {
	img, _, err := image.Decode(file)

	if err != nil {
		return ""
	}

	rect := processor.findGameArea(img)
	croppedImage := imaging.Crop(img, rect)
	croppedImage = imaging.AdjustContrast(croppedImage, 80)
	croppedImage = imaging.Grayscale(croppedImage)

	return cutImageToParts(croppedImage, rect)
}

func (processor *ImageProcessorV1) GetBattlefield(r *http.Request) string {
	logger := processor.Logger
	logger.Info("start process image")
	err := r.ParseMultipartForm(32 << 15)

	if err != nil {
		logger.Error("failed to parse data from request")
		return ""
	}

	file, _, err := r.FormFile(fileKey)
	img, _, err := image.Decode(file)
	if err != nil {
		return ""
	}

	rect := processor.findGameArea(img)
	croppedImage := imaging.Crop(img, rect)
	croppedImage = imaging.AdjustContrast(croppedImage, 80)
	croppedImage = imaging.Grayscale(croppedImage)

	return cutImageToParts(croppedImage, rect)
}

func cutImageToParts(croppedImage *image.NRGBA, rect image.Rectangle) string {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	now := time.Now()
	logger.Info("start cutting image")

	var result = strings.Builder{}

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

			val, ok := possibleNumberString[out]
			if ok {
				result.WriteString(val)
			} else {
				result.WriteString(".")
			}

			counter++
			coordinateX++
		}
		coordinateX = 0
		coordinateY++
	}

	logger.Info(fmt.Sprintf("finish cutting image. Time: %d ms", time.Now().Sub(now).Milliseconds()))
	return result.String()
}

func (processor *ImageProcessorV1) findGameArea(img image.Image) image.Rectangle {
	gray := rgbaToGray(img)
	width := gray.Bounds().Dx()
	height := gray.Bounds().Dy()
	processor.Logger.Info(fmt.Sprintf("process image with size %dx%d px", width, height))

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

	processor.Logger.Info(fmt.Sprintf("calculated game area with side %d px", quadSide))

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
