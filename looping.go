package main

import(
	"fmt"
)

func main(){
	var str string = "abcd"

	for i, j := range str {
		fmt.Printf("%d - %c\n", i, j)
	}

	for k := 0; k<5; k++ {
		fmt.Println(k)
	}

	z := 3
	for z > 0 {
		fmt.Println(z)
		z = z - 1
	}
	
}