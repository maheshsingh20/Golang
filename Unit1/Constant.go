package main
import "fmt"

const( //declaring multiple const at once
	port=2443
	host="localhost"
)

func Constant(){
	fmt.Println(port, host)
}