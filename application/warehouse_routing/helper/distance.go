package helper

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/models"
	"math"
)

func CalculateEuclideanDistance(loc1, loc2 *models.Coordinate) int {
	return int(math.Sqrt(math.Pow(float64(loc2.X-loc1.X), 2) + math.Pow(float64(loc2.Y-loc1.Y), 2)))
}

func CalculateManhattanDistance(loc1, loc2 *models.Coordinate) int {
	return utils.Abs(loc2.X-loc1.X) + utils.Abs(loc2.Y-loc1.Y)
}

func CalculateBfsDistance(pickerLoc *models.Coordinate, listRemainWalkLoc []*models.Coordinate) ([]*models.Coordinate, int) {
	distance := 0
	listQueueLoc := make([]*models.Coordinate, 0)
	listVisitedLoc := make([]*models.Coordinate, 0)
	listVisitedLoc = append(listVisitedLoc, pickerLoc)

	currentLoc := pickerLoc
	possibleLoc := GetPossibleLocation(pickerLoc, listRemainWalkLoc)
	for _, loc := range possibleLoc {
		distance += 1
		listQueueLoc = append(listQueueLoc, loc)
		for len(listQueueLoc) > 0 {
			currentLoc = listQueueLoc[0]
			listVisitedLoc = append(listVisitedLoc, currentLoc)
			listQueueLoc = listQueueLoc[1:] // popleft
			childPossibleLoc := GetPossibleLocation(currentLoc, listRemainWalkLoc)

			// Exclude visited location
			for _, loc := range listVisitedLoc {
				childPossibleLoc = utils.Where(childPossibleLoc, func(childLoc *models.Coordinate) bool {
					return childLoc.X != loc.X || childLoc.Y != loc.Y
				})
			}

			if len(childPossibleLoc) > 0 {
				distance += len(childPossibleLoc)
				listQueueLoc = append(listQueueLoc, childPossibleLoc...)
			}
		}
	}

	return listVisitedLoc, distance
}

type Route struct {
	Distance       int
	ListVisitedLoc []*models.Coordinate
}

func CalculateDfsDistance(pickerLoc, nextPickLoc *models.Coordinate, listRemainWalkLoc []*models.Coordinate) []*Route {
	i := 0
	distance := 0
	listQueueLoc := make([]*models.Coordinate, 0)
	listVisitedLoc := make([]*models.Coordinate, 0)
	listVisitedLoc = append(listVisitedLoc, pickerLoc)
	currentLoc := pickerLoc
	listQueueLoc = GetPossibleLocation(pickerLoc, listRemainWalkLoc)

	lastLoc := &models.Coordinate{}
	lastLoc = pickerLoc
	listRouteLoc := make([]*Route, 0)
	listRouteLoc = append(listRouteLoc, &Route{ListVisitedLoc: []*models.Coordinate{lastLoc}})

	for len(listQueueLoc) > 0 {
		distance += 1
		currentLoc = listQueueLoc[0]
		listQueueLoc = listQueueLoc[1:] // popleft
		listRouteLoc[i].Distance = distance
		listRouteLoc[i].ListVisitedLoc = append(listRouteLoc[i].ListVisitedLoc, currentLoc)
		listVisitedLoc = append(listVisitedLoc, currentLoc)
		childPossibleLoc := GetPossibleLocation(currentLoc, listRemainWalkLoc)

		// Exclude visited location
		for _, loc := range listVisitedLoc {
			childPossibleLoc = utils.Where(childPossibleLoc, func(childLoc *models.Coordinate) bool {
				return childLoc.X != loc.X || childLoc.Y != loc.Y
			})
		}
		if len(childPossibleLoc) > 0 && (childPossibleLoc[0].X != nextPickLoc.X || childPossibleLoc[0].Y != nextPickLoc.Y) {
			childPossibleLoc = SortLocationEuclidean(nextPickLoc, childPossibleLoc, false)
			listQueueLoc = append(childPossibleLoc[:1], listQueueLoc...)
		} else if len(listQueueLoc) > 0 {
			listRouteLoc[i].Distance += 1
			listRouteLoc[i].ListVisitedLoc = append(listRouteLoc[i].ListVisitedLoc, nextPickLoc)
			i++
			distance = 0
			listRouteLoc = append(listRouteLoc, &Route{ListVisitedLoc: []*models.Coordinate{lastLoc}})
		} else {
			listRouteLoc[i].Distance += 1
			listRouteLoc[i].ListVisitedLoc = append(listRouteLoc[i].ListVisitedLoc, nextPickLoc)
		}
	}
	return listRouteLoc
}
