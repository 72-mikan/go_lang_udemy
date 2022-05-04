package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json

//JSONから構造体への変換

type A struct {}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	//----------------------

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	//Unmarshal  JSONをデータに変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}