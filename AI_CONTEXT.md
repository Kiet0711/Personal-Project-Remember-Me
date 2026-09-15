# Remember Me — AI Agent Context

> **Read this file before making structural changes to the project.**

## 1. Project Purpose

Remember Me is a personal task and schedule management application for users who often forget tasks, appointments, deadlines, and other important activities.

Core goals:
- Create, edit, delete, and complete tasks
- Set task priority
- Set start time, end time, and deadline
- Detect schedule conflicts
- Recommend which task should be prioritised when schedules overlap
- Create reminders and snooze/reschedule tasks
- Create recurring tasks
- Create notes
- Organise tasks with categories
- Search and filter tasks
- View tasks in a calendar
- Send notifications, with Zalo as the preferred channel
- Support Messenger and Desktop notifications
- Keep the architecture ready for a future mobile application

## 2. Technology Stack

| Layer | Technology |
|-------|------------|
| Frontend | Svelte |
| Language | TypeScript |
| Bundler | Vite |
| Desktop | Tauri |
| Backend | Golang |
| API | REST |
| Database | PostgreSQL |
| Local dev | PostgreSQL running locally, no Docker |

**Do not change the stack without explicit approval.**
Do not introduce React, Next.js, PHP, Laravel, .NET, MySQL, MongoDB, or Docker unless the user explicitly requests a change.

## 3. Architecture

```
Svelte + Tauri  →  REST API  →  Golang Backend  →  PostgreSQL
```

Rules:
- Frontend must not connect directly to PostgreSQL.
- Database access belongs to the backend.
- Business logic belongs to the backend.
- API handlers should not contain large amounts of business logic.
- Database queries should stay in the repository/database layer.
- The API should be reusable by a future mobile application.

## 4. Coding Rules

- Do not change the technology stack without explicit approval.
- Do not introduce Docker.
- Do not connect Svelte directly to PostgreSQL.
- Keep business logic in Golang.
- Keep database operations in the repository/database layer.
- Keep HTTP handling in handlers.
- Organise backend code by feature.
- Avoid unnecessary abstractions.
- Do not create duplicate entities.
- Do not create database tables for backend-calculated features.
- Check the current specification before changing the database.
- Check existing API endpoints before creating a new endpoint.
- Do not delete existing functionality without approval.
- Do not rename database fields without checking the current specification.
- Prefer simple, maintainable code.
- Do not over-engineer the application.
- Preserve compatibility with future mobile clients.

## 5. Database Entities (current)

- `users`
- `categories`
- `tasks`
- `reminders`
- `notifications`
- `notification_settings`
- `notes`
- `recurring_tasks`

**Do not create tables for:** conflicts, priority recommendations, calendar, deadlines, completed tasks, search, filters. These are computed from existing task data.

## 6. API Surface

Base prefix: `/api`

Groups: `users`, `tasks`, `categories`, `reminders`, `notifications`, `notification-settings`, `notes`, `recurring`, `schedule`, `calendar`, `deadlines`.

AI endpoints are future functionality and must not be implemented during the structure-only phase.

## 7. Current Progress

- [x] **Phase 1** — Project structure
- [ ] Phase 2 — Database / ERD
- [ ] Phase 3 — PostgreSQL migrations
- [ ] Phase 4 — Backend foundation
- [ ] Phase 5 — Database connection
- [ ] Phase 6 — Authentication
- [ ] Phase 7 — User
- [ ] Phase 8 — Category
- [ ] Phase 9 — Task
- [ ] Phase 10 — Reminder
- [ ] Phase 11 — Notification
- [ ] Phase 12 — Notes
- [ ] Phase 13 — Recurring Tasks
- [ ] Phase 14 — Schedule / Conflict
- [ ] Phase 15 — Calendar
- [ ] Phase 16 — Frontend integration
- [ ] Phase 17 — Tauri integration
- [ ] Phase 18 — Testing
- [ ] Phase 19 — Future AI / Mobile

## 8. Important Decisions

- Local PostgreSQL, no Docker (per user preference).
- Zalo is the **preferred** notification channel; Messenger and Desktop are secondary.
- Conflicts, recommendations, calendar, deadlines, completed-tasks, search and filters are all **computed** features, not database tables.
- Frontend folder name is `frontend` (lowercase) per spec — legacy folder `FrontEnd` was renamed to comply with this rule.
- The original project brief lives in `Document/RememberMe_AGENT_CONTEXT.md` (kept verbatim per user request). `AI_CONTEXT.md` is a structured digest of it for AI agents.

## 9. Known Limitations

- AI endpoints not implemented.
- Mobile client not implemented (architecture reserved).
- Notifications: only the channel interfaces are planned — actual Zalo/Messenger SDK integration is a future task.

## 10. Things the AI Must Not Change

- The technology stack (no React, no Next.js, no PHP, no Laravel, no .NET, no MySQL, no MongoDB).
- The decision to run PostgreSQL locally without Docker.
- The principle that business logic lives in Golang.
- The principle that database queries live in the repository layer.
- The migration files once created (new changes → new migration file).
- The decision that conflicts/calendar/etc. are computed, not stored.
