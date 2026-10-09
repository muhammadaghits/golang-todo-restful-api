# Golang Todo RESTful API

A robust, modular Todo List RESTful API built with **Go (Golang)**, **Gin Web Framework**, and **PostgreSQL** using `pgx/v5` connection pooling. The project features **JWT (JSON Web Token)** authentication, password hashing with **bcrypt**, database migrations using **golang-migrate**, and hot-reloading development support via **Air**.

---

## 🚀 Key Features

- **User Authentication & Authorization:**
  - User registration with validation and password hashing (bcrypt).
  - User login issuing JWT tokens (24-hour expiration).
  - JWT middleware to secure private endpoints.
- **Todo Management (CRUD):**
  - Create new todos associated with the authenticated user.
  - List all todos belonging to the authenticated user.
  - Fetch a single todo by its ID.
  - Update a todo's title and/or completion status.
  - Delete a todo by ID.
- **Clean & Modular Architecture:**
  - Clear separation of concerns: *handlers*, *repository*, *models*, *middleware*, and *config*.
  - PostgreSQL connection pooling (`pgxpool`) for high performance.
- **Developer Experience:**
  - PowerShell migration automation script (`migrate.ps1`).
  - Hot reloading during local development using Air.

---

## 🛠️ Tech Stack

- **Programming Language:** [Go](https://go.dev/) (>= 1.22)
- **Web Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Database Driver:** [pgx v5](https://github.com/jackc/pgx)
- **Authentication & Security:** [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt) & [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- **Environment Management:** [godotenv](https://github.com/joho/godotenv)
- **Database Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Live Reload Tool:** [Air](https://github.com/air-verse/air)

---

## 📁 Project Structure

```text
golang-todo-restful-api/
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/              # Configuration & environment loader (.env)
│   ├── database/            # PostgreSQL pool initialization
│   ├── handlers/            # HTTP handlers / controllers (Auth & Todo)
│   ├── middleware/          # Middlewares (JWT Authentication)
│   ├── models/              # Data structs / models (User & Todo)
│   └── repository/          # Database query layer
├── migrations/              # SQL migration files (up & down)
├── scripts/
│   └── migrate.ps1          # Migration automation script (PowerShell)
├── .air.toml                # Air hot reload configuration
├── .env.example             # Environment template
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definitions
├── go.sum                   # Go checksum file
└── README.md                # Project documentation
```

---

## ⚙️ Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://go.dev/dl/) (version 1.22 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (optional, for running migrations)
- [Air](https://github.com/air-verse/air) (optional, for hot reloading)

---

## 🏁 Getting Started & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/muhammadaghits/golang-todo-restful-api.git
cd golang-todo-restful-api
```

### 2. Download Dependencies
```bash
go mod download
```

### 3. Configure Environment Variables
Copy the `.env.example` file to `.env`:
```bash
cp .env.example .env
```

Update `.env` with your local database credentials and JWT secret:
```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/golang_todo_restful_api?sslmode=disable
PORT=3000
JWT_SECRET=your_super_secret_jwt_key
```

> **Note:** Ensure the database specified in `DATABASE_URL` (e.g., `golang_todo_restful_api`) exists in your PostgreSQL server.

### 4. Run Database Migrations

Using the provided PowerShell script:
```powershell
.\scripts\migrate.ps1 up
```

Or using the standard `golang-migrate` CLI:
```bash
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/golang_todo_restful_api?sslmode=disable" up
```

### 5. Run the Server

**Standard Mode:**
```bash
go run ./cmd/api
```

**Hot Reload Mode (Air):**
```bash
air
```

The server will be running at `http://localhost:3000`.

---

## 📚 API Endpoints

### Public Endpoints

#### 1. Health Check
- **Endpoint:** `GET /`
- **Response:**
  ```json
  {
    "database": "Connected",
    "message": "Todo API is running well!",
    "status": "Success"
  }
  ```

#### 2. Register User
- **Endpoint:** `POST /api/auth/register`
- **Body (JSON):**
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "id": "uuid-or-user-id",
    "email": "user@example.com",
    "created_at": "2026-10-09T14:00:00Z"
  }
  ```

#### 3. Login User
- **Endpoint:** `POST /api/auth/login`
- **Body (JSON):**
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

---

### Protected Endpoints
> All endpoints below require the following header:  
> `Authorization: Bearer <TOKEN>`

#### 4. Test Protected Route
- **Endpoint:** `GET /protected-test`
- **Response (200 OK):**
  ```json
  {
    "message": "Protected route accessed successfully!",
    "user_id": "uuid-user"
  }
  ```

#### 5. Create a Todo
- **Endpoint:** `POST /api/todos`
- **Body (JSON):**
  ```json
  {
    "title": "Learn Go RESTful API",
    "completed": false
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "id": 1,
    "title": "Learn Go RESTful API",
    "completed": false,
    "user_id": "uuid-user",
    "created_at": "2026-10-09T14:00:00Z",
    "updated_at": "2026-10-09T14:00:00Z"
  }
  ```

#### 6. Get All Todos
- **Endpoint:** `GET /api/todos`
- **Response (200 OK):**
  ```json
  [
    {
      "id": 1,
      "title": "Learn Go RESTful API",
      "completed": false,
      "user_id": "uuid-user",
      "created_at": "2026-10-09T14:00:00Z",
      "updated_at": "2026-10-09T14:00:00Z"
    }
  ]
  ```

#### 7. Get Todo by ID
- **Endpoint:** `GET /api/todos/:id`
- **Response (200 OK):** Detailed todo object (only if owned by the authenticated user).

#### 8. Update Todo
- **Endpoint:** `PUT /api/todos/:id`
- **Body (JSON):**
  ```json
  {
    "title": "Learn Go RESTful API & Deployment",
    "completed": true
  }
  ```
- **Response (200 OK):** Updated todo object.

#### 9. Delete Todo
- **Endpoint:** `DELETE /api/todos/:id`
- **Response (200 OK):**
  ```json
  {
    "message": "Todo deleted successfully"
  }
  ```

---

## 📜 Database Migration Scripts

The PowerShell script `.\scripts\migrate.ps1` provides simple commands to manage database schemas:
- **Apply all pending migrations:**
  ```powershell
  .\scripts\migrate.ps1 up
  ```
- **Roll back migrations:**
  ```powershell
  .\scripts\migrate.ps1 down [count]
  ```
- **Create a new migration:**
  ```powershell
  .\scripts\migrate.ps1 create migration_name
  ```
- **Force a specific migration version:**
  ```powershell
  .\scripts\migrate.ps1 force [version_number]
  ```

---

## 📄 License

This project is created for educational and personal development purposes. Feel free to use and modify it.
