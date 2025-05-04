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

## Visualização de dados 

## 🔗 Acessando o Banco de Dados via pgAdmin

Para visualizar os dados no banco de dados, siga os passos abaixo:

### 1. Acesse o pgAdmin
Abra o navegador e vá até:

http://localhost:54321

Faça login utilizando o e-mail e a senha definidos no arquivo `docker-compose.yml`.  
**Exemplo:**
- **Email**: `cayo@email.com`
- **Senha**: `123456`

---

### 2. Criar uma conexão com o banco

Após o login, será necessário registrar o servidor do banco manualmente:

**Passos:**
1. Clique com o botão direito em **Servers**
2. Selecione **Create → Server...**

#### Aba "General"
- **Name**: Escolha qualquer nome (ex: `GoFruitsDB`)

#### Aba "Connection"
- **Host**:
  - No terminal do editor ou cmd, execute:
    ```bash
    docker-compose exec postgres sh
    hostname -i
    ```
  - Copie o IP retornado (ex: `172.21.0.2`) e cole no campo **Host**

- **Port**: `5432`  
- **Username**: `root`  
- **Password**: `root`  
- **Database**: `root`

---

### 3. Finalize
Clique em **Save**.  
A conexão com o banco de dados estará pronta e visível no painel esquerdo do pgAdmin.

