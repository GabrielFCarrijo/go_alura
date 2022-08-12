package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade float64
}

var html = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produto := []Produto{
		{Nome:"Camiseta", Descricao: "Azul, especial", Preco: 60, Quantidade: 10},
		{"Tenis","Confortavel", 89,3},
		{"Fone","Muito Bom", 59,2},
		{"Blusa","Muito Bom", 59,2},
		
	}
	html.ExecuteTemplate(w, "Index", produto)
}