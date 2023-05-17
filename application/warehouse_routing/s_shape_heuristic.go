package warehouseRouting

import (
	"fmt"
	"go-algorithms/application/utils"
	"math"
	"sort"
)

func distributeSubAisle(listLoc []*Coordinate) string {

	BlockSubAsile := make([]*BlockSubAsile, 0)

	// Group all location into block of subaisles
	for _, loc := range listLoc {
		// block sub aisle 1
		if IsBelongBlockAisle(loc, 0) {
			BlockSubAsile 
			continue
		}
		// block sub aisle 2
		if IsBelongBlockAisle(loc, 1){
			subAisles[1].Loc = append(subAisles[1].Loc, loc)
			subAisles[1].LocS = append(subAisles[1].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 3
		if IsBelongBlockAisle(loc, 2) {
			subAisles[2].Loc = append(subAisles[2].Loc, loc)
			subAisles[2].LocS = append(subAisles[2].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 4
		if IsBelongBlockAisle(loc, 3) || (loc.X == 3 && loc.Y >= 1 && loc.Y <= 4) {
			subAisles[3].Loc = append(subAisles[3].Loc, loc)
			subAisles[3].LocS = append(subAisles[3].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 5
		if IsBelongBlockAisle(loc, 4) || (loc.X == 3 && loc.Y >= 5 && loc.Y <= 8) {
			subAisles[4].Loc = append(subAisles[4].Loc, loc)
			subAisles[4].LocS = append(subAisles[4].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 6
		if IsBelongBlockAisle(loc,5) {
			subAisles[5].Loc = append(subAisles[5].Loc, loc)
			subAisles[5].LocS = append(subAisles[5].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 7
		if IsBelongBlockAisle(loc, 6){
			subAisles[6].Loc = append(subAisles[6].Loc, loc)
			subAisles[6].LocS = append(subAisles[6].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 8
		if IsBelongBlockAisle(loc, 7){
			subAisles[7].Loc = append(subAisles[7].Loc, loc)
			subAisles[7].LocS = append(subAisles[7].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
		// block sub aisle 9
		if IsBelongBlockAisle(loc, 8){
			subAisles[8].Loc = append(subAisles[8].Loc, loc)
			subAisles[8].LocS = append(subAisles[8].LocS, fmt.Sprintf("{X: %v, Y: %v}", loc.X, loc.Y))
			continue
		}
	}

	// Sort the blocks in decreasing distance from the depot
	subAisles = utils.Where(subAisles, func(i *SubAsile) bool {
		return len(i.LocS) > 0
	})
	sort.Slice(subAisles, func(i, j int) bool {
		return subAisles[i].Distance > subAisles[j].Distance
	})

	//  Determine the closest (left or rightmost) subaisle with pick in the furthest block from depot
	for _, loc := range subAisles[0].Loc {

	}

	routingLoc := make([]*Coordinate, 0)
	for _, subAisle := range subAisles {
		routingLoc = append(routingLoc, subAisle.Loc...)
	}

	return utils.PrettyPrint(routingLoc)
}

// FindPickingRouteSShape is a function to find route for picking (s-shape heuristic)
func FindPickingRouteSShape() {
	listLoc := []*Coordinate{
		{X: 1, Y: 6},
		{X: 1, Y: 11},
		{X: 1, Y: 15},
		{X: 2, Y: 1},
		{X: 2, Y: 6},
		{X: 2, Y: 7},
		{X: 4, Y: 3},
		{X: 4, Y: 11},
	}

	fmt.Println(distributeSubAisle(listLoc))
}

func CalculateDistanceFromDepot(loc *Coordinate) float64 {
	// square_root((x2-x1)^2 + (y2-y1)^2)
	return math.Sqrt(math.Pow(float64(loc.X), 2) + math.Pow(float64(loc.Y), 2))
}

func IsBelongBlockAisle(loc *Coordinate, blockIndex int) bool {
	listLoc := WAREHOUSE_LAYOUT[0].Loc
	for _, l := range listLoc {
		if l.X == loc.X && l.Y == loc.Y {
			return true
		}
	}
	return false
}
