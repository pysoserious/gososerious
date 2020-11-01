package main


import "fmt"

// Go passes everything my value hence the v here is new variable in new function stack
func modifyVariableValue(v int){
	v = 7
	fmt.Println(v)
}

//func modifyVariableRef(*v){
//	*v = 7
//	fmt.Println(*v)
//}

func main(){
	v := 6
	modifyVariableValue(v)
	fmt.Println(v)
}