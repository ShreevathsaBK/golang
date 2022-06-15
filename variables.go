package main

import "fmt"

// A generic function
func printVariable[T any](variable T) {
	fmt.Printf("str : %v \ntype : %T\n\n", variable, variable)
}

func main(){
	var name string = "Shreevathsa"
	var isPositive bool = true
	
	var num uint8 = 255 // max values of uint8 is 255
	var bigNum uint64 = 1234322323123234341
	var val int = 3455


	// implicit type
	str := "string"
	char := 'b'

	printVariable(name)
	printVariable(isPositive)
	printVariable(num)
	printVariable(bigNum)
	printVariable(val)
	printVariable(str)
	printVariable(char)
}