package main

import "log"

type coordinate struct {
	lat, lng float64
}

// Maps
func main(){
	log.Printf("start of program")

	my_map := map[string]int64{
		"ajinkya": 400000,
		"Kharatkar": 50000,
	}

	for k, v := range my_map{
		log.Print(k, v)
	}

	var m = make(map[string]coordinate)
	m["Nagpur"] = coordinate{4.54,234.23}

	log.Printf("%v", m)


	// while declaring map if we assign the values we dont need to mention the type
	var coord_map = map[string]coordinate{
		"Nagpur": {34,45},
	}

	log.Printf("%v", coord_map)



	// Mutating maps
	// delete an element in map
	delete(my_map, "ajinkya")
	log.Printf("%v", my_map)

	elem, ok := my_map["Kharatkar"]
	log.Print(elem, ok)

	elem, ok = my_map["ajinkya"]
	log.Print(elem, ok)
}