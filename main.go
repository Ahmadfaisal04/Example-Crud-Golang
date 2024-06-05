package main

import (
	"buah/config"
	"buah/controllers"
	"net/http"
)

func main() {
	config.Conn()
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8000", nil)
}
