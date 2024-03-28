package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MalcolmFuchs/Credit-Card-Validator/api/component"
)

func main() {
	const portNum string = ":8080"

	log.Println("Starting our simple http server.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmplIndex, err := template.ParseFiles(filepath.Join("..", "templates", "index.templ"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ccNumber := r.FormValue("cc_number")
		isValid := component.IsValidLuhn(ccNumber)

		err = tmplIndex.Execute(w, isValid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})

	log.Println("Started on port", portNum)

	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
