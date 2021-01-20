package main

// list of constants
const (
	ConstExample = "const value"
)

// list of variable
var (
	varGo = "variable"
)

type City struct {
	name, list string
	Location *Coordinate
}

type Coordinate struct {
	lat, long float64
}

func NewUser(firstName string) *City {

	return &City{"Gurgaon", &Coordinate{45.644, 5656.99}}

}


