package main

import (
	"fmt"
	"image"
	"image/color"
	"ocr-test/internal/image_processing/image_processor"
	"ocr-test/internal/image_processing/sudoku_solver"
)

func main() {
	processor := image_processor.NewImageProcessorV1()
	solver := sudoku_solver.NewSolver()

	data := processor.ProcessImage("123.jpg")

	solve := solver.GetScript(data)

	fmt.Println(solve)

	//client := gosseract.NewClient()

	//client.SetImageFromBytes()
	//client.Text()
	//
	//defer client.Close()
	//client.SetImage("images/5.jpg")
	//text, _ := client.Text()
	//fmt.Println(text)

	// Open the image file
	//file, err := os.Open("images/9.jpg")
	//if err != nil {
	//	log.Fatalf("failed to open image: %v", err)
	//}
	//defer file.Close()
	//
	//// Decode the image
	//img, _, err := image.Decode(file)
	//if err != nil {
	//	log.Fatalf("failed to decode image: %v", err)
	//}
	//
	//// Invert the colors of the image
	//invertedImg := invertColor(img)
	//
	//// Save the inverted image to a new file
	//err = imaging.Save(invertedImg, "inverted.jpg")
	//if err != nil {
	//	log.Fatalf("failed to save image: %v", err)
	//}

}

func invertColor(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	dst := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			r, g, b, a := oldColor.RGBA()

			invertedColor := color.NRGBA{
				R: 255 - uint8(r>>8),
				G: 255 - uint8(g>>8),
				B: 255 - uint8(b>>8),
				A: uint8(a >> 8),
			}

			dst.SetNRGBA(x, y, invertedColor)
		}
	}

	return dst
}
