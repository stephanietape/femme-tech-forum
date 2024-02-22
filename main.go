package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Go app...")

	type User struct {
		Id         string `json:"id"`
		Nom        string `json:"nom"`
		Prenom     string `json:"prenom"`
		Age        string `json:"age"`
		Contact    string `json:"conatct"`
		Email      string `json:"email"`
		Competence string `json:"competence"`
		Profession string `json:"prefession"`
		Fichier    string `json:"fichier"`
		Video      string `json:"video"`
	}

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

	inscription := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" && r.URL.Path == "/enregistrer" {
			id := uuid.New()
			fmt.Println(id.String())

			var mystruct []User
			var user_one User
			user_one.Id = id.String()
			user_one.Nom = r.PostFormValue("nom")
			user_one.Prenom = r.PostFormValue("prenom")
			user_one.Age = r.PostFormValue("age")
			user_one.Email = r.PostFormValue("email")
			user_one.Contact = r.PostFormValue("contact")
			user_one.Competence = r.PostFormValue("competence")
			user_one.Profession = r.PostFormValue("profession")

			mystruct = append(mystruct, user_one)

			fmt.Println(mystruct)
			jsonfromslice, err := json.MarshalIndent(mystruct, "", " ")
			if err != nil {
				fmt.Println("Error unmarshalling json", err)
			}
			fmt.Println(string(jsonfromslice))

			jsonData, err := json.MarshalIndent(jsonfromslice, "", "  ")
			if err != nil {
				fmt.Println("Erreur lors de la sérialisation JSON:", err)
				return
			}

			err = ioutil.WriteFile("users.json", jsonData, 0644)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier JSON:", err)
				return
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	// define handlers
	http.HandleFunc("/form", formulaire)
	http.HandleFunc("/", home)
	http.HandleFunc("/enregistrer", inscription)
	http.Handle("/ressources/", http.StripPrefix("/ressources/", http.FileServer(http.Dir("ressources"))))
	log.Fatal(http.ListenAndServe(":8000", nil))

}
