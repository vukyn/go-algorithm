package warehouseRouting

type Coordinate struct {
	X int
	Y int
}

type BlockSubAsile struct {
	Name     string
	Distance int
	Loc      []*Coordinate
	// LocS     []string
}

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
	SUB_AISLE_1  = []*Coordinate{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}, {X: 1, Y: 4}, {X: 1, Y: 5}}
	SUB_AISLE_2  = []*Coordinate{{X: 1, Y: 6}, {X: 1, Y: 7}, {X: 1, Y: 8}, {X: 1, Y: 9}, {X: 1, Y: 10}}
	SUB_AISLE_3  = []*Coordinate{{X: 1, Y: 11}, {X: 1, Y: 12}, {X: 1, Y: 13}, {X: 1, Y: 14}, {X: 1, Y: 15}, {X: 1, Y: 16}}
	SUB_AISLE_4  = []*Coordinate{{X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}, {X: 2, Y: 4}}
	SUB_AISLE_5  = []*Coordinate{{X: 2, Y: 5}, {X: 2, Y: 6}, {X: 2, Y: 7}, {X: 2, Y: 8}}
	SUB_AISLE_6  = []*Coordinate{{X: 2, Y: 9}, {X: 2, Y: 10}, {X: 2, Y: 11}, {X: 2, Y: 12}, {X: 2, Y: 13}}
	SUB_AISLE_7  = []*Coordinate{{X: 3, Y: 1}, {X: 3, Y: 2}, {X: 3, Y: 3}, {X: 3, Y: 4}}
	SUB_AISLE_8  = []*Coordinate{{X: 3, Y: 5}, {X: 3, Y: 6}, {X: 3, Y: 7}, {X: 3, Y: 8}}
	SUB_AISLE_9  = []*Coordinate{{X: 3, Y: 9}, {X: 3, Y: 10}, {X: 3, Y: 11}, {X: 3, Y: 12}, {X: 3, Y: 13}}
	SUB_AISLE_10 = []*Coordinate{{X: 4, Y: 1}, {X: 4, Y: 2}, {X: 4, Y: 3}, {X: 4, Y: 4}}
	SUB_AISLE_11 = []*Coordinate{{X: 4, Y: 5}, {X: 4, Y: 6}, {X: 4, Y: 7}, {X: 4, Y: 8}}
	SUB_AISLE_12 = []*Coordinate{{X: 4, Y: 9}, {X: 4, Y: 10}, {X: 4, Y: 11}, {X: 4, Y: 12}}
	SUB_AISLE_13 = []*Coordinate{{X: 5, Y: 1}, {X: 5, Y: 2}, {X: 5, Y: 3}, {X: 5, Y: 4}}
	SUB_AISLE_14 = []*Coordinate{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}, {X: 5, Y: 8}}
	SUB_AISLE_15 = []*Coordinate{{X: 5, Y: 9}, {X: 5, Y: 10}, {X: 5, Y: 11}, {X: 5, Y: 12}}
)

var WAREHOUSE_LAYOUT = []BlockSubAsile{
	{Name: "Block 1", Distance: DISTANCE_DEPOT_TO_BLOCK_1, Loc: SUB_AISLE_1},
	{Name: "Block 2", Distance: DISTANCE_DEPOT_TO_BLOCK_2, Loc: SUB_AISLE_2},
	{Name: "Block 3", Distance: DISTANCE_DEPOT_TO_BLOCK_3, Loc: SUB_AISLE_3},
	{Name: "Block 4", Distance: DISTANCE_DEPOT_TO_BLOCK_4, Loc: append(SUB_AISLE_4, SUB_AISLE_7...)},
	{Name: "Block 5", Distance: DISTANCE_DEPOT_TO_BLOCK_5, Loc: append(SUB_AISLE_5, SUB_AISLE_8...)},
	{Name: "Block 6", Distance: DISTANCE_DEPOT_TO_BLOCK_6, Loc: append(SUB_AISLE_6, SUB_AISLE_9...)},
	{Name: "Block 7", Distance: DISTANCE_DEPOT_TO_BLOCK_7, Loc: append(SUB_AISLE_10, SUB_AISLE_13...)},
	{Name: "Block 8", Distance: DISTANCE_DEPOT_TO_BLOCK_8, Loc: append(SUB_AISLE_11, SUB_AISLE_14...)},
	{Name: "Block 9", Distance: DISTANCE_DEPOT_TO_BLOCK_9, Loc: append(SUB_AISLE_12, SUB_AISLE_15...)},
}
