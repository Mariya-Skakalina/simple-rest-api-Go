package main

import (
	"log"

	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/config"
	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/db"
	_ "github.com/Mariya-Skakalina/simple-rest-api-Go/internal/migrations"
	"github.com/Mariya-Skakalina/simple-rest-api-Go/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pressly/goose/v3"
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

	err = goose.Up(db.DB, "./internal/migrations")
	if err != nil {
		panic(err)
	}

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))

}
