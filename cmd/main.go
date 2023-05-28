package main

import (
	"fmt"
	"go-algorithms/application/concurrent"
	"go-algorithms/application/constants"
	"go-algorithms/application/cryption"
	"go-algorithms/application/location"
	"go-algorithms/application/sort"
	"go-algorithms/application/utils"
	warehouseRouting "go-algorithms/application/warehouse_routing"
	warehouseConstants "go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
	nearestNeighbor "go-algorithms/application/warehouse_routing/nearest_neighbor"
	sshape "go-algorithms/application/warehouse_routing/s_shape"
	"go-algorithms/application/warehouse_routing/test"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// callGenFiles()
	// callGenFilesAndFolders()
	// callReadFileCountWord()
	// callGenLocation()
	// callFindPickingRouteNN()
	// callWarehouseRouting()
	// callEncriptMD5()

	pickerLoc := &models.Coordinate{}
	listWallLoc := make([]*models.Coordinate, 0)
	listWalkLoc := make([]*models.Coordinate, 0)
	listRemainWalkLoc := make([]*models.Coordinate, 0)
	listPickLoc := make([]*models.Coordinate, 0)
	listPickableLoc := make([]*models.Coordinate, 0)

	lines := helper.ReadMap()
	for y, line := range lines {
		for x, char := range strings.Split(line, "-") {
			if char == strconv.Itoa(warehouseConstants.WALL) {
				listWallLoc = append(listWallLoc, &models.Coordinate{X: x, Y: y})
			}
			if char == strconv.Itoa(warehouseConstants.WALK) || char == strconv.Itoa(warehouseConstants.PICKABLE) {
				listWalkLoc = append(listWalkLoc, &models.Coordinate{X: x, Y: y})
				listRemainWalkLoc = append(listRemainWalkLoc, &models.Coordinate{X: x, Y: y})
			}
			if char == strconv.Itoa(warehouseConstants.PICKABLE) {
				listPickableLoc = append(listPickableLoc, &models.Coordinate{X: x, Y: y})
			}
			if char == strconv.Itoa(warehouseConstants.DEPOT) {
				pickerLoc = &models.Coordinate{X: x + 1, Y: y}
			}
		}
	}

	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 1, X: 0, Y: 15})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 2, X: 0, Y: 14})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 3, X: 0, Y: 7})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 4, X: 0, Y: 2})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 5, X: 2, Y: 1})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 6, X: 3, Y: 1})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 7, X: 3, Y: 5})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 8, X: 3, Y: 8})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 9, X: 5, Y: 2})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 10, X: 6, Y: 2})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 11, X: 6, Y: 9})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 12, X: 5, Y: 16})
	listPickLoc = append(listPickLoc, &models.Coordinate{Id: 13, X: 6, Y: 16})
	refListPickLoc := listPickLoc
	for {
		nextPickLoc, stage := test.GetNextPickLocation(listPickLoc, listWalkLoc, pickerLoc, 1)
		// pickerLoc = nextPickLoc[len(nextPickLoc)-1]
		fmt.Printf("Stage: %v\n", stage)
		fmt.Print("Locations: ")
		for _, v := range nextPickLoc {
			listPickLoc = utils.Where(listPickLoc, func(c *models.Coordinate) bool {
				return c.Id != v.Id
			})
			if pickLoc := utils.Find(refListPickLoc, func(c *models.Coordinate) bool {
				return c.Id == v.Id
			}); pickLoc != nil {
				fmt.Printf("{X:%v, Y:%v}\t", pickLoc.X, pickLoc.Y)
			}
		}
		fmt.Println()
		if stage == 4 {
			break
		}
	}
	// pickerLoc := &models.Coordinate{X: 4, Y: 16}
	// nextPickLoc := &models.Coordinate{X: 7, Y: 16}
	// visitedLoc, distance := helper.CalculateDfsDistance(pickerLoc, nextPickLoc, listRemainWalkLoc)
}

func callEncriptMD5() {
	cryption.EncriptMD5()
}

func callWarehouseRouting() {
	warehouseRouting.Run(nil)
}

func callGenLocation() {
	location.GenLocation()
}

func callFindPickingRouteNN() {
	const LOCATION_FILE_PATH = "assets/location.txt"

	file, err := os.ReadFile(LOCATION_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	locations := string(file)
	tempCoordinate := make([]models.Coordinate, 0)
	coordinate := make([]models.Coordinate, 0)
	for _, v := range strings.Split(locations, "\n") {
		axis := strings.Split(v, "-")
		x, y := axis[0], axis[1]
		tempCoordinate = append(tempCoordinate, models.Coordinate{
			X: utils.ConvertStringToInt(x),
			Y: utils.ConvertStringToInt(y),
		})
	}
	x := 0
	y := 0
	for _, v := range tempCoordinate {
		if (x != v.X && y != v.Y) || (x == v.X && y != v.Y) {
			x = v.X
			y = v.Y
			coordinate = append(coordinate, models.Coordinate{
				X: x,
				Y: y,
			})
		}
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	tempCoordinate = make([]models.Coordinate, 0)
	tempCoordinate = append(tempCoordinate, models.Coordinate{X: 0, Y: 0})
	for i := 0; i < 10; i++ {
		index := r1.Intn(len(coordinate))
		fmt.Printf("Coornation: %v\n", coordinate[index])
		tempCoordinate = append(tempCoordinate, coordinate[index])
	}

	yLow := 0
	yHigh := 16
	waveDistance, routes := nearestNeighbor.FindPickingRouteNN(tempCoordinate[0], tempCoordinate[1:], yLow, yHigh)
	fmt.Printf("Wave distance: %d\n", waveDistance)
	fmt.Printf("Routes: %v\n", routes)
}

func callFindPickingRouteSShape() {
	sshape.FindPickingRouteSShape()
}

func callReadFileCountWord() {
	const (
		FOLDER_PATH = "assets/read_file_count_word"
		DETECT_WORD = "ut"
	)

	// Scan all files and folders in directory
	dir, err := os.ReadDir(FOLDER_PATH)
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	count := 0
	countCh := make(chan int, len(dir))

	for _, e := range dir {
		wg.Add(1)
		go func(e fs.DirEntry) {
			defer wg.Done()
			countCh <- concurrent.CountWordInFile(e, FOLDER_PATH, DETECT_WORD)
		}(e)
	}
	wg.Wait()
	close(countCh)
	for c := range countCh {
		count += c
	}
	fmt.Printf("Count [%s]: %d", DETECT_WORD, count)
}

func callGenFiles() {
	concurrent.GenFiles(100, constants.FOLDER_PATH)
}

func callGenFilesAndFolders() {
	concurrent.GenFilesAndFolders(100, constants.FOLDER_PATH)
}

func callSelectSort() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 20)
	for i := 0; i < cap(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Printf("nums: %d\n", sort.SelectSort(nums))
}
