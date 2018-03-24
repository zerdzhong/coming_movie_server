package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func comingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "coming")
}

func filmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("film handler")
}
