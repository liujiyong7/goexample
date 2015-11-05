package main

import (
	"fmt"
)

func main() {

	///////// base array /////////////////////////////
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println("\n")

	//////// direct create ////////////////////////////

	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1, 2, 3, 4, 5}

	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)
	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] = ", v)
	}

	//////// len cap func /////////////
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}
