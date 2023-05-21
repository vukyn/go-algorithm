package greedyheuristic

import (
	"fmt"
	"go-algorithms/application/utils"
	warehouseRouting "go-algorithms/application/warehouse_routing"
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	isMoving   bool
	currentLoc struct {
		X float32
		Y float32
	}
	depotLoc          *warehouseRouting.Coordinate
	pickerLoc         *warehouseRouting.Coordinate
	listLoc           []*warehouseRouting.Coordinate
	listWalkLoc       []*warehouseRouting.Coordinate
	listRemainWalkLoc []*warehouseRouting.Coordinate
	listWallLoc       []*warehouseRouting.Coordinate
	listPickLoc       []*warehouseRouting.Coordinate
}

func (g *Game) Update() error {

	if !g.isMoving {
		// Get possible walk location
		listPossibleLoc := getPossibleLocation(g.pickerLoc, g.listRemainWalkLoc)

		// Calulate distance from possible walk location to nearest pick location (manhattan distance)
		minDistance := 99999
		minPossibleLoc := &warehouseRouting.Coordinate{}
		for _, possibleLoc := range listPossibleLoc {
			for _, pickLoc := range g.listPickLoc {
				distance := warehouseRouting.CalculateEuclideanDistance(possibleLoc, pickLoc)
				if distance < minDistance {
					minDistance = distance
					minPossibleLoc = possibleLoc
				}
			}
		}

		// Move picker to nearest possible walk location
		g.isMoving = true
		if len(g.listPickLoc) > 0 {
			g.pickerLoc = minPossibleLoc
		} else {
			g.pickerLoc = g.depotLoc
		}
		g.listRemainWalkLoc = utils.Where(g.listRemainWalkLoc, func(loc *warehouseRouting.Coordinate) bool {
			return loc.X != minPossibleLoc.X || loc.Y != minPossibleLoc.Y
		})
	} else {
		if strconv.Itoa(g.pickerLoc.X) == fmt.Sprintf("%.f", g.currentLoc.X) && strconv.Itoa(g.pickerLoc.Y) == fmt.Sprintf("%.f", g.currentLoc.Y) {
			g.isMoving = false
			// Pick item
			// var isPicked bool
			g.listPickLoc, _ = pickItem(g.listPickLoc, g.pickerLoc)

			//Comment if don't want to repeat walk path
			// if isPicked {
			// 	g.listRemainWalkLoc = g.listWalkLoc
			// }
		}
		// Move left
		if float32(g.pickerLoc.X) < g.currentLoc.X {
			g.currentLoc.X -= 0.1
		}
		// Move right
		if float32(g.pickerLoc.X) > g.currentLoc.X {
			g.currentLoc.X += 0.1
		}
		// Move up
		if float32(g.pickerLoc.Y) < g.currentLoc.Y {
			g.currentLoc.Y -= 0.1
		}
		// Move down
		if float32(g.pickerLoc.Y) > g.currentLoc.Y {
			g.currentLoc.Y += 0.1
		}
	}

	// fmt.Printf("Picker move to x: %d, y: %d\n", g.pickerLoc.X, g.pickerLoc.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// // Draw background
	// drawSquarePattern(screen, color.RGBA{96, 96, 96, 255}, RECT_WIDTH, true)

	offsetX := 40

	// Draw wall
	for _, loc := range g.listWallLoc {
		vector.DrawFilledRect(screen, float32(offsetX+loc.X*RECT_WIDTH), float32(loc.Y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{255, 0, 0, 255}, true)
	}

	// Draw walk path
	// for _, loc := range g.listWalkLoc {
	// 	vector.DrawFilledRect(screen, float32(offsetX+loc.X*RECT_WIDTH), float32(loc.Y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{96, 96, 96, 255}, true)
	// }

	// Draw pick location
	for _, loc := range g.listPickLoc {
		drawCrossX(screen, color.RGBA{255, 255, 255, 255}, offsetX+loc.X*RECT_WIDTH, loc.Y*RECT_HEIGHT, RECT_WIDTH, RECT_HEIGHT)
	}

	// Draw depot
	vector.DrawFilledRect(screen, float32(offsetX+g.depotLoc.X*RECT_WIDTH), float32(g.depotLoc.Y*RECT_HEIGHT), RECT_WIDTH, RECT_HEIGHT, color.RGBA{0, 128, 255, 255}, true)

	// Draw current picker
	vector.DrawFilledRect(screen, float32(offsetX)+g.currentLoc.X*RECT_WIDTH, g.currentLoc.Y*RECT_HEIGHT, RECT_WIDTH, RECT_HEIGHT, color.RGBA{0, 255, 153, 255}, true)

	// Draw list location coordinate
	for _, loc := range g.listLoc {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(("x: %d\ny: %d"), loc.X, loc.Y), offsetX+RECT_WIDTH*(loc.X)+5, RECT_HEIGHT*(loc.Y)+3)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Depot: x: %d, y: %d", g.depotLoc.X, g.depotLoc.Y), 0, SCREEN_HEIGHT-60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Picker: x: %d, y: %d", g.pickerLoc.X, g.pickerLoc.Y), 0, SCREEN_HEIGHT-40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current: x: %f, y: %f", g.currentLoc.X, g.currentLoc.Y), 0, SCREEN_HEIGHT-20)

	// Draw overlay white square pattern
	drawSquarePattern(screen, color.RGBA{255, 255, 255, 255}, RECT_WIDTH, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func Run(screen *ebiten.Image) {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Greedy Heuristic")
	g := newGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func newGame() ebiten.Game {
	g := &Game{}
	g.init()
	g.isMoving = false
	return g
}

func (g *Game) init() {
	// Init warehouse coordinate
	lines := readMap()
	for y, line := range lines {
		for x, char := range strings.Split(line, "-") {
			g.listLoc = append(g.listLoc, toCoordinate(x, y))
			if char == strconv.Itoa(WALL) {
				g.listWallLoc = append(g.listWallLoc, toCoordinate(x, y))
				continue
			}
			if char == strconv.Itoa(WALK) {
				g.listWalkLoc = append(g.listWalkLoc, toCoordinate(x, y))
				g.listRemainWalkLoc = append(g.listWalkLoc, toCoordinate(x, y))
				continue
			}
			if char == strconv.Itoa(DEPOT) {
				g.depotLoc = toCoordinate(x, y)
				continue
			}
		}
	}

	// Random pick location
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	// for i := 0; i < 10; i++ {
	// 	index := r1.Intn(len(g.listWallLoc))
	// 	if len(g.listPickLoc) > 0 {
	// 		if loc := utils.Find(g.listPickLoc, func(loc *warehouseRouting.Coordinate) bool {
	// 			return loc.X == g.listWallLoc[index].X && loc.Y == g.listWallLoc[index].Y
	// 		}); loc == nil {
	// 			g.listPickLoc = append(g.listPickLoc, g.listWallLoc[index])
	// 		}
	// 	} else {
	// 		g.listPickLoc = append(g.listPickLoc, g.listWallLoc[index])
	// 	}
	// }

	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 0, Y: 15})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 0, Y: 14})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 0, Y: 7})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 0, Y: 2})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 2, Y: 1})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 6, Y: 2})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 5, Y: 16})
	g.listPickLoc = append(g.listPickLoc, &warehouseRouting.Coordinate{X: 6, Y: 16})

	// Set picker init location
	g.pickerLoc = g.depotLoc
	g.currentLoc.X = float32(g.pickerLoc.X)
	g.currentLoc.Y = float32(g.pickerLoc.Y)

	// fmt.Printf("Wall: %v", utils.PrettyPrint(g.listPickLoc))
}
