package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

func Mod(a int, b int) int {
	return a % b
}

//function with string return type

func Greeting() string {
	return "Hello Buddy"
}

func giveDetail(a int, b string) {
	fmt.Printf("My name is %s and my age is %d\n", b, a)
}
