package websockethtml

import (
	"net/http"
)

func RenderizarHTML(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"index.html")
}