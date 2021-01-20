package main

import "log"

func main(){

	// 
	var arr [9]int
	arr[0] = 1
	log.Print("%q", arr)

	// you can also use to set names as the entries.
	a := [...]string{"ajinkuya", "kharatkar"}
	log.Printf("%q", a)

	log.Printf("len=%d | capacity=%d ", len(a), cap(a))

}