package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"vanio/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // Se a requisição for POST
		nome := r.FormValue("nome") //Receber o valor inputado no formulário para o campo NAME
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantiadade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}
