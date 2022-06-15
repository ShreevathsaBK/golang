package main

import "fmt"

func main(){
	// var arr [3]int = [3]int{1,2,3}
	// var arr = [3]int{1,2,3}
	arr := [3]int{1,2,3}
	names := [2]string{"steven", "marcus"}

	fmt.Println("arr :",arr, ", len(arr) :", len(arr))
	fmt.Println("names :",names, ", len(names) :", len(names))

	// slices
	slice := []int{1,2,3,4,3}
	slice = append(slice, 5)
	fmt.Println("slice :", slice)

	rangeSlice := slice[2:4]
	fmt.Println("slice[2:4] :", rangeSlice)
	fmt.Println("slice[2:] :", slice[2:])
}