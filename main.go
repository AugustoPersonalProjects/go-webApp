package main

import (
    "html/template"
    "net/http"
)

type Produto struct {
	Nome      string
	Descricao string
	Preco     float64
	Quantidade int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", Index)
    http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul Bem bonita", Preco: 39.99, Quantidade: 10},
		{"Tenis", "Confortável", 299.99, 3},
		{"Fone", "Muito Bom", 59.99, 20},
	}

    tmpl.ExecuteTemplate(w, "Index", produtos)
}	