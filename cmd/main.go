package main

import (
	"fmt"
	"go-algorithms/application/concurrent"
	"go-algorithms/application/constants"
	"go-algorithms/application/location"
	"go-algorithms/application/sort"
	"go-algorithms/application/utils"
	warehouseRouting "go-algorithms/application/warehouse_routing"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// callGenFiles()
	// callGenFilesAndFolders()
	// callReadFileCountWord()
	// callGenLocation()
	callFindPickingRoute()
}

func callGenLocation() {
	location.GenLocation()
}

func callFindPickingRoute() {
	const LOCATION_FILE_PATH = "assets/location.txt"

	file, err := os.ReadFile(LOCATION_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	locations := string(file)
	tempCoordinate := make([]warehouseRouting.Coordinate, 0)
	coordinate := make([]warehouseRouting.Coordinate, 0)
	for _, v := range strings.Split(locations, "\n") {
		axis := strings.Split(v, "-")
		x, y := axis[0], axis[1]
		tempCoordinate = append(tempCoordinate, warehouseRouting.Coordinate{
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
			coordinate = append(coordinate, warehouseRouting.Coordinate{
				X: x,
				Y: y,
			})
		}
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	tempCoordinate = make([]warehouseRouting.Coordinate, 0)
	tempCoordinate = append(tempCoordinate, warehouseRouting.Coordinate{X: 0, Y: 0})
	for i := 0; i < 10; i++ {
		index := r1.Intn(len(coordinate))
		fmt.Printf("Coornation: %v\n", coordinate[index])
		tempCoordinate = append(tempCoordinate, coordinate[index])
	}

	yLow := 0
	yHigh := 16
	waveDistance, routes := warehouseRouting.FindPickingRoute(tempCoordinate[0], tempCoordinate[1:], yLow, yHigh)
	fmt.Printf("Wave distance: %d\n", waveDistance)
	fmt.Printf("Routes: %v\n", routes)
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
