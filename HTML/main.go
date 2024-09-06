package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pessoa struct {
	Nome      string
	Sobrenome string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := pessoa{"Guilherme", "Kunsch"}
		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Servidor escutando")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
