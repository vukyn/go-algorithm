package warehouseRouting

import (
	"fmt"
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/greedy"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	IsMoving   bool
	CurrentLoc struct {
		X float32
		Y float32
	}
	DepotLoc          *models.Coordinate
	PickerLoc         *models.Coordinate
	ListLoc           []*models.Coordinate
	ListWalkLoc       []*models.Coordinate
	ListRemainWalkLoc []*models.Coordinate
	ListWallLoc       []*models.Coordinate
	ListPickLoc       []*models.Coordinate
	ListPickableLoc   []*models.Coordinate
}

func (g *Game) Update() error {

	if !g.IsMoving {

		// call greedy heuristic
		minPossibleLoc := greedy.CallGreedyHeuristic(g.PickerLoc, g.ListPickLoc, g.ListRemainWalkLoc)

		// Move picker to nearest possible walk location
		g.IsMoving = true
		if len(g.ListPickLoc) > 0 {
			g.PickerLoc = minPossibleLoc
		} else {
			g.PickerLoc = g.DepotLoc
		}
		g.ListRemainWalkLoc = utils.Where(g.ListRemainWalkLoc, func(loc *models.Coordinate) bool {
			return loc.X != minPossibleLoc.X || loc.Y != minPossibleLoc.Y
		})
	} else {
		if strconv.Itoa(g.PickerLoc.X) == fmt.Sprintf("%.f", g.CurrentLoc.X) && strconv.Itoa(g.PickerLoc.Y) == fmt.Sprintf("%.f", g.CurrentLoc.Y) {
			g.IsMoving = false
			// Pick item
			// var isPicked bool
			g.ListPickLoc, _ = helper.PickItem(g.ListPickLoc, g.ListPickableLoc, g.PickerLoc)

			//Comment if don't want to repeat walk path
			// if isPicked {
			// 	g.ListRemainWalkLoc = g.ListWalkLoc
			// }
		}
		// Move left
		if float32(g.PickerLoc.X) < g.CurrentLoc.X {
			g.CurrentLoc.X -= 0.1
		}
		// Move right
		if float32(g.PickerLoc.X) > g.CurrentLoc.X {
			g.CurrentLoc.X += 0.1
		}
		// Move up
		if float32(g.PickerLoc.Y) < g.CurrentLoc.Y {
			g.CurrentLoc.Y -= 0.1
		}
		// Move down
		if float32(g.PickerLoc.Y) > g.CurrentLoc.Y {
			g.CurrentLoc.Y += 0.1
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// // Draw background
	// drawSquarePattern(screen, color.RGBA{96, 96, 96, 255}, constants.RECT_WIDTH, true)

	offsetX := 40

	// Draw wall
	for _, loc := range g.ListWallLoc {
		vector.DrawFilledRect(screen, float32(offsetX+loc.X*constants.RECT_WIDTH), float32(loc.Y*constants.RECT_HEIGHT), constants.RECT_WIDTH, constants.RECT_HEIGHT, color.RGBA{255, 0, 0, 255}, true)
	}

	// Draw walk path
	// for _, loc := range g.ListWalkLoc {
	// 	vector.DrawFilledRect(screen, float32(offsetX+loc.X*constants.RECT_WIDTH), float32(loc.Y*constants.RECT_HEIGHT), constants.RECT_WIDTH, constants.RECT_HEIGHT, color.RGBA{96, 96, 96, 255}, true)
	// }

	// Draw pick location
	for _, loc := range g.ListPickLoc {
		helper.DrawCrossX(screen, color.RGBA{255, 255, 255, 255}, offsetX+loc.X*constants.RECT_WIDTH, loc.Y*constants.RECT_HEIGHT, constants.RECT_WIDTH, constants.RECT_HEIGHT)
	}

	// Draw depot
	vector.DrawFilledRect(screen, float32(offsetX+g.DepotLoc.X*constants.RECT_WIDTH), float32(g.DepotLoc.Y*constants.RECT_HEIGHT), constants.RECT_WIDTH, constants.RECT_HEIGHT, color.RGBA{0, 128, 255, 255}, true)

	// Draw current picker
	vector.DrawFilledRect(screen, float32(offsetX)+g.CurrentLoc.X*constants.RECT_WIDTH, g.CurrentLoc.Y*constants.RECT_HEIGHT, constants.RECT_WIDTH, constants.RECT_HEIGHT, color.RGBA{0, 255, 153, 255}, true)

	// Draw list location coordinate
	for _, loc := range g.ListLoc {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(("x: %d\ny: %d"), loc.X, loc.Y), offsetX+constants.RECT_WIDTH*(loc.X)+5, constants.RECT_HEIGHT*(loc.Y)+3)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Depot: x: %d, y: %d", g.DepotLoc.X, g.DepotLoc.Y), 0, constants.SCREEN_HEIGHT-60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Picker: x: %d, y: %d", g.PickerLoc.X, g.PickerLoc.Y), 0, constants.SCREEN_HEIGHT-40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current: x: %f, y: %f", g.CurrentLoc.X, g.CurrentLoc.Y), 0, constants.SCREEN_HEIGHT-20)

	// Draw overlay white square pattern
	helper.DrawSquarePattern(screen, color.RGBA{255, 255, 255, 255}, constants.RECT_WIDTH, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT
}

func Run(screen *ebiten.Image) {
	ebiten.SetWindowSize(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Greedy Heuristic")
	g := newGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func newGame() ebiten.Game {
	g := &Game{}
	g.init()
	g.IsMoving = false
	return g
}

func (g *Game) init() {
	// Init warehouse coordinate
	lines := helper.ReadMap()
	for y, line := range lines {
		for x, char := range strings.Split(line, "-") {
			g.ListLoc = append(g.ListLoc, helper.ToCoordinate(x, y))
			if char == strconv.Itoa(constants.WALL) {
				g.ListWallLoc = append(g.ListWallLoc, helper.ToCoordinate(x, y))
			}
			if char == strconv.Itoa(constants.WALK) || char == strconv.Itoa(constants.PICKABLE) {
				g.ListWalkLoc = append(g.ListWalkLoc, helper.ToCoordinate(x, y))
				g.ListRemainWalkLoc = append(g.ListRemainWalkLoc, helper.ToCoordinate(x, y))
			}
			if char == strconv.Itoa(constants.PICKABLE) {
				g.ListPickableLoc = append(g.ListPickableLoc, helper.ToCoordinate(x, y))
			}
			if char == strconv.Itoa(constants.DEPOT) {
				g.DepotLoc = helper.ToCoordinate(x, y)
			}
		}
	}

	// Random pick location
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	// for i := 0; i < 10; i++ {
	// 	index := r1.Intn(len(g.ListWallLoc))
	// 	if len(g.ListPickLoc) > 0 {
	// 		if loc := utils.Find(g.ListPickLoc, func(loc *warehouseRouting.Coordinate) bool {
	// 			return loc.X == g.ListWallLoc[index].X && loc.Y == g.ListWallLoc[index].Y
	// 		}); loc == nil {
	// 			g.ListPickLoc = append(g.ListPickLoc, g.ListWallLoc[index])
	// 		}
	// 	} else {
	// 		g.ListPickLoc = append(g.ListPickLoc, g.ListWallLoc[index])
	// 	}
	// }

	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 15})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 14})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 7})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 2})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 2, Y: 1})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 6, Y: 2})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 5, Y: 16})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 6, Y: 16})

	// Set picker init location
	g.PickerLoc = g.DepotLoc
	g.CurrentLoc.X = float32(g.PickerLoc.X)
	g.CurrentLoc.Y = float32(g.PickerLoc.Y)
}
