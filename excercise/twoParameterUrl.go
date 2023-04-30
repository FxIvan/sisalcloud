package twoparameter

import (
	"net/http"
	"github.com/gorilla/mux"
)

func twoParameter(w http.ResponseWrite, r *http.Request ){
	vars := mux.Vars(r)

	title := vars["title"]
	pag = vars["page"]

	fmt.Println("Title: " + title)
	fmt.Println("Page: " + page)
}
