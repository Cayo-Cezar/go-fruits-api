package models

type Fruit struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Family     string    `json:"family"`
	Genus      string    `json:"genus"`
	Order      string    `json:"order"`
	Nutritions Nutrition `json:"nutritions" gorm:"embedded"`
}

type Nutrition struct {
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
	Fat           float64 `json:"fat"`
	Calories      float64 `json:"calories"`
	Sugar         float64 `json:"sugar"`
}

type FruitSummary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
