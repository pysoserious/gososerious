package main

import (
"fmt"
)

type User struct {
	FirstName, LastName string
}

func (refVar User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", refVar.FirstName, refVar.LastName)
}

func main() {
	u := User{"Ajinkya", "Kharatkar"}
	fmt.Println(u.Greeting())
}