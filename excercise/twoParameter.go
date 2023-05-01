package twoparameter


import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func TwoParameter(w http.ResponseWriter, r *http.Request ){
	vars := mux.Vars(r)

	title := vars["title"]
	pag := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, pag)
}
