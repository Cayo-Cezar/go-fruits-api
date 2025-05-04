# 🍍 Go Fruits API

API desenvolvida em Go que coleta dados nutricionais de frutas da [Fruityvice API](https://www.fruityvice.com/) diariamente e disponibiliza um relatório via endpoint HTTP.

---

## 🚀 Funcionalidades

- Crawler diário que importa frutas da Fruityvice API e salva no PostgreSQL
- Endpoint `/api/fruits/report-sugar` que retorna um relatório de frutas classificadas por teor de açúcar
- Classificação:
  - **High Sugar**: açúcar ≥ 10g/100g
  - **Low Sugar**: açúcar < 10g/100g

---

## ⚙️ Tecnologias utilizadas

- Go
- Gin (framework HTTP)
- GORM (ORM)
- PostgreSQL (via Docker)
- Cron jobs (`robfig/cron`)
- PgAdmin para interface visual com o banco

---

## 🛠️ Requisitos

- Go instalado (versão 1.18+)
- Docker e Docker Compose

--- 

## 📦 Como rodar o projeto

# 1. Clone o repositório
git clone https://github.com/Cayo-Cezar/go-fruits-api.git
cd go-fruits-api

# 2. Suba o banco de dados com Docker
docker-compose up -d

# 3. Instale as dependências Go
go mod tidy

# 4. Execute a aplicação
go run main.go


## 5. Testando o endpoint 

http://localhost:8080/api/fruits/report-sugar

# Exemplo de respostas
{
  "high_sugar": [
    { "id": 1, "name": "Banana" },
    { "id": 52, "name": "Persimmon" }
  ],
  "low_sugar": [
    { "id": 3, "name": "Strawberry" },
    { "id": 5, "name": "Tomato" }
  ],
  "total_high_sugar": 13,
  "total_low_sugar": 36
}

## estrutura do projeto

go-fruits-api/
├── clients/            # Client HTTP da Fruityvice
├── database/           # Conexão com PostgreSQL
├── handlers/           # Endpoints HTTP (Gin)
├── models/             # Structs de domínio (Fruit, Nutrition)
├── repositories/       # Acesso ao banco (GORM)
├── schedulers/         # Cron job diário
├── services/           # Lógica de negócio
├── main.go             # Entry point
├── docker-compose.yml
└── README.md
