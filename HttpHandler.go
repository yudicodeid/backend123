package main

import (
	"net/http"
	"log"
	"html/template"
)

func root(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}

}

func main() {

	http.HandleFunc("/",  root)
	err := http.ListenAndServe(":8080", nil)
	if err!=nil {
		log.Fatal("ListenAndServe", err)
	}

}
