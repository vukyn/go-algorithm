package greedy

import (
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
)

func CallGreedyHeuristic(pickerLoc *models.Coordinate, listPickLoc []*models.Coordinate, listRemainWalkLoc []*models.Coordinate) (minPossibleLoc *models.Coordinate) {
	// Get possible walk location
	listPossibleLoc := helper.GetPossibleLocation(pickerLoc, listRemainWalkLoc)

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
