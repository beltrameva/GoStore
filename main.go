package main

import (
	"html/template"
	"net/http"
	"github.com/Vanio/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) 
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
