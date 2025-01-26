# Banking Service API

A RESTful banking service API built with Go, Fiber, GORM, and PostgreSQL. Features account management, transactions, and JWT authentication.

![Go](https://img.shields.io/badge/Go-1.23-blue)
![Fiber](https://img.shields.io/badge/Fiber-2.x-9cf)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-brightgreen)
![Docker](https://img.shields.io/badge/Docker-3.x-important)

## Features

- 💳 Account registration
- 💰 Deposit/Withdraw funds
- 🔒 JWT Authentication
- 📊 Balance inquiry
- 🐳 Dockerized deployment
- 📝 Structured logging
- 🔄 Database migrations

## Tech Stack

- **Language**: Go 1.23
- **Framework**: Fiber
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Containerization**: Docker + Docker Compose
- **Password Hashing**: bcrypt

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.23+ (for local development)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/meterai07/bank-api.git
cd banking-service
```

2. Create `.env` file:
```env
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=bank-db
DB_PORT=5432
JWT_SECRET=your_secure_secret
PORT=8080
```

3. Start services:
```bash
docker-compose up --build
```

## Environment Variables

| Variable        | Description                  | Example Value       |
|-----------------|------------------------------|---------------------|
| DB_HOST         | PostgreSQL host              | postgres            |
| DB_USER         | Database user                | postgres            |
| DB_PASSWORD     | Database password            | securepassword123   |
| DB_NAME         | Database name                | bank-db             |
| DB_PORT         | Database port                | 5432                |
| JWT_SECRET      | JWT signing key              | super_secret_key    |
| PORT            | API server port              | 8080                |

## API Documentation

### Authentication

#### Register Account
```http
POST /api/v1/daftar
```
**Request:**
```json
{
  "nama": "John Doe",
  "nik": "1234567890123456",
  "no_hp": "081234567890",
  "password": "securepass123"
}
```

#### Login
```http
POST /api/v1/login
```
**Request:**
```json
{
  "nik": "1234567890123456",
  "password": "securepass123"
}
```

### Account Management

#### Deposit Funds
```http
POST /api/v1/tabung
```
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```
**Request:**
```json
{
  "no_rekening": "REK-123456",
  "nominal": 500000
}
```

#### Check Balance
```http
GET /api/v1/saldo/REK-123456
```
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

## Project Structure

```
banking-service/
├── src/
│   ├── database/       # Database connection
│   ├── handlers/       # API handlers
│   ├── models/         # Data models
│   ├── repository/     # Database operations
│   ├── middleware/     # Authentication middleware
│   └── utils/          # Helper functions
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── main.go
```

## Deployment

### Development
```bash
docker-compose up --build
```

## Testing

Example cURL commands:

1. Register account:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"nama":"Test User","nik":"1234567890123456","no_hp":"081234567890","password":"testpass"}' \
  http://localhost:8080/api/v1/daftar
```

2. Login:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"nik":"1234567890123456","password":"testpass"}' \
  http://localhost:8080/api/v1/login
```

---

**Contributing**  
Feel free to submit issues or PRs!  
**Contact**: tengkumrafir@gmail.com
**API Demo**: [https://documenter.getpostman.com/view/20875079/2sAYQgfnq3]