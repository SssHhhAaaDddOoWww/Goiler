# Goiler 🛠️

Goiler is a **Go boilerplate** designed to help you start backend services quickly with a clean, scalable, and idiomatic Go project structure.

---

## 📦 Folder Structure

```
Goiler/
├── cmd/
│   └── app/
│       └── main.go        # Application entry point
│
├── internal/
│   ├── config/            # App configuration (env, constants, setup)
│   ├── handlers/          # HTTP handlers / controllers
│   ├── routes/            # Route definitions
│   ├── services/          # Business logic layer
│   └── repository/        # Database access layer
│
│
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
└── README.md              # Project documentation
```

---

## 📁 Folder Explanation

### `cmd/`
Contains the **entry point** of the application.  
Each subfolder inside `cmd` represents a runnable app.

> Why?
- Makes it easy to run multiple services from the same repo
- Keeps `main.go` isolated from business logic

---

### `internal/`
Private application code that **cannot be imported** by other projects.

#### `config/`
Handles:
- Environment variables
- App configuration
- Startup settings

#### `handlers/`
- HTTP request handlers
- Request validation
- Response formatting

#### `routes/`
- API route definitions
- Maps endpoints to handlers

#### `services/`
- Core business logic
- Keeps handlers thin
- No HTTP or DB code here

#### `repository/`
- Database queries
- ORM / SQL logic
- Data persistence layer

---


## 🚀 Getting Started

### Prerequisites
- Go 1.18+

### Run the app

```bash
go run ./cmd/app
```

### Install dependencies

```bash
go mod tidy
```

---


