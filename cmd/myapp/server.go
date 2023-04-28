package main

import (
	"net/http"
)

func main() {

	
	fs := http.FileServer(http.Dir("../../static/"))
	http.Handle("/static/", http.StripPrefix("/static/",fs))
	


	http.ListenAndServe("127.0.0.1:8080", nil)



}