package main

import "fmt"

func update(arg *string) {
	fmt.Println("&arg :", &arg)
	fmt.Println("*arg :", *arg)
	*arg = "shreevathsa"
}

func main(){
	var num int = 10
	var numptr *int = &num

	fmt.Println("num :", num)
	fmt.Println("numptr :", numptr)
	fmt.Println("&num :", &num)
	fmt.Println("*numptr :", *numptr)
	fmt.Println()

	// updating a varible by passing a pointer to function
	name := "shree"
	fmt.Println("name :", name)
	fmt.Println("&name :", &name)
	update(&name)
	fmt.Println("name :", name)

}