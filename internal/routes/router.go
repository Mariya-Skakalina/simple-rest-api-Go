package routes

import (
	"database/sql"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/books", handlers.AllBooks(db))
	e.POST("/books", handlers.AddBook(db))
	e.GET("/book/:id", handlers.GetBook(db))
	e.PUT("/book/:id", handlers.UpdateBook(db))
	e.DELETE("/book/:id", handlers.DeleteBook(db))
}
