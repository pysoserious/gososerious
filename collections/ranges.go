package main

import "log"

func main(){

	my_arr := []int{11,21,13,41,15,16,17}

	for idx, v := range my_arr {

		log.Printf("idx=%d and val=%d", idx, v)
		if idx % 2 == 0 {
			// skipping even index
			continue
		}

	}

	// skip the idx
	for _, v := range my_arr {
		log.Print(v)
	}

}