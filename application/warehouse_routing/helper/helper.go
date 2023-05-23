package helper

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/models"
	"math"
)

func CalculateEuclideanDistance(loc1, loc2 *models.Coordinate) int {
	return int(math.Sqrt(math.Pow(float64(loc2.X-loc1.X), 2) + math.Pow(float64(loc2.Y-loc1.Y), 2)))
}

func IsBelongBlockAisle(layout []*models.BlockSubAsile, loc *models.Coordinate, blockIndex int) bool {
	listLoc := layout[blockIndex]
	for _, l := range listLoc.Locs {
		if l.X == loc.X && l.Y == loc.Y {
			return true
		}
	}
	return false
}

func IsBelongSubAisle(loc *models.Coordinate) int {
	// Sub aisle 1
	for _, l := range constants.SUB_AISLE_1 {
		if l.X == loc.X && l.Y == loc.Y {
			return 1
		}
	}
	// Sub aisle 2
	for _, l := range constants.SUB_AISLE_2 {
		if l.X == loc.X && l.Y == loc.Y {
			return 2
		}
	}
	// Sub aisle 3
	for _, l := range constants.SUB_AISLE_3 {
		if l.X == loc.X && l.Y == loc.Y {
			return 3
		}
	}
	// Sub aisle 4
	for _, l := range constants.SUB_AISLE_4 {
		if l.X == loc.X && l.Y == loc.Y {
			return 4
		}
	}
	// Sub aisle 5
	for _, l := range constants.SUB_AISLE_5 {
		if l.X == loc.X && l.Y == loc.Y {
			return 5
		}
	}
	// Sub aisle 6
	for _, l := range constants.SUB_AISLE_6 {
		if l.X == loc.X && l.Y == loc.Y {
			return 6
		}
	}
	// Sub aisle 7
	for _, l := range constants.SUB_AISLE_8 {
		if l.X == loc.X && l.Y == loc.Y {
			return 7
		}
	}
	// Sub aisle 8
	for _, l := range constants.SUB_AISLE_8 {
		if l.X == loc.X && l.Y == loc.Y {
			return 8
		}
	}
	// Sub aisle 9
	for _, l := range constants.SUB_AISLE_9 {
		if l.X == loc.X && l.Y == loc.Y {
			return 9
		}
	}
	// Sub aisle 10
	for _, l := range constants.SUB_AISLE_10 {
		if l.X == loc.X && l.Y == loc.Y {
			return 10
		}
	}
	// Sub aisle 11
	for _, l := range constants.SUB_AISLE_11 {
		if l.X == loc.X && l.Y == loc.Y {
			return 11
		}
	}
	// Sub aisle 12
	for _, l := range constants.SUB_AISLE_12 {
		if l.X == loc.X && l.Y == loc.Y {
			return 12
		}
	}
	// Sub aisle 13
	for _, l := range constants.SUB_AISLE_13 {
		if l.X == loc.X && l.Y == loc.Y {
			return 13
		}
	}
	// Sub aisle 14
	for _, l := range constants.SUB_AISLE_14 {
		if l.X == loc.X && l.Y == loc.Y {
			return 14
		}
	}
	// Sub aisle 15
	for _, l := range constants.SUB_AISLE_15 {
		if l.X == loc.X && l.Y == loc.Y {
			return 15
		}
	}
	return 0
}

func ToCoordinate(x, y int) *models.Coordinate {
	return &models.Coordinate{X: x, Y: y}
}

func PickItem(listPickLoc, listPickableLoc []*models.Coordinate, pickerLoc *models.Coordinate) ([]*models.Coordinate, bool) {
	isPicked := false
	if pickableLoc := utils.Find(listPickableLoc, func(loc *models.Coordinate) bool {
		return loc.X == pickerLoc.X && loc.Y == pickerLoc.Y
	}); pickableLoc != nil {
		if northLoc := utils.IndexOf(listPickLoc, func(loc *models.Coordinate) bool {
			return loc.X == pickerLoc.X && loc.Y == pickerLoc.Y-1
		}); northLoc != -1 {
			isPicked = true
			listPickLoc = append(listPickLoc[:northLoc], listPickLoc[northLoc+1:]...)
		}
		if southLoc := utils.IndexOf(listPickLoc, func(loc *models.Coordinate) bool {
			return loc.X == pickerLoc.X && loc.Y == pickerLoc.Y+1
		}); southLoc != -1 {
			isPicked = true
			listPickLoc = append(listPickLoc[:southLoc], listPickLoc[southLoc+1:]...)
		}
		if westLoc := utils.IndexOf(listPickLoc, func(loc *models.Coordinate) bool {
			return loc.X == pickerLoc.X-1 && loc.Y == pickerLoc.Y
		}); westLoc != -1 {
			isPicked = true
			listPickLoc = append(listPickLoc[:westLoc], listPickLoc[westLoc+1:]...)
		}
		if eastLoc := utils.IndexOf(listPickLoc, func(loc *models.Coordinate) bool {
			return loc.X == pickerLoc.X+1 && loc.Y == pickerLoc.Y
		}); eastLoc != -1 {
			isPicked = true
			listPickLoc = append(listPickLoc[:eastLoc], listPickLoc[eastLoc+1:]...)
		}
	}
	return listPickLoc, isPicked
}
