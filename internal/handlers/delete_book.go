package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteBook(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := `DELETE FROM books WHERE id=$1`
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Println("Error deleting from database: ", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}
