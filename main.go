package main

import (
	"net/http"

	"victor.com/module/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
