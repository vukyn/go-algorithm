package models

type Coordinate struct {
	X int
	Y int
}

type Route struct {
	Distance       int
	ListVisitedLoc []*Coordinate
}

type BlockSubAsile struct {
	Id         int
	Name       string
	Distance   int
	Coordinate *Coordinate
	Locs       []*Coordinate
}
