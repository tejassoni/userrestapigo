# User REST API (Go)

A versioned REST API for managing users, built with Go, Gorilla Mux, and MySQL.

## Requirements

- Go 1.25.8 or later (see `go.mod`)
- MySQL 8+ (or another MySQL-compatible server)

## Quick start

1. Create your local environment file and update the database values:

   ```bash
   cp .env.example .env
   ```

2. Create the database named by `DB_NAME` (for example, `users_api`):

   ```bash
   mysql -u root -p -e 'CREATE DATABASE users_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;'
   ```

3. Set `DB_NAME=users_api` (or the name of an existing compatible database), `DB_USER`, and `DB_PASSWORD` in `.env`.

4. Apply database migrations:

   ```bash
   make migrate
   # or: go run ./cmd/migrate
   ```

5. Start the API:

   ```bash
   make run
   # or: go run ./cmd/server
   ```

The default server address is `http://localhost:8080`.

## Packages and imports

Dependencies and exact versions are declared in `go.mod`. Go downloads them automatically when you run or build the project; to download them ahead of time, run `go mod download`.

| Package | Version | Imported in | Purpose |
| --- | --- | --- | --- |
| `github.com/go-sql-driver/mysql` | `v1.10.0` | `internal/config/database.go`, `internal/errors/errors.go` | MySQL database driver and MySQL error handling. The database package imports the driver for its registration side effect. |
| `github.com/gorilla/mux` | `v1.8.1` | `internal/router/*.go`, `internal/handlers/user_handler.go` | HTTP router, path prefixes, method matching, and `{id}` path parameters. |
| `github.com/joho/godotenv` | `v1.5.1` | `internal/config/config.go` | Loads local settings from `.env`. |
| `github.com/go-playground/validator/v10` | `v10.30.3` | `internal/validator/validator.go` | Validates user and registration payloads, including custom validation rules. |
| `golang.org/x/crypto/bcrypt` | `v0.52.0` | `internal/handlers/auth_handler.go`, `internal/handlers/user_handler.go` | Hashes passwords before user records are stored. |

`go.mod` also records transitive modules marked `// indirect`. The validator and bcrypt packages above are imported by application code even though their current entries are marked indirect; the remaining indirect modules are dependencies of these packages. The project also uses Go standard-library packages such as `net/http`, `encoding/json`, `database/sql`, `log`, `time`, and `os` for HTTP handling, JSON responses, database access, logging, time parsing, and environment access.

## Server configuration

The application loads `.env` automatically on startup. In production, supply the same variables through the process environment instead; a missing `.env` file is expected when `APP_ENV=production`.

| Variable | Default | Purpose |
| --- | --- | --- |
| `APP_NAME` | `User REST API` | Displayed by the health endpoint. |
| `APP_ENV` | `development` | Application environment. |
| `APP_URL` | `http://localhost` | Used in the startup log. |
| `APP_PORT` | `8080` | TCP port on which the HTTP server listens. |
| `APP_VERSION` | `v1.0` | Displayed by the health endpoint. |
| `DB_HOST` | `127.0.0.1` | MySQL host. |
| `DB_PORT` | `3306` | MySQL port. |
| `DB_NAME` | _required_ | Database name. |
| `DB_USER` | _required_ | MySQL user. |
| `DB_PASSWORD` | empty | MySQL password. |

`DB_NAME` and `DB_USER` are required at startup. The current database connection uses the MySQL settings above and verifies connectivity before the HTTP server starts.

`.env.example` also lists settings for JWT, Redis, SMTP, CORS, logging, uploads, pagination, rate limits, and database pool tuning. They are loaded into configuration, but they are not yet connected to the running routes or database connection; configure them when adding those features.

## API overview

Base URL: `http://localhost:8080`

User CRUD is available under both `/api/v1` and `/api/v2`, with identical behavior. Registration is currently available only in v1. No authentication middleware is applied to these routes.

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/` | Health/status response. |
| `GET` | `/api/v1/users` | List all users. |
| `GET` | `/api/v1/users/{id}` | Get a user by numeric ID. |
| `POST` | `/api/v1/users` | Create a user. |
| `PUT` | `/api/v1/users/{id}` | Update a user. |
| `DELETE` | `/api/v1/users/{id}` | Delete a user. |
| `POST` | `/api/v1/auth/register` | Register a user. |
| `GET` | `/api/v2/users` | List all users (v2). |
| `GET` | `/api/v2/users/{id}` | Get a user (v2). |
| `POST` | `/api/v2/users` | Create a user (v2). |
| `PUT` | `/api/v2/users/{id}` | Update a user (v2). |
| `DELETE` | `/api/v2/users/{id}` | Delete a user (v2). |

`API_PREFIX` and `API_VERSION` do not currently alter these paths; the router registers `/api/v1` and `/api/v2` directly.

## Request payloads

Use `Content-Type: application/json` for requests with a body.

### Create a user

`POST /api/v1/users`

```json
{
  "name": "Ada Lovelace",
  "email": "ada@example.com",
  "gender": "female",
  "birthdate": "1815-12-10",
  "is_active": true,
  "password": "strong-password",
  "confirm_password": "strong-password"
}
```

### Update a user

`PUT /api/v1/users/1`

```json
{
  "name": "Ada Byron",
  "email": "ada@example.com",
  "gender": "female",
  "birthdate": "1815-12-10",
  "is_active": true
}
```

### Register a user

`POST /api/v1/auth/register`

```json
{
  "name": "Ada Lovelace",
  "email": "ada@example.com",
  "gender": "female",
  "birthdate": "1815-12-10",
  "is_active": true,
  "password": "strong-password",
  "confirm_password": "strong-password"
}
```

For create, update, and registration: `name` must be 3–100 characters, `email` must be valid, `gender` must be `male`, `female`, or `other`, `birthdate` must use `YYYY-MM-DD` (and be in the past for create/update), and passwords must be at least eight characters. Emails are unique.

## Response format and examples

Responses are JSON and generally follow this shape:

```json
{
  "status": true,
  "message": "User created successfully.",
  "data": { "id": 1 }
}
```

Common status codes are `200 OK` for successful reads, updates, and deletes; `201 Created` for successful creates/registrations; `400 Bad Request` for invalid IDs, JSON, or validation; `404 Not Found` for missing users; and `409 Conflict` for an existing email.

Try the API:

```bash
curl http://localhost:8080/

curl -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Ada Lovelace","email":"ada@example.com","gender":"female","birthdate":"1815-12-10","is_active":true,"password":"strong-password","confirm_password":"strong-password"}'

curl http://localhost:8080/api/v1/users/1
```

## Development commands

| Command | Description |
| --- | --- |
| `make run` | Start the API locally. |
| `make build` | Build `bin/server`. |
| `make test` | Run all Go tests. |
| `make vet` | Run Go static checks. |
| `make fmt` | Format Go source files. |
| `make migrate` | Run the migration command. |
| `make worker` | Start the background worker. |
| `make db-schema` | Load the schema using `MYSQL_USER`, `MYSQL_PASSWORD`, and `MYSQL_DATABASE`. |

## Database migrations

Migrations live in `cmd/migrate/migrations` and are embedded in the migration binary. Each `.sql` file is applied once, in filename order, and recorded in the `schema_migrations` table. Add a new sequentially named SQL file (for example, `002_add_user_role.sql`) for every schema change; do not edit a migration that has already been applied.

## Notes

- User passwords are hashed with bcrypt and are excluded from user response objects.
- The service must connect to MySQL successfully before it begins listening for HTTP requests.
