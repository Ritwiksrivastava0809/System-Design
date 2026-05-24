# Gatekeeper

Gatekeeper is a Go-based authentication and user management service built with Gin, GORM, and PostgreSQL. It provides a simple REST API for user creation, login, and logout, with token-based authentication support.

## Features

- Gin HTTP server with CORS enabled
- PostgreSQL persistence using GORM
- JWT/PASETO-style token generation for authentication
- User creation and login flow
- Structured configuration using Viper
- Docker Compose setup for local PostgreSQL

## Project Structure

- `cmd/api/main.go` - application entry point
- `internal/server` - router and server bootstrap
- `internal/Authentication` - authentication service, token generation, and handlers
- `internal/user` - user service, repository, and handlers
- `internal/storage/postgres` - PostgreSQL user repository
- `platform/database` - PostgreSQL connection and migration helpers
- `config` - environment configuration loader
- `environment/development.yaml` - development environment configuration
- `docker-compose.yaml` - local PostgreSQL service

## Requirements

- Go 1.25+
- Docker (for local database with Docker Compose)

## Setup

1. Start PostgreSQL locally

```sh
cd solids/Projects/Gatekeeper
make compose-up
```

2. Run the server in development mode

```sh
make server
```

Or directly:

```sh
go run cmd/api/main.go -e development
```

## Configuration

The service reads configuration from `environment/development.yaml` by default. Key settings include:

- `server.host`
- `server.port`
- `database.username`
- `database.password`
- `database.host`
- `database.port`
- `database.name`
- `database.sslmode`
- `token.internal`
- `token.symmetric`
- `token.access_token_duration`

## API Endpoints

Base path: `/api/v0`

### Health check

- `GET /api/v0/health`
- Response: `200 OK`
- Body: `{ "status": "ok" }`

### Users

- `POST /api/v0/users/`
- Create a new user
- Request body: JSON with required user fields

### Authentication

- `POST /api/v0/authenticate/login`
- Login with username and password
- Request body: JSON credentials
- `POST /api/v0/authenticate/logout`
- Logout endpoint; expects `Authorization` header

## Notes

- Database migrations run automatically on startup.
- The service currently exposes a minimal API and is designed for extension with additional user management and authentication routes.

## Useful commands

```sh
make compose-up   # start database
make server       # start API server
```

## License

This repository does not specify a license. Add one if you plan to open source the project.
