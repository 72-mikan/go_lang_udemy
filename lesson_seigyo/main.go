package main

import (
	"fmt"
	//"strconv"
	//"os"
	"time"
)

func anything(a interface{}) {
	//fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	/*
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't Know")
	}

	if b := 100; b == 100 {
		fmt.Println("one handred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
	*/

	/*
	var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
	*/

	/*
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}
	*/

	/*
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
	*/

	/*
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	*/

	/*
	arr := [3]int{1, 2, 3}
	for i := 0; i< len(arr); i++ {
		fmt.Println(arr[i])
	}
	*/

	/*
	arr := [3]int{1, 2, 3}

	for i := range arr {
		fmt.Println(i)
	}
	*/

	/*
	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	*/

	/*
	m := map[string]int{"apple": 100, "banana": 200}

	for k := range m {
		fmt.Println(k)
	}
	*/

	/*
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}
	*/

	/*
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}
	*/

	/*
	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}
	*/

	/*
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	default:
		fmt.Println("I don't Know")
	}
	*/

	/*
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i, Int := x.(int)
	fmt.Println(i, Int)
	//fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("I don't Know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't Know")
	}
	*/
	/*
	Loop:
	  for {
			for {
				for {
					fmt.Println("START")
					break Loop
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("END")
		*/

		/*
	Loop:
	  for i := 0; i < 3; i++ {
			for j := 1; j < 3; j++ {
				if j > 1 {
					continue Loop
				}
				fmt.Println(i, j, i*j)
			}
			fmt.Println("処理をしない")
		}
		*/
		
	/*
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
	*/

	/*
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
	*/

	/*
	go sub()
	go sub()
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
	*/

	fmt.Println("main")
}