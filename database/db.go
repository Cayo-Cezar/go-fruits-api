package database

import (
	"fmt"
	"log"

	"github.com/Cayo-Cezar/go-fruits-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	err = database.AutoMigrate(&models.Fruit{})
	if err != nil {
		log.Fatal("Erro ao fazer AutoMigrate: ", err)
	}

	// Armazena a conexão em uma variável global
	DB = database
	fmt.Println("Banco de dados conectado com sucesso.")
}
