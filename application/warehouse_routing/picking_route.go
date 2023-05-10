package warehouseRouting

import "go-algorithms/application/utils"

// Reference: https://www.samirsaci.com/improve-warehouse-productivity-using-order-batching-with-python

type Coordinate struct {
	X int
	Y int
}

// FindPickingRoute is a function to find route for picking
func FindPickingRoute(originLoc Coordinate, listLoc []Coordinate, yLow, yHigh int) (int, []Coordinate) {

	// Total distance
	waveDistance := 0

	// Current location
	startLoc := originLoc

	// List of locations to visit
	listChemin := make([]Coordinate, 0)
	listChemin = append(listChemin, startLoc)

	for {
		// Going to next location
		var nextLoc Coordinate
		var distanceNext int
		listLoc, _, nextLoc, distanceNext = nextLocation(startLoc, listLoc, yLow, yHigh)
		// If there is no more location to visit
		if len(listLoc) == 0 {
			break
		}
		// Update start_loc
		startLoc = nextLoc
		listChemin = append(listChemin, startLoc)
		// Update distance
		waveDistance = waveDistance + distanceNext
	}

	// Final distance from last storage location to origin
	waveDistance = waveDistance + distanceBetweenTwoLocations(startLoc, originLoc, yLow, yHigh)
	listChemin = append(listChemin, originLoc)

	return waveDistance, listChemin
}

// DistanceBetweenTwoLocations is a function to calculate distance between two locations
// loc1(Start Point) and loc2(End Point) are coordinate of two locations
func distanceBetweenTwoLocations(loc1, loc2 Coordinate, yLow, yHigh int) int {
	var (
		distanceX  int
		distanceY  int
		distanceY1 int
		distanceY2 int
	)

	// Distance x-axis
	distanceX = utils.Abs(loc2.X - loc1.X)

	// Distance y-axis
	if loc1.X == loc2.X {
		distanceY1 = utils.Abs(loc2.Y - loc1.Y)
		distanceY2 = distanceY1
	} else {
		distanceY1 = (yHigh - loc1.Y) + (yHigh - loc2.Y)
		distanceY2 = (loc1.Y - yLow) + (loc2.Y - yLow)
	}

	// Minimum distance on y-axis
	distanceY = utils.Min(distanceY1, distanceY2)

	// Total distance
	return distanceX + distanceY
}

// NextLocation is a function to find closest next location
func nextLocation(startLoc Coordinate, listLoc []Coordinate, yLow, yHigh int) ([]Coordinate, Coordinate, Coordinate, int) {

	// Distance to every next points candidate
	listDistance := make([]int, len(listLoc))
	for i, loc := range listLoc {
		listDistance[i] = distanceBetweenTwoLocations(startLoc, loc, yLow, yHigh)
	}

	// Minimum Distance
	distanceNext := utils.Min(listDistance...)

	// Location of minimum distance
	index := utils.IndexOf(listDistance, func(dist int) bool {
		return dist == distanceNext
	})

	// Next location is the first location with distance = min
	nextLoc := listLoc[index]

	// Next location is removed from the list of candidates
	listLoc = append(listLoc[:index], listLoc[index+1:]...)

	return listLoc, startLoc, nextLoc, distanceNext
}
