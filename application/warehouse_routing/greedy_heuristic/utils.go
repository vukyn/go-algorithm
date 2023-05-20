package greedyheuristic

import (
	"bufio"
	warehouseRouting "go-algorithms/application/warehouse_routing"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func readMap() []string {
	file, err := os.Open(MAP_PATH)
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

func drawSquarePattern(screen *ebiten.Image, clr color.Color, width int, filled bool) {
	for i := 0; i < SCREEN_HEIGHT; i += width {
		for j := 0; j < SCREEN_WIDTH; j += width {
			if filled {
				vector.DrawFilledRect(screen, float32(j), float32(i), float32(width), float32(width), clr, true)
			} else {
				vector.StrokeRect(screen, float32(j), float32(i), float32(width), float32(width), 1, clr, true)
			}
		}
	}
}

func toCoordinate(x, y int) *warehouseRouting.Coordinate {
	return &warehouseRouting.Coordinate{X: x, Y: y}
}
