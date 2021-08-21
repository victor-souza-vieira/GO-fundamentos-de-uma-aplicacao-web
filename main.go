package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Bem maneira fi", Preco: 150, Quantidade: 10},
		{"Notebook", "Brabo", 2500, 2},
		{"Headset", "Redragon", 250, 2},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
