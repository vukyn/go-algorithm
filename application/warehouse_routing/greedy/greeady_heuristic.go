package greedy

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
)

func getPossibleLocation(currentLoc *models.Coordinate, locations []*models.Coordinate) []*models.Coordinate {
	listPossibleLoc := make([]*models.Coordinate, 0)
	if northLoc := utils.Find(locations, func(loc *models.Coordinate) bool {
		return loc.X == currentLoc.X && loc.Y == currentLoc.Y-1
	}); northLoc != nil {
		listPossibleLoc = append(listPossibleLoc, northLoc)
	}
	if southLoc := utils.Find(locations, func(loc *models.Coordinate) bool {
		return loc.X == currentLoc.X && loc.Y == currentLoc.Y+1
	}); southLoc != nil {
		listPossibleLoc = append(listPossibleLoc, southLoc)
	}
	if westLoc := utils.Find(locations, func(loc *models.Coordinate) bool {
		return loc.X == currentLoc.X-1 && loc.Y == currentLoc.Y
	}); westLoc != nil {
		listPossibleLoc = append(listPossibleLoc, westLoc)
	}
	if eastLoc := utils.Find(locations, func(loc *models.Coordinate) bool {
		return loc.X == currentLoc.X+1 && loc.Y == currentLoc.Y
	}); eastLoc != nil {
		listPossibleLoc = append(listPossibleLoc, eastLoc)
	}
	return listPossibleLoc
}

func CallGreedyHeuristic(pickerLoc *models.Coordinate, listPickLoc []*models.Coordinate, listRemainWalkLoc []*models.Coordinate) (minPossibleLoc *models.Coordinate) {
	// Get possible walk location
	listPossibleLoc := getPossibleLocation(pickerLoc, listRemainWalkLoc)

	// Calulate distance from possible walk location to nearest pick location (euclidean distance)
	minDistance := 99999
	for _, possibleLoc := range listPossibleLoc {
		for _, pickLoc := range listPickLoc {
			distance := helper.CalculateEuclideanDistance(possibleLoc, pickLoc)
			if distance < minDistance {
				minDistance = distance
				minPossibleLoc = possibleLoc
			}
		}
	}
	return minPossibleLoc
}
