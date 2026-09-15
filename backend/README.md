# Backend — Remember Me (Golang)

REST API server for the Remember Me application.

## Layout

```
backend/
├── cmd/server/main.go       # entry point
├── internal/                # feature modules
│   ├── auth/                # Phase 6
│   ├── user/                # Phase 7
│   ├── category/            # Phase 8
│   ├── task/                # Phase 9
│   ├── reminder/            # Phase 10
│   ├── notification/        # Phase 11 (with channels/zalo, channels/messenger, channels/desktop)
│   ├── note/                # Phase 12
│   ├── recurring/           # Phase 13
│   └── schedule/            # Phase 14
├── middleware/              # logger, recover, CORS
├── config/                  # env loader (.env + os.Getenv)
├── database/                # pgxpool connection (Phase 5)
├── router/                  # gin router
├── migrations/              # migration runner (Phase 5)
├── go.mod
├── go.sum
└── .env
```

## Stack

- Go 1.21+
- Gin (HTTP framework)
- pgx v5 (PostgreSQL driver)
- godotenv (.env loader)

## Run (later phases)

```sh
cd backend
cp ../.env.example .env   # fill in DB_PASSWORD and JWT_SECRET
go mod tidy
go run cmd/server/main.go
```

## Current Phase

**Phase 4 — Backend foundation (structure-only).**

The server is wired but `database.NewConnection` is intentionally a stub that
returns an error. Database wiring is implemented in **Phase 5**.
