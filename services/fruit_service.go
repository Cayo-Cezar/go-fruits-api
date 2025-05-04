package services

import (
	"fmt"

	"github.com/Cayo-Cezar/go-fruits-api/clients"
	"github.com/Cayo-Cezar/go-fruits-api/models"
	"github.com/Cayo-Cezar/go-fruits-api/repositories"
)

type FruitService struct {
	repo *repositories.FruitRepository
}

func NewFruitService(repo *repositories.FruitRepository) *FruitService {
	return &FruitService{repo: repo}
}

// busca da API e salva no banco
func (s *FruitService) FetchAndSaveFruits() error {
	fruits, err := clients.FetchFruitsFromAPI()
	if err != nil {
		return fmt.Errorf("erro ao buscar frutas da API: %w", err)
	}

	for _, fruit := range fruits {
		if err := s.repo.CreateOrUpdate(&fruit); err != nil {
			return fmt.Errorf("erro ao salvar fruta '%s': %w", fruit.Name, err)
		}
	}

	fmt.Printf("%d frutas atualizadas com sucesso\n", len(fruits))
	return nil
}

// relatório de açúcar
type SugarReport struct {
	HighSugar      []models.FruitSummary `json:"high_sugar"`
	LowSugar       []models.FruitSummary `json:"low_sugar"`
	TotalHighSugar int                   `json:"total_high_sugar"`
	TotalLowSugar  int                   `json:"total_low_sugar"`
}

// Summary com apenas id e name
type FruitSummary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// separa frutas com açúcar >= 10 ou < 10
func (s *FruitService) GenerateSugarReport() (*SugarReport, error) {
	fruits, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar frutas: %w", err)
	}

	var high, low []models.FruitSummary

	for _, fruit := range fruits {
		summary := models.FruitSummary{ID: fruit.ID, Name: fruit.Name}

		if fruit.Nutritions.Sugar >= 10 {
			high = append(high, summary)
		} else {
			low = append(low, summary)
		}
	}

	report := &SugarReport{
		HighSugar:      high,
		LowSugar:       low,
		TotalHighSugar: len(high),
		TotalLowSugar:  len(low),
	}

	return report, nil
}
