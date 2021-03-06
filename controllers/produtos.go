package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"victor.com/module/models"
)

const returnStatus = 301

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)

	}
	http.Redirect(w, r, "/", returnStatus)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)

	http.Redirect(w, r, "/", returnStatus)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro ao converter o ID:", err)
		}
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro ao converter o preço:", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro ao converter a quantidade:", err)
		}

		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", returnStatus)
}
