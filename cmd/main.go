package main

import (
	"crud_api/pkg/db"
	"crud_api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db.InitDB()

	// Инициализация Gin роутера
	router := gin.Default()

	// Подключение маршрутов
	routes.RegisterRoutes(router)

	// Запуск сервера
	router.Run("localhost:8090")
}
