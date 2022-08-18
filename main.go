package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	age        uint16
	money      int16
	avg_grades float64
	happiness  float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. "+
		"He is %d and he had money equal: %d", u.Name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}

	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8888", nil)

}

func main() {
	handleRequest()
}
