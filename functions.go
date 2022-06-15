package main

import(
	"fmt"
)

func sum(a int, b int) int {
	return a + b;
}

func factorial(n int) int {
	if(n == 0){
		return 1
	}
	return n * factorial(n-1)
}

func f(n int) int {
	return n * 2;
}

func myMap(arr []int, predicate func(int) int) {
	for i:=0; i<len(arr); i++ {
		arr[i] = predicate(arr[i])
	}
}

func getCubeAreaVolume(side int) (int, int) {
	return 6 * side * side, side * side * side
}


func main(){
	fmt.Println("sum(4,5) :", sum(4, 5))
	fmt.Println("factorial :", factorial(5))
	
	// passing function as an argument
	arr := []int{1,2,3,4}
	myMap(arr, f)

	fmt.Println("arr :", arr)

	// multiple return value function
	area, volume := getCubeAreaVolume(25)
	fmt.Println(area, volume)
}