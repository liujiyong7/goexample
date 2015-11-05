package main

import (
	"fmt"
)

func main() {
	var c byte
	var a1 [32]byte
	var a12 [32][3]byte
	var b1 [32]*byte

	a1[0] = 'c'
	a12[0][2] = 'c'
	b1[0] = &c
	c = 'c'
	fmt.Println(a1)
	fmt.Println(a12)
	fmt.Println(*b1[0])
}
