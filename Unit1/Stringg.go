package main

import "fmt"

func Stringg(){
	val:="Mahesh Singh" //short hand syntax 
	var email string="maheshsingh0905a@gmail.com"
	fmt.Printf("%T\n",val)
	/*If we define variable but not assigning then we need to use belo mehtod onlt */
	var course string
	course="Engineering"
	fmt.Println("Course is",course)
	fmt.Println(email)
}