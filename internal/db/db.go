package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg config.Config) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Database connected")

	createTable()
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        author VARCHAR(100) NOT NULL,
        title VARCHAR(300) NOT NULL
    );`

	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("Table 'books' checked/created")
}
