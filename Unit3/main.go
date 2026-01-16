package main

import "fmt"

func main() {

	var varSum = Add(12, 45)
	fmt.Println(varSum)

	var varMinus=Minus(45,23)
	fmt.Println(varMinus)

	fmt.Println(Greeting())
	
	giveDetail(22,"Mahesh Singh")
}