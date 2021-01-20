package main

import (
	"fmt"
	"log"
)

func main(){
	var f [30]int

	// slice temporary
	log.Print(f[1:9])
	f[3] = 5
	g := f[0:10]
	// slice refers to same array
	log.Print(f, g)

	g[0] = 9

	// cannot append any other data type to sliced literal
	// slice creates a copy of the array
	//g[1] = "cccc"
	log.Print(g)
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)
	// [2 3 5 7 11 13]

	fmt.Println(mySlice[1:4])
	// [3 5 7]

	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// [2 3 5]

	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
	// [11 13]

	//append function
	mystring := []string{}
	mystring = append(mystring, "nagpur")
	mystring = append(mystring, "gondia", "chandrapur")

	log.Print(mystring, len(mystring), cap(mystring))

	make_string := make([]string, 3) // second argument determines the string
	log.Print(make_string, len(make_string), cap(make_string))
	make_string[0] = "asfhaokfv"
	log.Print(make_string)
	log.Print(make_string, len(make_string), cap(make_string))


	// can append a alice to another using ellipses
	cities := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities = append(cities, otherCities...) // same as list.extend in python
	fmt.Printf("%q", cities)


	// nil slices

}


