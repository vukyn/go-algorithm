package main

import (
	"fmt"
	"go-algorithms/application/concurrent"
	"go-algorithms/application/constants"
	"go-algorithms/application/location"
	"go-algorithms/application/sort"
	"go-algorithms/application/warehouse_routing"
	"io/fs"
	"log"
	"math/rand"
	"os"
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
	waveDistance, routes := warehouseRouting.FindPickingRoute(warehouseRouting.Coordinate{X: 0, Y: 0}, []warehouseRouting.Coordinate{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 3, Y: 3},
		{X: 5, Y: 4},
		{X: 1, Y: 5},
	},0 , 10)
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
