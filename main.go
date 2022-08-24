package main

import (
	_ "database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
}

// r *http.Request - содержит информацию о пришедшем запросе

func index_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil) // этот метод нужен для динамического подлючения шаблонов
}

func create_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil) // этот метод нужен для динамического подлючения шаблонов
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")         // можно забрать значение из формы, название должно быть равно name контрола
	anons := r.FormValue("anons")         // можно забрать значение из формы, название должно быть равно name контрола
	full_text := r.FormValue("full_text") // можно забрать значение из формы, название должно быть равно name контрола

}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index_page) // По умолчанию, если не будет найдено страници go откроет этот адрес
	http.HandleFunc("/create", create_page)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
