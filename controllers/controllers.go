package controllers

import (
	"buah/entities"
	"buah/models"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	buah := models.GetAll()

	data := map[string]any{
		"buah": buah,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var buah entities.Buah

		buah.Nama = r.FormValue("buah")

		if ok := models.Create(buah); !ok {
			temp, err := template.ParseFiles("views/create.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)

		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
