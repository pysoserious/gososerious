package exports

import "log"

func Add(a, b int) int{
	log.Println("Add function")
	return a+b
}