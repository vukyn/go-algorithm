package greedyheuristic

import (
	"bufio"
	"go-algorithms/application/utils"
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

func drawCrossX(screen *ebiten.Image, clr color.Color, x, y, width, height int) {
	vector.StrokeLine(screen, float32(x), float32(y), float32(x+width), float32(y+height), 1, clr, true)
	vector.StrokeLine(screen, float32(x+width), float32(y), float32(x), float32(y+height), 1, clr, true)
}

func toCoordinate(x, y int) *warehouseRouting.Coordinate {
	return &warehouseRouting.Coordinate{X: x, Y: y}
}

func getPossibleLocation(currentLoc *warehouseRouting.Coordinate, locations []*warehouseRouting.Coordinate) []*warehouseRouting.Coordinate {
	listPossibleLoc := make([]*warehouseRouting.Coordinate, 0)
	if northLoc := utils.Find(locations, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == currentLoc.X && loc.Y == currentLoc.Y-1
	}); northLoc != nil {
		listPossibleLoc = append(listPossibleLoc, northLoc)
	}
	if southLoc := utils.Find(locations, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == currentLoc.X && loc.Y == currentLoc.Y+1
	}); southLoc != nil {
		listPossibleLoc = append(listPossibleLoc, southLoc)
	}
	if westLoc := utils.Find(locations, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == currentLoc.X-1 && loc.Y == currentLoc.Y
	}); westLoc != nil {
		listPossibleLoc = append(listPossibleLoc, westLoc)
	}
	if eastLoc := utils.Find(locations, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == currentLoc.X+1 && loc.Y == currentLoc.Y
	}); eastLoc != nil {
		listPossibleLoc = append(listPossibleLoc, eastLoc)
	}
	return listPossibleLoc
}

func pickItem(listPickLoc []*warehouseRouting.Coordinate, pickerLoc *warehouseRouting.Coordinate) ([]*warehouseRouting.Coordinate, bool) {
	isPicked := false
	if northLoc := utils.IndexOf(listPickLoc, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == pickerLoc.X && loc.Y == pickerLoc.Y-1
	}); northLoc != -1 {
		isPicked = true
		listPickLoc = append(listPickLoc[:northLoc], listPickLoc[northLoc+1:]...)
	}
	if southLoc := utils.IndexOf(listPickLoc, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == pickerLoc.X && loc.Y == pickerLoc.Y+1
	}); southLoc != -1 {
		isPicked = true
		listPickLoc = append(listPickLoc[:southLoc], listPickLoc[southLoc+1:]...)
	}
	if westLoc := utils.IndexOf(listPickLoc, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == pickerLoc.X-1 && loc.Y == pickerLoc.Y
	}); westLoc != -1 {
		isPicked = true
		listPickLoc = append(listPickLoc[:westLoc], listPickLoc[westLoc+1:]...)
	}
	if eastLoc := utils.IndexOf(listPickLoc, func(loc *warehouseRouting.Coordinate) bool {
		return loc.X == pickerLoc.X+1 && loc.Y == pickerLoc.Y
	}); eastLoc != -1 {
		isPicked = true
		listPickLoc = append(listPickLoc[:eastLoc], listPickLoc[eastLoc+1:]...)
	}
	return listPickLoc, isPicked
}
