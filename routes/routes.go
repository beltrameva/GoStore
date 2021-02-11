package routes

import (
	"net/http"
	"vanio/controlers"
)

func CarregaRotas() {
	http.HandleFunc("/", controlers.Index)
}
