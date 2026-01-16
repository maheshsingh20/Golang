package main

import (
	"fmt"
	// "sort"
)

func printArray(){
	/*
	arr:=[5]int{56,26,33,44,55}
	for i:=0; i<5;i++{
		fmt.Println(arr[i])
	}
	arr[4]=109
	fmt.Println(len(arr))

	using range printing index and value 
	 for i,v:= range arr{
		fmt.Println(i," ",v)
	 }
	//printing only value in range

	for _,v:=range arr{
		fmt.Println(v);
	}
	var sum int
	for _ , v:=range arr{
		sum+=v
	}
	fmt.Println(sum)\

	

	arr:=[]int{11,22,33,44,55}
	arr = append(arr, 66)
	for _,v:= range arr{
		fmt.Println(v);
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for _,v:= range arr{
		fmt.Println(v);
	}
	// arrCopy := make([]int, len(arr))
	// copy(arrCopy, arr)

	fmt.Printf("arr: %v\n", arr)
	*/

	students:=[]string{"Mahesh", "Kartik","Ankit","Yash","Avinash"}
	fmt.Printf("%v\n",students)

}