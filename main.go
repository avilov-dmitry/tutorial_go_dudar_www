package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Id        uint16 `json: id`
	Title     string `json: title`
	Anons     string `json: anons`
	Full_text string `json: full_text`
}

// r *http.Request - содержит информацию о пришедшем запросе

func index_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:S3cret@tcp(127.0.0.1:3306)/tutorial_go_dudar_www")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	articles := []Article{}

	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Full_text)

		if err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}

	t.ExecuteTemplate(w, "index", articles) // этот метод нужен для динамического подлючения шаблонов
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

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполненны")
	} else {
		fmt.Println(fmt.Sprintf("Recieved new form: \n %s \n %s \n %s", title, anons, full_text))

		db, err := sql.Open("mysql", "root:S3cret@tcp(127.0.0.1:3306)/tutorial_go_dudar_www")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles(title, anons, full_text) VALUES('%s', '%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func post_by_id_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show_post.html", "templates/header.html", "templates/footer.html")
	vars := mux.Vars(r)

	db, err := sql.Open("mysql", "root:S3cret@tcp(127.0.0.1:3306)/tutorial_go_dudar_www")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	post := Article{}
	for res.Next() {
		var article Article

		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Full_text)
		if err != nil {
			panic(err)
		}

		post = article
	}
	t.ExecuteTemplate(w, "post", post) // этот метод нужен для динамического подлючения шаблонов

}

func handleFunc() {
	router := mux.NewRouter() // Нужен для извлечения QP

	router.HandleFunc("/", index_page).Methods("GET")
	router.HandleFunc("/create", create_page).Methods("GET")
	router.HandleFunc("/save-article", save_article).Methods("POST")
	router.HandleFunc("/posts/{id:[0-9]+}", post_by_id_page).Methods("GET")

	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
