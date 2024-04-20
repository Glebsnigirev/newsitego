package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		// Если шаблоны не могут бresponseWriterыть загружены, возвращаем ошибку
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Убираем все строки, связанные с базой данных

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		// Если шаблоны не могут быть загружены, возвращаем ошибку
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Заполните все строки")
	} else {
		// Убираем все строки, связанные с базой данных

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		// Если шаблоны не могут быть загружены, возвращаем ошибку
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "show", showPost)
}

func handlefunc() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handlefunc()
}
