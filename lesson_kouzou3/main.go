package main

import (
	"fmt"
)

type T struct {
	User User
	//User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	//型の指定をUserのみにしたときに有効
	//fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
	
}