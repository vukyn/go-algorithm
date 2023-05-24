package constants

import "go-algorithms/application/warehouse_routing/models"

const (
	MAX_RACK_AISLE_1 = 16
	MAX_RACK_AISLE_2 = 13
	MAX_RACK_AISLE_3 = 13
	MAX_RACK_AISLE_4 = 12
	MAX_RACK_AISLE_5 = 12

	DISTANCE_DEPOT_TO_BLOCK_1 = 1
	DISTANCE_DEPOT_TO_BLOCK_2 = 2
	DISTANCE_DEPOT_TO_BLOCK_3 = 3
	DISTANCE_DEPOT_TO_BLOCK_4 = 1
	DISTANCE_DEPOT_TO_BLOCK_5 = 2
	DISTANCE_DEPOT_TO_BLOCK_6 = 3
	DISTANCE_DEPOT_TO_BLOCK_7 = 2
	DISTANCE_DEPOT_TO_BLOCK_8 = 3
	DISTANCE_DEPOT_TO_BLOCK_9 = 4
)

var (
	SUB_AISLE_1  = []*models.Coordinate{{X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}, {X: 0, Y: 4}, {X: 0, Y: 5}}
	SUB_AISLE_2  = []*models.Coordinate{{X: 0, Y: 6}, {X: 0, Y: 7}, {X: 0, Y: 8}, {X: 0, Y: 9}, {X: 0, Y: 10}}
	SUB_AISLE_3  = []*models.Coordinate{{X: 0, Y: 11}, {X: 0, Y: 12}, {X: 0, Y: 13}, {X: 0, Y: 14}, {X: 0, Y: 15}, {X: 0, Y: 16}}
	SUB_AISLE_4  = []*models.Coordinate{{X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}, {X: 2, Y: 4}, {X: 2, Y: 5}}
	SUB_AISLE_5  = []*models.Coordinate{{X: 2, Y: 7}, {X: 2, Y: 8}, {X: 2, Y: 9}, {X: 2, Y: 10}}
	SUB_AISLE_6  = []*models.Coordinate{{X: 2, Y: 13}, {X: 2, Y: 14}, {X: 2, Y: 15}, {X: 2, Y: 16}}
	SUB_AISLE_7  = []*models.Coordinate{{X: 3, Y: 1}, {X: 3, Y: 2}, {X: 3, Y: 3}, {X: 3, Y: 4}, {X: 3, Y: 5}}
	SUB_AISLE_8  = []*models.Coordinate{{X: 3, Y: 7}, {X: 3, Y: 8}, {X: 3, Y: 9}, {X: 3, Y: 10}}
	SUB_AISLE_9  = []*models.Coordinate{{X: 3, Y: 13}, {X: 3, Y: 14}, {X: 3, Y: 15}, {X: 3, Y: 16}}
	SUB_AISLE_10 = []*models.Coordinate{{X: 5, Y: 2}, {X: 5, Y: 3}, {X: 5, Y: 4}, {X: 5, Y: 5}}
	SUB_AISLE_11 = []*models.Coordinate{{X: 5, Y: 7}, {X: 5, Y: 8}, {X: 5, Y: 9}, {X: 5, Y: 10}}
	SUB_AISLE_12 = []*models.Coordinate{{X: 5, Y: 13}, {X: 5, Y: 14}, {X: 5, Y: 15}, {X: 5, Y: 16}}
	SUB_AISLE_13 = []*models.Coordinate{{X: 6, Y: 2}, {X: 6, Y: 3}, {X: 6, Y: 4}, {X: 6, Y: 5}}
	SUB_AISLE_14 = []*models.Coordinate{{X: 6, Y: 7}, {X: 6, Y: 8}, {X: 6, Y: 9}, {X: 6, Y: 10}}
	SUB_AISLE_15 = []*models.Coordinate{{X: 6, Y: 13}, {X: 6, Y: 14}, {X: 6, Y: 15}, {X: 6, Y: 16}}
)

var WAREHOUSE_LAYOUT = []*models.BlockSubAsile{
	{Id: 1, Coordinate: &models.Coordinate{X: 1, Y: 1}, Name: "Block 1", Distance: DISTANCE_DEPOT_TO_BLOCK_1},
	{Id: 2, Coordinate: &models.Coordinate{X: 1, Y: 2}, Name: "Block 2", Distance: DISTANCE_DEPOT_TO_BLOCK_2},
	{Id: 3, Coordinate: &models.Coordinate{X: 1, Y: 3}, Name: "Block 3", Distance: DISTANCE_DEPOT_TO_BLOCK_3},
	{Id: 4, Coordinate: &models.Coordinate{X: 2, Y: 1}, Name: "Block 4", Distance: DISTANCE_DEPOT_TO_BLOCK_4},
	{Id: 5, Coordinate: &models.Coordinate{X: 2, Y: 2}, Name: "Block 5", Distance: DISTANCE_DEPOT_TO_BLOCK_5},
	{Id: 6, Coordinate: &models.Coordinate{X: 2, Y: 3}, Name: "Block 6", Distance: DISTANCE_DEPOT_TO_BLOCK_6},
	{Id: 7, Coordinate: &models.Coordinate{X: 3, Y: 1}, Name: "Block 7", Distance: DISTANCE_DEPOT_TO_BLOCK_7},
	{Id: 8, Coordinate: &models.Coordinate{X: 3, Y: 2}, Name: "Block 8", Distance: DISTANCE_DEPOT_TO_BLOCK_8},
	{Id: 9, Coordinate: &models.Coordinate{X: 3, Y: 3}, Name: "Block 9", Distance: DISTANCE_DEPOT_TO_BLOCK_9},
}
