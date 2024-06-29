package main

import (
	"log"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/config"
	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/db"
	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	// Инициализация базы данных
	db.InitDB(cfg)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Инициализация маршрутов
	routes.InitRoutes(e, db.DB)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
