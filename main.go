package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "./gotut.db")
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
			);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	username := "hogdoe"
	password := "secret"
	createdAt := time.Now()
	result, err := db.Exec("INSERT INTO users (username,password, created_at) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userID)
	{
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)
		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
		err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
