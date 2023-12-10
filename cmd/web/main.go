package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}