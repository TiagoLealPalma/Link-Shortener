package data

import (
	"database/sql"
	"time"
)
import _ "modernc.org/sqlite"

var err error
var db *sql.DB

func InitDB() {
	db, err = sql.Open("sqlite", "./link-shortener.db")
	if err != nil {
		panic(err)
	}

	// Criação da tabela se não existir
	statement := `
		CREATE TABLE IF NOT EXISTS links (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			token TEXT NOT NULL UNIQUE,
			original_url TEXT NOT NULL,
		    date DATE NOT NULL                             
		);
		`
	_, err = db.Exec(statement)
	if err != nil {
		panic(err)
	}
}

func SaveUrl(token string, originalUrl string) error {
	println(token)
	print(originalUrl)
	date := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO links (token, original_url, date) VALUES (?, ?, ?)", token, originalUrl, date)
	return err
}

func GetOriginalLink(token string) (string, error) {
	var url string
	err := db.QueryRow("SELECT original_url FROM links WHERE token = ?", token).Scan(&url)
	return url, err
}
