package helper

import (
	"bufio"
	"go-algorithms/application/warehouse_routing/constants"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func ReadMap() []string {
	file, err := os.Open(constants.MAP_PATH)
	if err != nil {
		log.Fatal("Cannot open file")
	}
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func DrawSquarePattern(screen *ebiten.Image, clr color.Color, width int, filled bool) {
	for i := 0; i < constants.SCREEN_HEIGHT; i += width {
		for j := 0; j < constants.SCREEN_WIDTH; j += width {
			if filled {
				vector.DrawFilledRect(screen, float32(j), float32(i), float32(width), float32(width), clr, true)
			} else {
				vector.StrokeRect(screen, float32(j), float32(i), float32(width), float32(width), 1, clr, true)
			}
		}
	}
}

func DrawCrossX(screen *ebiten.Image, clr color.Color, x, y, width, height int) {
	vector.StrokeLine(screen, float32(x), float32(y), float32(x+width), float32(y+height), 1, clr, true)
	vector.StrokeLine(screen, float32(x+width), float32(y), float32(x), float32(y+height), 1, clr, true)
}
