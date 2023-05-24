package test

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
)

func GetNextPickLocation(listPickLoc []*models.Coordinate, pickerLoc *models.Coordinate) []*models.Coordinate {

	listCurrentPickLoc := make([]*models.Coordinate, 0)
	if len(listPickLoc) > 0 {
		// Pick Aisle 1 or Asile 2
		for _, loc := range constants.SUB_AISLE_6 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_3 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_5 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_2 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_4 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_1 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		if len(listCurrentPickLoc) > 0 {
			return helper.SortLocationEuclidean(&models.Coordinate{X: 1, Y: 17}, listCurrentPickLoc, false)
		}

		// Pick furthest Aisle 3 or Asile 4 or Asile 5
		for _, loc := range constants.SUB_AISLE_7 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_10 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_13 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		if len(listCurrentPickLoc) > 0 {
			return helper.SortLocationManhattan(&models.Coordinate{X: 4, Y: 0}, listCurrentPickLoc, false)
		}

		// Pick middle Aisle 3 or Asile 4 or Asile 5
		for _, loc := range constants.SUB_AISLE_8 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_11 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_14 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		if pickerLoc.X == 4 || pickerLoc.X == 5 {
			if len(listCurrentPickLoc) > 0 {
				return helper.SortLocationManhattan(&models.Coordinate{X: 4, Y: 6}, listCurrentPickLoc, false)
			}
		}
		if pickerLoc.X == 6 || pickerLoc.X == 7 {
			if len(listCurrentPickLoc) > 0 {
				return helper.SortLocationManhattan(&models.Coordinate{X: 7, Y: 6}, listCurrentPickLoc, false)
			}
		}

		// Pick nearest Aisle 3 or Asile 4 or Asile 5
		for _, loc := range constants.SUB_AISLE_9 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_12 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X - 1, Y: nextPick.Y})
			}
		}
		for _, loc := range constants.SUB_AISLE_15 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listCurrentPickLoc = append(listCurrentPickLoc, &models.Coordinate{X: nextPick.X + 1, Y: nextPick.Y})
			}
		}
		if pickerLoc.X == 4 || pickerLoc.X == 5 {
			if len(listCurrentPickLoc) > 0 {
				return helper.SortLocationManhattan(&models.Coordinate{X: 4, Y: 12}, listCurrentPickLoc, false)
			}
		}
		if pickerLoc.X == 6 || pickerLoc.X == 7 {
			if len(listCurrentPickLoc) > 0 {
				return helper.SortLocationManhattan(&models.Coordinate{X: 7, Y: 12}, listCurrentPickLoc, false)
			}
		}

		// // Pick furthest Aisle 6

		// if len(listCurrentPickLoc) > 0 {
		// 	return helper.SortLocation(&models.Coordinate{X: 7, Y: 6}, listCurrentPickLoc, false)
		// }
	}
	return listCurrentPickLoc
}