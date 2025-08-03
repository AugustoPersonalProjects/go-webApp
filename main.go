package main

import (
	"net/http"

	"github.com/lopeslyra/go-webApp/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}