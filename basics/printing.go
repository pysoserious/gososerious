package main

import (

	"gososerious/exports"
	"log"
)

// init function is called before main and when package is initialized
// in lexicographical order( alphabetically of file names in a package)
func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate| log.Lmicroseconds | log.Llongfile)
	log.Println("Log initialized")
}

func main(){

	sum := exports.Add(3 ,4)
	log.Println(sum)
}