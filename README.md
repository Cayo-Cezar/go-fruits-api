# Go Fruits API

API developed in Go that daily collects nutritional data about fruits from the [Fruityvice API](https://www.fruityvice.com/) and exposes a report through an HTTP endpoint.

---

## Features

- Daily crawler that imports fruits from the Fruityvice API and saves them to PostgreSQL
- `/api/fruits/report-sugar` endpoint that returns a report of fruits classified by sugar content
- Classification:
  - **High Sugar**: sugar ≥ 10g/100g
  - **Low Sugar**: sugar < 10g/100g

---

## Technologies Used

- Go
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL (via Docker)
- Cron jobs (`robfig/cron`)
- PgAdmin for a visual database interface

---

## Requirements

- Go installed (version 1.18+)
- Docker and Docker Compose

--- 

## How to Run the Project

# 1. Clone the repository
git clone https://github.com/Cayo-Cezar/go-fruits-api.git  
cd go-fruits-api

# 2. Start the database with Docker
docker-compose up -d

# 3. Install Go dependencies
go mod tidy

# 4. Run the application
go run main.go

## 5. Testing the endpoint

http://localhost:8080/api/fruits/report-sugar

# Response example
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

## 6. Data Visualization

# Accessing the Database via pgAdmin

To view the data in the database, follow the steps below:

### 1. Access pgAdmin
Open your browser and go to:

http://localhost:54321

Log in using the email and password defined in the `docker-compose.yml` file.  
**Example:**
- **Email**: `cayo@email.com`
- **Password**: `123456`

---

### 2. Create a database connection

After logging in, you will need to manually register the database server:

**Steps:**
1. Right-click on **Servers**
2. Select **Create → Server...**

#### "General" tab
- **Name**: Choose any name (e.g., `GoFruitsDB`)

#### "Connection" tab
- **Host**:
  - In the editor terminal or command prompt, run:
    ```bash
    docker-compose exec postgres sh
    hostname -i
    ```
  - Copy the returned IP (e.g., `172.21.0.2`) and paste it into the **Host** field

- **Port**: `5432`  
- **Username**: `root`  
- **Password**: `root`  
- **Database**: `root`

---

### 3. Finish
Click **Save**.  
The database connection will be ready and visible in the left panel of pgAdmin.
