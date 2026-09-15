# Remember Me

A personal task & schedule management application for users who often forget tasks, appointments, deadlines, and other important activities.

## Features

- Create, edit, delete, and complete tasks
- Task priority, start/end time, deadline
- Detect schedule conflicts
- Recommend which task to prioritise when schedules overlap
- Reminders with snooze/reschedule
- Recurring tasks (daily/weekly/monthly/yearly)
- Notes
- Categories for organising tasks
- Search & filter
- Calendar view
- Notifications via Zalo (preferred), Messenger, Desktop
- Architecture ready for a future mobile application

## Technology Stack

| Layer | Technology |
|-------|------------|
| Frontend | Svelte + TypeScript + Vite |
| Desktop wrapper | Tauri |
| Backend | Golang (REST API) |
| Database | PostgreSQL (local, no Docker) |
| OS | Windows + VS Code |

## Project Structure

```
remember-me/
├── frontend/        Svelte + Tauri
├── backend/         Golang REST API
├── database/        SQL migrations + seeds
├── docs/            Brainstorm, ERD, API, architecture
├── scripts/         Utility scripts
├── AI_CONTEXT.md    Project context for AI agents
├── README.md
├── .env.example
└── .gitignore
```

## Development Phases

1. Project structure
2. Database / ERD
3. PostgreSQL migrations
4. Backend foundation
5. Database connection
6. Authentication
7. User
8. Category
9. Task
10. Reminder
11. Notification
12. Notes
13. Recurring Tasks
14. Schedule / Conflict
15. Calendar
16. Frontend integration
17. Tauri integration
18. Testing
19. Future AI / Mobile

## Local Setup

### Prerequisites

- PostgreSQL installed locally (no Docker)
- Go 1.21+
- Node.js + npm
- Rust (for Tauri)

### Steps

1. Create database in PostgreSQL: `remember_me`
2. Copy `.env.example` to `backend/.env` and fill values
3. Run migrations: `psql -U postgres -d remember_me -f database/migrations/*.sql`
4. Run backend: `cd backend && go run cmd/server/main.go`

## Documentation

- Project context for AI agents: `AI_CONTEXT.md`
- Architecture: `docs/architecture/system-architecture.md`
- Database design: `docs/database/database-design.md`
- API documentation: `docs/api/api-documentation.md`
