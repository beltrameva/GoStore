package main

import (
	"net/http"

	"github.com/beltrameva/Vanio/routes"
)

func main() {
	routes.CarregaRotas
	http.ListenAndServe(":8000", nil)
}
