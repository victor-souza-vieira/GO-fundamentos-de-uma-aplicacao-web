package routes

import (
	"net/http"

	"victor.com/module/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
