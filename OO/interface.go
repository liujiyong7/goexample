package main

import (
	"fmt"
)

type A struct {
	a int
}

type B struct {
	A
}

func Egg(b B) {

}

func IsZero(a interface{})b {
	return nil
}
func main() {
	d := make(B)
	Egg(d)
	fmt.Println("hello")
}
