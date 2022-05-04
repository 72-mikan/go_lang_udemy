package main

import (
	"fmt"
)

/*
type Stringer interface {
	String() string
}
*/

//&{100 ABC} => <<100, ABC>>
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

type Point struct {
	A int
	B string
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}