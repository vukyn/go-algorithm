package sshape

import (
	"fmt"
	"go-algorithms/application/utils"
	"go-algorithms/application/warehouse_routing/constants"
	"go-algorithms/application/warehouse_routing/helper"
	"go-algorithms/application/warehouse_routing/models"
	"log"
	"math"
	"sort"
)

func distributeBlockSubAisle(listLoc []*models.Coordinate) []*models.BlockSubAsile {

	warehouseDefaultLayout := initWarehouseLayout()
	blockSubAisle := make([]*models.BlockSubAsile, 0)
	for _, v := range warehouseDefaultLayout {
		blockSubAisle = append(blockSubAisle, &models.BlockSubAsile{
			Id:         v.Id,
			Name:       v.Name,
			Distance:   v.Distance,
			Coordinate: v.Coordinate,
		})
	}

	// Group all location into block of subaisles
	for _, loc := range listLoc {
		// block sub aisle 1
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 0) {
			blockSubAisle[0].Locs = append(blockSubAisle[0].Locs, loc)
			continue
		}
		// block sub aisle 2
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 1) {
			blockSubAisle[1].Locs = append(blockSubAisle[1].Locs, loc)
			continue
		}
		// block sub aisle 3
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 2) {
			blockSubAisle[2].Locs = append(blockSubAisle[2].Locs, loc)
			continue
		}
		// block sub aisle 4
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 3) || (loc.X == 3 && loc.Y >= 1 && loc.Y <= 4) {
			blockSubAisle[3].Locs = append(blockSubAisle[3].Locs, loc)
			continue
		}
		// block sub aisle 5
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 4) || (loc.X == 3 && loc.Y >= 5 && loc.Y <= 8) {
			blockSubAisle[4].Locs = append(blockSubAisle[4].Locs, loc)
			continue
		}
		// block sub aisle 6
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 5) {
			blockSubAisle[5].Locs = append(blockSubAisle[5].Locs, loc)
			continue
		}
		// block sub aisle 7
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 6) {
			blockSubAisle[6].Locs = append(blockSubAisle[6].Locs, loc)
			continue
		}
		// block sub aisle 8
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 7) {
			blockSubAisle[7].Locs = append(blockSubAisle[7].Locs, loc)
			continue
		}
		// block sub aisle 9
		if helper.IsBelongBlockAisle(warehouseDefaultLayout, loc, 8) {
			blockSubAisle[8].Locs = append(blockSubAisle[8].Locs, loc)
			continue
		}
	}

	// Sort the blocks in decreasing distance from the depot
	blockSubAisle = utils.Where(blockSubAisle, func(i *models.BlockSubAsile) bool {
		return len(i.Locs) > 0
	})
	sort.Slice(blockSubAisle, func(i, j int) bool {
		return blockSubAisle[i].Distance < blockSubAisle[j].Distance
	})

	// Write file location with block aisles
	filePath := "assets/picking_route.json"
	if err := utils.WriteFile(utils.PrettyPrint(blockSubAisle), filePath); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Write file: %s\n", filePath)
	}

	//  Determine the closest (left or rightmost) subaisle with pick in the furthest block from depot
	distanceFurthestBlock := make([]float64, 0)
	for _, loc := range blockSubAisle[0].Locs {
		distanceFurthestBlock = append(distanceFurthestBlock, distanceFromDepot(loc))
	}
	minDistance := utils.MinF(distanceFurthestBlock...)
	indexMinDistance := utils.IndexOf(distanceFurthestBlock, func(d float64) bool { return d == minDistance })

	fmt.Printf("Closest subaisle with pick in the furthest block from depot: %v\n", helper.IsBelongSubAisle(blockSubAisle[0].Locs[indexMinDistance]))

	// listLoc2 := []*Coordinate{
	// 	{X: 4, Y: 9},
	// 	{X: 4, Y: 10},
	// 	{X: 4, Y: 11},
	// 	{X: 4, Y: 12},
	// 	{X: 5, Y: 9},
	// 	{X: 5, Y: 10},
	// 	{X: 5, Y: 11},
	// 	{X: 5, Y: 12},
	// }

	// for _, loc := range listLoc2 {
	// 	fmt.Printf("Distance: %.2f\n", calculateDistanceFromDepot(loc))
	// }
	return blockSubAisle
}

