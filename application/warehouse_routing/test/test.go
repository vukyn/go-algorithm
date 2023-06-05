package test

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
)

func GetNextPickLocation(listPickLoc, ListWalkLoc []*models.Coordinate, pickerLoc *models.Coordinate, stage int) ([]*models.Coordinate, int) {

	listNextPickLoc := make([]*models.Coordinate, 0)

	return listNextPickLoc, 0
}

func DAQPickLocation(listPickLoc []*models.Coordinate, lastLoc *models.Coordinate, stage int) (listRemainingPickLoc []*models.Coordinate, listNextPickLoc []*models.Coordinate) {

	// listNextPickLoc := make([]*models.Coordinate, 0)
	switch stage {
	case 1:
		// Pick Aisle 1 or Asile 2
		for _, loc := range constants.SUB_AISLE_6 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		for _, loc := range constants.SUB_AISLE_3 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		for _, loc := range constants.SUB_AISLE_5 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		for _, loc := range constants.SUB_AISLE_2 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		for _, loc := range constants.SUB_AISLE_4 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		for _, loc := range constants.SUB_AISLE_1 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(&models.Coordinate{X: 1, Y: 17}, listNextPickLoc, true)
	case 2:
		// Pick furthest Aisle 3 or Asile 4 or Asile 5 (with half furthest)
		for _, loc := range constants.SUB_AISLE_7 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(&models.Coordinate{X: 4, Y: 0}, listNextPickLoc, false)
	case 3:
		// Pick furthest Aisle 3 or Asile 4 or Asile 5 (with half nearest)
		for _, loc := range constants.SUB_AISLE_8 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(lastLoc, listNextPickLoc, false)
	case 4:
		// Pick middle Aisle 3 or Asile 4 or Asile 5 (with half furthest)
		for _, loc := range constants.SUB_AISLE_9 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(lastLoc, listNextPickLoc, false)
	case 5:
		// Pick middle Aisle 3 or Asile 4 or Asile 5 (with half nearest)
		for _, loc := range constants.SUB_AISLE_10 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(lastLoc, listNextPickLoc, false)
	case 6:
		// Pick nearest Aisle 3 or Asile 4 or Asile 5 (with half furthest)
		for _, loc := range constants.SUB_AISLE_11 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(lastLoc, listNextPickLoc, false)
	case 7:
		// Pick nearest Aisle 3 or Asile 4 or Asile 5 (with half nearest)
		for _, loc := range constants.SUB_AISLE_12 {
			if nextPick := utils.Find(listPickLoc, func(i *models.Coordinate) bool {
				return i.X == loc.X && i.Y == loc.Y
			}); nextPick != nil {
				listPickLoc = utils.Where(listPickLoc, func(i *models.Coordinate) bool {
					return i.Id != nextPick.Id
				})
				listNextPickLoc = append(listNextPickLoc, &models.Coordinate{Id: nextPick.Id, X: nextPick.X, Y: nextPick.Y, PickX: loc.PickX, PickY: loc.PickY})
			}
		}
		listNextPickLoc = helper.SortLocationEuclidean(lastLoc, listNextPickLoc, false)
	}

	listRemainingPickLoc = listPickLoc
	return listRemainingPickLoc, listNextPickLoc
}
