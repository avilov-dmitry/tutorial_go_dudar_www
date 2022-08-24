package main

import (
	"database/sql"
	"fmt"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:name`
	Age  uint16 `json:age`
}

func main() {
	db, err := sql.Open("mysql", "citizix_user:An0thrS3crt@tcp(127.0.0.1:3306)/citizix_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Установка дданных
	// insert, err := db.Query("INSERT INTO users(name, age) VALUES ('Alex', 25)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)

		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

	fmt.Println("Connected to MySQL")
}
