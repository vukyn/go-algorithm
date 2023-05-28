package helper

import (
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/models"
	"math"
	"sort"
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

func CalculateDfsDistance(pickerLoc, nextPickLoc *models.Coordinate, listRemainWalkLoc []*models.Coordinate) []*models.Route {
	i := 0
	distance := 0
	listQueueLoc := make([]*models.Coordinate, 0)
	listVisitedLoc := make([]*models.Coordinate, 0)
	listVisitedLoc = append(listVisitedLoc, pickerLoc)
	currentLoc := pickerLoc
	listQueueLoc = GetPossibleLocation(pickerLoc, listRemainWalkLoc)

	lastLoc := &models.Coordinate{}
	lastLoc = pickerLoc
	listRouteLoc := make([]*models.Route, 0)
	listRouteLoc = append(listRouteLoc, &models.Route{ListVisitedLoc: []*models.Coordinate{lastLoc}})

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
			if !isDeadEnd(currentLoc, nextPickLoc, listRemainWalkLoc) {
				listRouteLoc[i].Distance += 1
				listRouteLoc[i].ListVisitedLoc = append(listRouteLoc[i].ListVisitedLoc, nextPickLoc)
			}
			i++
			distance = 0
			listRouteLoc = append(listRouteLoc, &models.Route{ListVisitedLoc: []*models.Coordinate{lastLoc}})
		} else {
			if !isDeadEnd(currentLoc, nextPickLoc, listRemainWalkLoc) {
				listRouteLoc[i].Distance += 1
				listRouteLoc[i].ListVisitedLoc = append(listRouteLoc[i].ListVisitedLoc, nextPickLoc)
			}
		}
	}
	return listRouteLoc
}

func isDeadEnd(currentLoc, nextPickLoc *models.Coordinate, listWalkLoc []*models.Coordinate) bool {
	possibleLocs := GetPossibleLocation(currentLoc, listWalkLoc)
	if loc := utils.Find(possibleLocs, func(loc *models.Coordinate) bool {
		return loc.X == nextPickLoc.X && loc.Y == nextPickLoc.Y
	}); loc != nil {
		return false
	}
	return true
}

func SortLocationDfs(seedLoc *models.Coordinate, listLoc []*models.Coordinate, listWalkLoc []*models.Coordinate) []*models.Coordinate {
	dfsRoutes := make([]*models.Route, 0)
	sortedLocs := make([]*models.Coordinate, 0)
	length := len(listLoc)
	for i := -1; i < length-1; i++ {
		if i >= 0 {
			seedLoc = sortedLocs[i]
			listLoc = utils.Where(listLoc, func(loc *models.Coordinate) bool {
				return loc.X != seedLoc.X || loc.Y != seedLoc.Y
			})
		}
		if len(listLoc) > 1 {
			for _, loc := range listLoc {
				dfsLocs := CalculateDfsDistance(seedLoc, loc, listWalkLoc)
				// Remove any dead end route
				min := 99999
				resLoc := &models.Coordinate{}
				for _, route := range dfsLocs {
					if route.ListVisitedLoc[len(route.ListVisitedLoc)-1].X == loc.X && route.ListVisitedLoc[len(route.ListVisitedLoc)-1].Y == loc.Y && route.Distance < min {
						min = route.Distance
						resLoc = route.ListVisitedLoc[len(route.ListVisitedLoc)-1]
						resLoc.Id = loc.Id
					}
				}
				if resLoc.X != 0 && resLoc.Y != 0 {
					route := &models.Route{
						Distance:       min,
						ListVisitedLoc: []*models.Coordinate{resLoc},
					}
					dfsRoutes = append(dfsRoutes, route)
				}
			}
			listDistance := make([]int, 0)
			for _, dfsLoc := range dfsRoutes {
				listDistance = append(listDistance, dfsLoc.Distance)
			}
			min := utils.Min(listDistance...)
			idx := utils.IndexOf(dfsRoutes, func(dfsLoc *models.Route) bool {
				return dfsLoc.Distance == min
			})

			// Assign id
			// if dfsRoutes[idx].ListVisitedLoc[0].Id == 0 {

			// }
			sortedLocs = append(sortedLocs, dfsRoutes[idx].ListVisitedLoc[0])
		} else if len(listLoc) == 1 {
			sortedLocs = append(sortedLocs, listLoc[0])
		}
		dfsRoutes = make([]*models.Route, 0)
	}
	return sortedLocs
}

func SortLocationEuclidean(seedLoc *models.Coordinate, listLoc []*models.Coordinate, asc bool) []*models.Coordinate {
	if asc {
		sort.Slice(listLoc, func(i, j int) bool {
			return CalculateEuclideanDistance(seedLoc, listLoc[i]) > CalculateEuclideanDistance(seedLoc, listLoc[j])
		})
	} else {
		sort.Slice(listLoc, func(i, j int) bool {
			return CalculateEuclideanDistance(seedLoc, listLoc[i]) < CalculateEuclideanDistance(seedLoc, listLoc[j])
		})
	}
	return listLoc
}

func SortLocationManhattan(seedLoc *models.Coordinate, listLoc []*models.Coordinate, asc bool) []*models.Coordinate {
	if asc {
		sort.Slice(listLoc, func(i, j int) bool {
			return CalculateManhattanDistance(seedLoc, listLoc[i]) > CalculateManhattanDistance(seedLoc, listLoc[j])
		})
	} else {
		sort.Slice(listLoc, func(i, j int) bool {
			return CalculateManhattanDistance(seedLoc, listLoc[i]) < CalculateManhattanDistance(seedLoc, listLoc[j])
		})
	}
	return listLoc
}
