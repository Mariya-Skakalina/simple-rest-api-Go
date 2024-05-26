package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/models"
	"github.com/labstack/echo/v4"
)

func AllBooks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT id, author, title FROM books")
		if err != nil {
			log.Println("Error querying database:", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer rows.Close()

		books := []models.Book{}
		for rows.Next() {
			var book models.Book
			if err := rows.Scan(&book.ID, &book.Author, &book.Title); err != nil {
				log.Println("Error scanning row:", err)
				return c.JSON(http.StatusInternalServerError, err)
			}
			books = append(books, book)
		}
		return c.JSON(http.StatusOK, books)
	}
}
