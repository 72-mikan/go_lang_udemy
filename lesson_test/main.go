package main

import (
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

//テスト

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
}