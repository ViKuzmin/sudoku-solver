package script_creator

import (
	"strconv"
	"strings"
)

type ScriptCreator struct {
}

func NewScriptCreator() *ScriptCreator {
	return &ScriptCreator{}
}

// Координаты точек для redmi note 10s
func (creator *ScriptCreator) GetScript(data [][]int) string {
	pixelX := 90
	pixelY := 505
	stepX := 112
	stepY := 114

	pixelAnswerX := 77
	pixelAnswerY := 2011
	pixelAnswerXStep := 116

	var sb strings.Builder

	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[x]); y++ {
			number := data[x][y]
			sb.WriteString(creator.Tap(pixelX, pixelY))

			switch number {
			case 1:
				sb.WriteString(creator.Tap(pixelAnswerX, pixelAnswerY))
			case 2:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep, pixelAnswerY))
			case 3:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*2, pixelAnswerY))
			case 4:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*3, pixelAnswerY))
			case 5:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*4, pixelAnswerY))
			case 6:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*5, pixelAnswerY))
			case 7:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*6, pixelAnswerY))
			case 8:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*7, pixelAnswerY))
			case 9:
				sb.WriteString(creator.Tap(pixelAnswerX+pixelAnswerXStep*8, pixelAnswerY))
			}

			pixelX += stepX
		}
		pixelX = 130
		pixelY += stepY
	}

	return sb.String()
}

func (creator *ScriptCreator) Tap(x, y int) string {
	return "input tap " + strconv.Itoa(x) + " " + strconv.Itoa(y) + "; "
}

func (creator ScriptCreator) Sleep(time int) string {
	return "sleep " + strconv.Itoa(time) + "; "
}
