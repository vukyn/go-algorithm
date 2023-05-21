package warehouseRouting

type Coordinate struct {
	X    int
	Y    int
}

type BlockSubAsile struct {
	Id         int
	Name       string
	Distance   int
	Coordinate *Coordinate
	Loc        []*Coordinate
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

var WAREHOUSE_LAYOUT = []*BlockSubAsile{
	{Id: 1, Coordinate: &Coordinate{X: 1, Y: 1}, Name: "Block 1", Distance: DISTANCE_DEPOT_TO_BLOCK_1},
	{Id: 2, Coordinate: &Coordinate{X: 1, Y: 2}, Name: "Block 2", Distance: DISTANCE_DEPOT_TO_BLOCK_2},
	{Id: 3, Coordinate: &Coordinate{X: 1, Y: 3}, Name: "Block 3", Distance: DISTANCE_DEPOT_TO_BLOCK_3},
	{Id: 4, Coordinate: &Coordinate{X: 2, Y: 1}, Name: "Block 4", Distance: DISTANCE_DEPOT_TO_BLOCK_4},
	{Id: 5, Coordinate: &Coordinate{X: 2, Y: 2}, Name: "Block 5", Distance: DISTANCE_DEPOT_TO_BLOCK_5},
	{Id: 6, Coordinate: &Coordinate{X: 2, Y: 3}, Name: "Block 6", Distance: DISTANCE_DEPOT_TO_BLOCK_6},
	{Id: 7, Coordinate: &Coordinate{X: 3, Y: 1}, Name: "Block 7", Distance: DISTANCE_DEPOT_TO_BLOCK_7},
	{Id: 8, Coordinate: &Coordinate{X: 3, Y: 2}, Name: "Block 8", Distance: DISTANCE_DEPOT_TO_BLOCK_8},
	{Id: 9, Coordinate: &Coordinate{X: 3, Y: 3}, Name: "Block 9", Distance: DISTANCE_DEPOT_TO_BLOCK_9},
}

func isBelongBlockAisle(layout []*BlockSubAsile, loc *Coordinate, blockIndex int) bool {
	listLoc := layout[blockIndex]
	for _, l := range listLoc.Loc {
		if l.X == loc.X && l.Y == loc.Y {
			return true
		}
	}
	return false
}

func isBelongSubAisle(loc *Coordinate) int {
	// Sub aisle 1
	for _, l := range SUB_AISLE_1 {
		if l.X == loc.X && l.Y == loc.Y {
			return 1
		}
	}
	// Sub aisle 2
	for _, l := range SUB_AISLE_2 {
		if l.X == loc.X && l.Y == loc.Y {
			return 2
		}
	}
	// Sub aisle 3
	for _, l := range SUB_AISLE_3 {
		if l.X == loc.X && l.Y == loc.Y {
			return 3
		}
	}
	// Sub aisle 4
	for _, l := range SUB_AISLE_4 {
		if l.X == loc.X && l.Y == loc.Y {
			return 4
		}
	}
	// Sub aisle 5
	for _, l := range SUB_AISLE_5 {
		if l.X == loc.X && l.Y == loc.Y {
			return 5
		}
	}
	// Sub aisle 6
	for _, l := range SUB_AISLE_6 {
		if l.X == loc.X && l.Y == loc.Y {
			return 6
		}
	}
	// Sub aisle 7
	for _, l := range SUB_AISLE_8 {
		if l.X == loc.X && l.Y == loc.Y {
			return 7
		}
	}
	// Sub aisle 8
	for _, l := range SUB_AISLE_8 {
		if l.X == loc.X && l.Y == loc.Y {
			return 8
		}
	}
	// Sub aisle 9
	for _, l := range SUB_AISLE_9 {
		if l.X == loc.X && l.Y == loc.Y {
			return 9
		}
	}
	// Sub aisle 10
	for _, l := range SUB_AISLE_10 {
		if l.X == loc.X && l.Y == loc.Y {
			return 10
		}
	}
	// Sub aisle 11
	for _, l := range SUB_AISLE_11 {
		if l.X == loc.X && l.Y == loc.Y {
			return 11
		}
	}
	// Sub aisle 12
	for _, l := range SUB_AISLE_12 {
		if l.X == loc.X && l.Y == loc.Y {
			return 12
		}
	}
	// Sub aisle 13
	for _, l := range SUB_AISLE_13 {
		if l.X == loc.X && l.Y == loc.Y {
			return 13
		}
	}
	// Sub aisle 14
	for _, l := range SUB_AISLE_14 {
		if l.X == loc.X && l.Y == loc.Y {
			return 14
		}
	}
	// Sub aisle 15
	for _, l := range SUB_AISLE_15 {
		if l.X == loc.X && l.Y == loc.Y {
			return 15
		}
	}
	return 0
}
