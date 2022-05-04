package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	/*
		0.813527291469711
		0.5598026045235738
		0.6695717783859498
	*/

}