package main

import (
	"fmt"
)

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	/*
	独自型と一般型では計算することができない
	i := 100
	fmt.Println(mi + i)
	*/

	mi.Print()

}