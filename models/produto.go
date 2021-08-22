package models

import "victor.com/module/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()
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

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("insert into produtos() values (null, ?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
}

func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletaProduto, err := db.Prepare("delete from produtos where id=?")
	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(idProduto)
}

func EditaProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produtoDoBanco, err := db.Query("select * from produtos where id = ?", idProduto)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	return produtoParaAtualizar
}

func UpdateProduto(id string, nome string, descricao string, preco string, quantidade string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	atualizaDadosProduto, err := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id=?")
	if err != nil {
		panic(err.Error())
	}

	atualizaDadosProduto.Exec(nome, descricao, preco, quantidade, id)
}
