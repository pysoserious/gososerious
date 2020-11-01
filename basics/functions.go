package main

import "fmt"



func Add1(x int, y int) int {
	return x + y
}

// function with explicit return type
// if return references are assigned a value. same is returned.
func addSub(x, y int) (sum, sub int){
	sum = x+y
	sub = x-y // if not initialized, int is assigned as 0 and 0 is returned
	return // return all references
}

func main(){
	//var sum int
	//var sub int
	sum, sub := addSub(4,3)
	fmt.Printf("sum %d and sub %d", sum, sub)
}