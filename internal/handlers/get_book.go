package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/models"
	"github.com/labstack/echo/v4"
)

func GetBook(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := `SELECT id, author, title FROM books WHERE id=$1`
		row := db.QueryRow(sqlStatement, id)

		var book models.Book
		err := row.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, "Book not found")
			}
			log.Println("Error querying database:", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, book)
	}
}
