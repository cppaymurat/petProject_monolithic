package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/syllables", func(w http.ResponseWriter, r *http.Request) {
		submit := r.FormValue("submit")
		w.Write([]byte("Вы ввели: " + submit + "\n"))
		w.Write([]byte("Результат: "))
		getSyllables(submit, w)
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":8080", router)
}