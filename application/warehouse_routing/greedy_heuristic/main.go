package greedyheuristic

import (
	"fmt"
	warehouseRouting "go-algorithms/application/warehouse_routing"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	depotLoc    *warehouseRouting.Coordinate
	listLoc     []*warehouseRouting.Coordinate
	listWalkLoc []*warehouseRouting.Coordinate
	listWallLoc []*warehouseRouting.Coordinate
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	depotLoc := &warehouseRouting.Coordinate{}
	listLoc := make([]*warehouseRouting.Coordinate, 0)
	listWallLoc := make([]*warehouseRouting.Coordinate, 0)
	listWalkLoc := make([]*warehouseRouting.Coordinate, 0)

	// Draw background
	// drawSquarePattern(screen, color.RGBA{96, 96, 96, 255}, RECT_WIDTH, true)

	// Draw warehouse
	lines := readMap()
	for y, line := range lines {
		offset := 40
		for x, char := range strings.Split(line, "-") {
			listLoc = append(listLoc, toCoordinate(x, y))
			if char == WALL {
				listWallLoc = append(listWallLoc, toCoordinate(x, y))
				vector.DrawFilledRect(screen, float32(offset), float32(y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{255, 0, 0, 255}, true)
				offset += RECT_WIDTH
				continue
			}
			if char == WALK {
				listWalkLoc = append(listWalkLoc, toCoordinate(x, y))
				vector.DrawFilledRect(screen, float32(offset), float32(y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{51, 255, 51, 255}, true)
				offset += RECT_WIDTH
				continue
			}
			if char == DEPOT {
				depotLoc = toCoordinate(x, y)
				vector.DrawFilledRect(screen, float32(offset), float32(y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{0, 128, 255, 255}, true)
				offset += RECT_WIDTH
				continue
			}
		}
	}
	

	offset := 40
	for _, loc := range listLoc {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(("x: %d\ny: %d"), loc.X, loc.Y), offset+RECT_WIDTH*(loc.X)+5, RECT_HEIGHT*(loc.Y)+3)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Depot: x: %d, y: %d", depotLoc.X, depotLoc.Y), 0, SCREEN_HEIGHT-20)
	// Draw overlay white square pattern
	drawSquarePattern(screen, color.RGBA{255, 255, 255, 255}, RECT_WIDTH, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func Run(screen *ebiten.Image) {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Greedy Heuristic")
	g := &Game{}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
