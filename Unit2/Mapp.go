package main
import "fmt"

func Mapp(){
	mpp:=make(map[string]int)
	mpp["Mahesh"]=22
	mpp["Anshika"]=18
	for u,v  := range mpp {
		if(u!="Anshika"){
			fmt.Println(v)
		}else{
			fmt.Println("She is Anshika")
		}
	}
	
}