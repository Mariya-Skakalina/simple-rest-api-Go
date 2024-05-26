package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/models"
	"github.com/labstack/echo/v4"
)

func AddBook(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var book models.Book

		// Связываем данные из тела запроса с объектом book
		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if book.Author == "" && book.Title == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Author and Title are required fields"})
		}

		// SQL запрос для вставки данных и возврата ID новой записи
		sqlStatement := `INSERT INTO books (author, title) VALUES ($1,$2) RETURNING id`
		err := db.QueryRow(sqlStatement, book.Author, book.Title).Scan(&book.ID)
		if err != nil {
			log.Println("Error inserting into db: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not insert book into database"})
		}
		return c.JSON(http.StatusOK, book)
	}
}
