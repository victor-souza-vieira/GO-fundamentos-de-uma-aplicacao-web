package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func conectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/LojinhaGO")

	if err != nil {
		panic(err.Error())
	}

	return db
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
