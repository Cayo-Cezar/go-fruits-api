// Esse arquivo busca os dados da API

package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cayo-Cezar/go-fruits-api/models"
)

const apiURL = "https://www.fruityvice.com/api/fruit/all"

func FetchFruitsFromAPI() ([]models.Fruit, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var fruits []models.Fruit
	if err := json.Unmarshal(body, &fruits); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal do JSON: %w", err)
	}

	return fruits, nil
}
