// nesse arquivo ele faz todas as operações dentro do banco

package repositories

import (
	"errors"

	"github.com/Cayo-Cezar/go-fruits-api/models"
	"gorm.io/gorm"
)

type FruitRepository struct {
	db *gorm.DB
}

func NewFruitRepository(db *gorm.DB) *FruitRepository {
	return &FruitRepository{db: db}
}

func (r *FruitRepository) CreateOrUpdate(fruit *models.Fruit) error {
	var existing models.Fruit

	err := r.db.First(&existing, fruit.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Não existe: cria nova fruta
		return r.db.Create(fruit).Error
	}

	if err != nil {

		return err
	}

	// Existe: atualiza os dados
	return r.db.Model(&existing).Updates(fruit).Error
}

func (r *FruitRepository) FindAll() ([]models.Fruit, error) {
	var fruits []models.Fruit
	err := r.db.Find(&fruits).Error
	return fruits, err
}
