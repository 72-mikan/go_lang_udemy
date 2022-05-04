package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json

//構造体からJSONテキストへの変換

type A struct {}

type User struct {
	ID      int       `json:"id,omitempty"` //json形式で表示する際の名前を変えることができ、型も変更することができるs
	Name    string    `json:"name"`      //表示したくない場合`json:"-"`で表示されないようにできる
	Email   string    `json:"email"`     //データが初期値の時隠したい場合`json"id,omitempty"`
	Created time.Time `json:"created"`
	A       *A         `json:"A,omitempty"`
}

func main() {
	u := new(User)
	u.ID = 0
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}