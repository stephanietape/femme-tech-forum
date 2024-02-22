package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("Go app...")

	home := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" && r.URL.Path == "/" {
			tmpl := template.Must(template.ParseFiles("Templates/index.html"))
			err := tmpl.Execute(w, tmpl)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.NotFound(w, r)

		}
	}

	formulaire := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" && r.URL.Path == "/form" {
			tmpl := template.Must(template.ParseFiles("Templates/formulaire.html"))
			err := tmpl.Execute(w, tmpl)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.NotFound(w, r)

		}
	}

	// define handlers
	http.HandleFunc("/form", formulaire)
	http.HandleFunc("/", home)
	http.Handle("/ressources/", http.StripPrefix("/ressources/", http.FileServer(http.Dir("ressources"))))
	log.Fatal(http.ListenAndServe(":8000", nil))

}
