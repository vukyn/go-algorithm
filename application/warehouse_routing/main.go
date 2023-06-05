package warehouseRouting

import (
	"fmt"
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
	"go-algorithms/application/warehouse_routing/test"
	"image/color"
	"sort"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	IsLock     bool
	IsMoving   bool
	Stage      int
	CurrentLoc struct {
		X float32
		Y float32
	}
	DepotLoc           *models.Coordinate
	PickerLoc          *models.Coordinate
	NextPickLoc        *models.Coordinate
	ListLoc            []*models.Coordinate
	ListWalkLoc        []*models.Coordinate
	ListRemainWalkLoc  []*models.Coordinate
	ListWallLoc        []*models.Coordinate
	ListPickLoc        []*models.Coordinate
	ListCurrentPickLoc []*models.Coordinate
	ListPickableLoc    []*models.Coordinate
	ListWalkingLoc     []*models.Coordinate
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.IsLock = !g.IsLock
	}

	if !g.IsLock {
		if len(g.ListPickLoc) == 0 {
			g.ListRemainWalkLoc = g.ListWalkLoc
		}
		if g.Stage == 1 {
			if !g.IsMoving {
				// call greedy heuristic
				// nextPick := greedy.CallGreedyHeuristic(g.PickerLoc, g.ListPickLoc, g.ListRemainWalkLoc)

				// call test
				if len(g.ListCurrentPickLoc) == 0 {
					g.ListCurrentPickLoc, g.Stage = test.GetNextPickLocation(g.ListPickLoc, g.ListWalkLoc, g.PickerLoc, g.Stage)
					fmt.Printf("Next pick: %v\n", utils.PrettyPrint(g.ListCurrentPickLoc))
				}

				// Move picker to nearest possible walk location
				if g.Stage == 1 {
					g.IsMoving = true
					if len(g.ListCurrentPickLoc) > 0 {
						g.NextPickLoc = g.ListCurrentPickLoc[0]
						g.ListCurrentPickLoc = g.ListCurrentPickLoc[1:]
					} else {
						g.NextPickLoc = g.DepotLoc
					}
				}
			} else {
				// moving
				if g.PickerLoc != g.NextPickLoc {
					if strconv.Itoa(g.PickerLoc.X) == fmt.Sprintf("%.f", g.CurrentLoc.X) && strconv.Itoa(g.PickerLoc.Y) == fmt.Sprintf("%.f", g.CurrentLoc.Y) {
						var isPicked bool
						g.ListPickLoc, isPicked = helper.PickItem(g.ListPickLoc, g.ListPickableLoc, g.PickerLoc)
						if isPicked {
							g.IsMoving = false
						} else {
							// listPossibleLoc := helper.GetPossibleLocation(g.PickerLoc, g.ListRemainWalkLoc)
							// Calulate distance from possible walk location to nearest pick location (euclidean distance)
							// minDistance := 99999
							// for _, possibleLoc := range listPossibleLoc {
							// 	distance := helper.CalculateEuclideanDistance(possibleLoc, g.NextPickLoc)
							// 	if distance < minDistance {
							// 		minDistance = distance
							// 		g.PickerLoc = possibleLoc
							// 	}
							// }
							// g.ListRemainWalkLoc = utils.Where(g.ListRemainWalkLoc, func(loc *models.Coordinate) bool {
							// 	return loc.X != g.PickerLoc.X || loc.Y != g.PickerLoc.Y
							// })
						}
					}
					g.moving()
				}
			}
		}
		if g.Stage != 1 {
			if !g.IsMoving {
				// call greedy heuristic
				// nextPick := greedy.CallGreedyHeuristic(g.PickerLoc, g.ListPickLoc, g.ListRemainWalkLoc)

				// call test
				if len(g.ListCurrentPickLoc) == 0 {
					g.ListCurrentPickLoc, g.Stage = test.GetNextPickLocation(g.ListPickLoc, g.ListWalkLoc, g.PickerLoc, g.Stage)
					fmt.Printf("Next pick: %v\n", utils.PrettyPrint(g.ListCurrentPickLoc))
				}

				// Move picker to nearest possible walk location
				g.IsMoving = true
				if len(g.ListCurrentPickLoc) > 0 {
					min := 99999
					for _, loc := range g.ListCurrentPickLoc {
						routes := helper.CalculateDfsDistance(g.PickerLoc, loc, g.ListWalkLoc)
						sort.Slice(routes, func(i, j int) bool {
							return routes[i].Distance < routes[j].Distance
						})
						if routes[0].Distance < min {
							min = routes[0].Distance
							g.ListWalkingLoc = routes[0].ListVisitedLoc
						}
					}
					fmt.Printf("ListCurrentPickLoc: %v\n", utils.PrettyPrint(g.ListCurrentPickLoc))
					fmt.Printf("Walking loc: %v\n", utils.PrettyPrint(g.ListWalkingLoc))
					// g.NextPickLoc = g.ListCurrentPickLoc[0]
					g.ListCurrentPickLoc = utils.Where(g.ListCurrentPickLoc, func(loc *models.Coordinate) bool {
						return loc.X != g.ListWalkingLoc[len(g.ListWalkingLoc)-1].X || loc.Y != g.ListWalkingLoc[len(g.ListWalkingLoc)-1].Y
					})
				} else {
					g.NextPickLoc = g.DepotLoc
				}
			} else {
				// moving
				if strconv.Itoa(g.PickerLoc.X) == fmt.Sprintf("%.f", g.CurrentLoc.X) && strconv.Itoa(g.PickerLoc.Y) == fmt.Sprintf("%.f", g.CurrentLoc.Y) {
					var isPicked bool
					g.ListPickLoc, isPicked = helper.PickItem(g.ListPickLoc, g.ListPickableLoc, g.PickerLoc)
					if isPicked {
						g.IsMoving = false
					} else {
						if len(g.ListWalkingLoc) > 0 {
							g.PickerLoc = g.ListWalkingLoc[0]
							g.ListWalkingLoc = g.ListWalkingLoc[1:]
						}
					}
				}
				g.moving()
			}
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
	g.IsLock = false
	g.Stage = 1
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
	// g.ListPickLoc = helper.GenerateRandomPickLocation(8, g.ListWallLoc)

	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 15})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 14})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 7})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 0, Y: 2})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 2, Y: 1})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 3, Y: 1})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 3, Y: 5})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 3, Y: 8})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 5, Y: 2})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 6, Y: 2})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 6, Y: 9})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 5, Y: 16})
	g.ListPickLoc = append(g.ListPickLoc, &models.Coordinate{X: 6, Y: 16})

	// Set picker init location
	g.PickerLoc = g.DepotLoc
	g.CurrentLoc.X = float32(g.PickerLoc.X)
	g.CurrentLoc.Y = float32(g.PickerLoc.Y)
}

func (g *Game) moving() {
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
