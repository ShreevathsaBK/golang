package main

import(
	"fmt"
)

func main(){
	var str string
	var age int

	fmt.Print("Enter your name : ")
	fmt.Scan(&str)

	fmt.Print("Enter your age : ")
	fmt.Scan(&age)

	// Sprintf can save the formatted string
	result := fmt.Sprintf("Your name is %s and you are %d years old", str, age)
	fmt.Println(result)
	
}