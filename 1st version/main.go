package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	// name       string - с маленькой буквы приватное свойство, которое не видно в шаблонизаторе
	Name       string
	Age        uint16
	Money      int16
	avg_grades float64
	happiness  float64
	Hobbies    []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. "+
		"He is %d and he had money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	hobbies := []string{"Footbal", "Skate"}
	bob := User{"Bob", 25, -50, 4.2, 0.8, hobbies}

	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func get_info_page(w http.ResponseWriter, r *http.Request) {
	hobbies := []string{"Footbal", "Skate", "Swim"}
	bob := User{"Bob", 25, -50, 4.1, 0.7, hobbies}
	bob.setNewName("Alex")
	fmt.Fprint(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	hobbies := []string{}
	bob := User{"Alex", 25, -50, 4.2, 0.8, hobbies}

	tmpl, _ := template.ParseFiles("templates/contacts_page.html")
	tmpl.Execute(w, bob)
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/info", get_info_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8888", nil)
}

func mainFirst() {
	handleRequest()
}
