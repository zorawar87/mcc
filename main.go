package main

import (
	"html/template"
	"net/http"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html", "templates/pch.html"))

  http.HandleFunc("/pch", func(w http.ResponseWriter, r *http.Request){
		if err := templates.ExecuteTemplate(w, "pch.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
  })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":80", nil)
}
