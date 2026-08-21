# Project Development Guide

## 1. Install prerequisites

🔧 Install Go 1.22+, Docker, and Docker Compose.

## 2. Run the services

- 🐳 From the project root, start Postgres, run migrations, and start the API:

```bash
docker compose up -d --build
```

- 🗄️ The default Postgres username and password are `postgres` and `postgres` (see `.env.example`).

## 3. Configure local settings

📄 Copy `.env.example` to `.env`, and 🛠️ modify it as per your needs (ports, credentials, etc.):

```bash
cp .env.example .env
```

## 4. Run the API

🚀 With the stack up, the API is available at `http://localhost:${APP_PORT}` (default `9001`):

```bash
curl http://localhost:9001/api/students
```

To run the Go binary directly on the host instead of in Docker (Postgres must already be reachable and migrated):

```bash
make run
```

## 5. Understand the project architecture

🏗️ The project follows **Layered Architecture**. It separates the application into distinct layers to improve maintainability, testability, and scalability:

- **Controller**: 🎛️ Handles HTTP requests, parses/validates inputs, and sends responses. Entry point for user interaction.
- **Service**: 🛠️ Implements business logic and acts as an intermediary between controllers and repositories.
- **Repository**: 📂 Manages database interactions, encapsulating persistence logic (CRUD operations) via GORM.
- **Models / Request**: 🧩 `models/` holds the GORM entities; `request/` holds the request DTOs with validation tags.

This separation follows the **Single Responsibility Principle** and ensures a clean dependency flow: `controller → service → repository → database`.

## 6. Format code

🖋️ Run the following command after every change to format your code:

```bash
make fmt
```

## 7. Code analysis

🔍 Run the following command to perform static code analysis:

```bash
make vet
```

## 8. Manage database migrations

- 🛠️ Create new migration files (prompts for a name if `NAME` isn't given):

```bash
make migrate-new
```

- 🚀 Apply pending migrations:

```bash
make migrate-up
```

- ⏪ Roll back the last migration:

```bash
make migrate-down
```

Migration files live in `migrations/` and follow the `golang-migrate` `NNNNNN_description.up/down.sql` convention.

## 9. Explore Makefile commands

📜 The project includes a `Makefile` in the root directory. Run `make help` to list all available commands with descriptions.

## 10. Environment variables

| Variable       | Default     | Description                                      |
|----------------|-------------|---------------------------------------------------|
| `DB_HOST`      | `127.0.0.1` | Postgres host                                      |
| `DB_PORT`      | `5432`      | Postgres port                                      |
| `DB_USER`      | `postgres`  | Postgres user                                      |
| `DB_PASSWORD`  | `postgres`  | Postgres password                                  |
| `DB_NAME`      | `prashanth` | Database name                                      |
| `DB_SSLMODE`   | `disable`   | Postgres SSL mode                                  |
| `APP_PORT`     | `9001`      | Port the API listens on                            |
| `DB_HOST_PORT` | `5433`      | Docker Compose only: host-side port for Postgres   |

