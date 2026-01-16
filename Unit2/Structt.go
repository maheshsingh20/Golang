package main

import "fmt"

type User struct {
	ID     int
	Name   string
	skills []string
}
var Students=[]User{
	{
		ID:	12210582,
		Name:"Mahesh Singh",
		skills:[]string{"ReactJs","CPP","NodeJs","CP"},
	},
	{
		ID:12210583,
		Name:"Chandan Singh",
		skills:[]string{"Software Engineering","Frontend Engineer","Devops Engineer"},
	},
}
func PrintDetail() {
	// fmt.Println(Students)

	for i := 0; i < len(Students); i++ {
		fmt.Printf("Student %d detail ",i+1)
		fmt.Println()
		fmt.Println("Name is:",Students[i].Name)
		fmt.Println("Id is:",Students[i].ID)
		fmt.Println("Skills Are: ")
		for j := 0; j < len(Students[i].skills); j++ {
			fmt.Println(Students[i].skills[j])
		}
	}
}