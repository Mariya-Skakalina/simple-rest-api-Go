package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/models"
	"github.com/labstack/echo/v4"
)

func UpdateBook(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var book models.Book
		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		if book.Author == "" || book.Title == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Author or Title are required fields"})
		}

		sqlStatement := `UPDATE books SET author=$1, title=$2 WHERE id=$3`

		_, err := db.Exec(sqlStatement, book.Author, book.Title, id)
		if err != nil {
			log.Println("error updating db: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Updating error"})
		}
		return c.JSON(http.StatusOK, book)
	}
}
