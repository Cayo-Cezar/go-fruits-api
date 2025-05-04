package main

import (
	"github.com/Cayo-Cezar/go-fruits-api/database"
	"github.com/Cayo-Cezar/go-fruits-api/handlers"
	"github.com/Cayo-Cezar/go-fruits-api/repositories"
	"github.com/Cayo-Cezar/go-fruits-api/schedulers"
	"github.com/Cayo-Cezar/go-fruits-api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDatabase()

	repo := repositories.NewFruitRepository(database.DB)
	service := services.NewFruitService(repo)
	service.FetchAndSaveFruits()

	schedulers.StartCron(service)

	handler := handlers.NewFruitHandler(service)
	router := gin.Default()

	router.GET("/api/fruits/report-sugar", handler.GetSugarReport)

	router.Run(":8080")
}
