package main

import (
	"net/http"
	//"html/template"
	//"log"
	"fmt"
)

//net/http
//サーバー

/*
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
*/

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111!")
}

func main() {
	http.HandleFunc("/top", top)
	//サーバーの立ち上げ
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", &MyHandler{})
}