// FindPickingRouteSShape is a function to find route for picking (s-shape heuristic)
func FindPickingRouteSShape() {
	listLoc := []*models.Coordinate{
		{X: 1, Y: 6},
		{X: 1, Y: 11},
		{X: 1, Y: 15},
		{X: 2, Y: 1},
		{X: 2, Y: 6},
		{X: 2, Y: 7},
		{X: 4, Y: 3},
		{X: 4, Y: 11},
		{X: 5, Y: 10},
	}
	listBlockSubAisle := distributeBlockSubAisle(listLoc)

	// routingLoc := make([]*Coordinate, 0)
	nextLoc := &models.BlockSubAsile{}
	for {
		listBlockDistance := make([]int, len(listBlockSubAisle))
		for i, loc := range listBlockSubAisle {
			listBlockDistance[i] = helper.CalculateEuclideanDistance(listBlockSubAisle[0].Coordinate, loc.Coordinate)
		}

		// Minimum Distance
		distanceNext := utils.Min(listBlockDistance...)

		// Location of minimum distance
		index := utils.IndexOf(listBlockDistance, func(dist int) bool {
			return dist == distanceNext
		})

		// Next location is the first location with distance = min
		nextLoc = listBlockSubAisle[index]

		// Next location is removed from the list of candidates
		listBlockSubAisle = append(listBlockSubAisle[:index], listBlockSubAisle[index+1:]...)

		fmt.Printf("Block: %v\n", utils.PrettyPrint(nextLoc))
		if len(listBlockSubAisle) == 0 {
			break
		}
	}
}

func distanceFromDepot(loc *models.Coordinate) float64 {
	// square_root((x2-x1)^2 + (y2-y1)^2)
	return math.Sqrt(math.Pow(float64(loc.X), 2) + math.Pow(float64(loc.Y), 2))
}

func initWarehouseLayout() []*models.BlockSubAsile {
	blockSubAsile := make([]*models.BlockSubAsile, 0)
	blockSubAsile = append(blockSubAsile, constants.WAREHOUSE_LAYOUT...)

	blockSubAsile[0].Locs = append(blockSubAsile[0].Locs, constants.SUB_AISLE_1...)
	blockSubAsile[1].Locs = append(blockSubAsile[1].Locs, constants.SUB_AISLE_2...)
	blockSubAsile[2].Locs = append(blockSubAsile[2].Locs, constants.SUB_AISLE_3...)
	blockSubAsile[3].Locs = append(blockSubAsile[3].Locs, constants.SUB_AISLE_4...)
	blockSubAsile[3].Locs = append(blockSubAsile[3].Locs, constants.SUB_AISLE_7...)
	blockSubAsile[4].Locs = append(blockSubAsile[4].Locs, constants.SUB_AISLE_5...)
	blockSubAsile[4].Locs = append(blockSubAsile[4].Locs, constants.SUB_AISLE_8...)
	blockSubAsile[5].Locs = append(blockSubAsile[5].Locs, constants.SUB_AISLE_6...)
	blockSubAsile[5].Locs = append(blockSubAsile[5].Locs, constants.SUB_AISLE_9...)
	blockSubAsile[6].Locs = append(blockSubAsile[6].Locs, constants.SUB_AISLE_10...)
	blockSubAsile[6].Locs = append(blockSubAsile[6].Locs, constants.SUB_AISLE_13...)
	blockSubAsile[7].Locs = append(blockSubAsile[7].Locs, constants.SUB_AISLE_11...)
	blockSubAsile[7].Locs = append(blockSubAsile[7].Locs, constants.SUB_AISLE_14...)
	blockSubAsile[8].Locs = append(blockSubAsile[8].Locs, constants.SUB_AISLE_12...)
	blockSubAsile[8].Locs = append(blockSubAsile[8].Locs, constants.SUB_AISLE_15...)

	return blockSubAsile
}
