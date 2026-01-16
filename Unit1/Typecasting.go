package main

import "fmt"

func TypeCasting() {
	
	val1 := 12
	val2 := float64(val1)
	fmt.Println(val2)
	fmt.Printf("%T\n",val2)
	val2=34.4
	fmt.Println(val2)
	val3:=int64(val2)
	fmt.Println(val3)
	fmt.Printf("%T\n",val3)
}