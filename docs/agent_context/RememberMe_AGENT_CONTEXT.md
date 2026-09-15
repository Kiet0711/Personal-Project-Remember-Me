Remember Me --- Agent Project Structure

1. Project Purpose

Remember Me is a personal task and schedule management application for
users who often forget tasks, appointments, deadlines, and other
important activities.

Core goals: - Create, edit, delete, and complete tasks. - Set task
priority. - Set start time, end time, and deadline. - Detect schedule
conflicts. - Recommend which task should be prioritised when schedules
overlap. - Create reminders and snooze/reschedule tasks. - Create
recurring tasks. - Create notes. - Organise tasks with categories. -
Search and filter tasks. - View tasks in a calendar. - Send
notifications, with Zalo as the preferred channel. - Support Messenger
and Desktop notifications. - Keep the architecture ready for a future
mobile application.

This file is an instruction for an AI coding agent. Read it before
making structural changes.

2. Technology Stack

Frontend

Svelte

TypeScript

Vite

Desktop

Tauri

Backend

Golang

REST API

Database

PostgreSQL

Development

Windows

VS Code

Local development

PostgreSQL running locally

No Docker

Do not change the stack without explicit approval

Do not introduce React, Next.js, PHP, Laravel, .NET, MySQL, MongoDB, or
Docker unless the user explicitly requests a change.

3. Architecture

Svelte + Tauri
|
| REST API
v
Golang Backend
|
| SQL
v
PostgreSQL

Rules: - Frontend must not connect directly to PostgreSQL. - Database
access belongs to the backend. - Business logic belongs to the
backend. - API handlers should not contain large amounts of business
logic. - Database queries should stay in the repository/database
layer. - The API should be reusable by a future mobile application.

4. Root Structure

Use a single repository:

remember-me/
├── frontend/
├── backend/
├── database/
├── docs/
├── scripts/
├── AI_CONTEXT.md
├── README.md
├── .env.example
└── .gitignore

Do not create unnecessary folders or placeholder files.

5. Frontend Structure

Target structure:

frontend/
├── src/
│ ├── components/
│ │ ├── common/
│ │ ├── task/
│ │ ├── reminder/
│ │ ├── note/
│ │ ├── calendar/
│ │ └── notification/
│ │
│ ├── pages/
│ │ ├── Login/
│ │ ├── Register/
│ │ ├── Dashboard/
│ │ ├── Tasks/
│ │ ├── Calendar/
│ │ ├── Notes/
│ │ ├── Reminders/
│ │ └── Settings/
│ │
│ ├── services/
│ │ └── api/
│ ├── stores/
│ │ ├── auth/
│ │ ├── task/
│ │ ├── notification/
│ │ └── settings/
│ ├── types/
│ ├── utils/
│ └── routes/
│
├── src-tauri/
├── package.json
└── vite.config.ts

Create feature folders when they are actually needed. Do not create
dozens of empty files.

6. Frontend Responsibilities

Frontend is responsible for: - UI. - User input. - Calling REST APIs. -
Displaying API results. - Displaying validation errors. - Displaying
task priority/status. - Displaying reminders and notifications. -
Displaying calendar data. - Displaying conflicts and recommendations.

Frontend must not: - Connect directly to PostgreSQL. - Contain database
queries. - Reimplement backend data-integrity rules unnecessarily.

7. Backend Structure

Use feature-oriented Golang modules:

backend/
├── cmd/
│ └── server/
│ └── main.go
│
├── internal/
│ ├── auth/
│ ├── user/
│ ├── task/
│ ├── category/
│ ├── reminder/
│ ├── notification/
│ ├── note/
│ ├── recurring/
│ └── schedule/
│
├── middleware/
├── config/
├── database/
├── router/
├── migrations/
├── .env
├── go.mod
└── go.sum

For a normal feature such as Task:

internal/task/
├── handler.go
├── service.go
├── repository.go
├── model.go
├── dto.go
└── validation.go

Flow:

HTTP Request
|
v
handler.go
|
v
service.go
|
v
repository.go
|
v
PostgreSQL

8. Backend Module Responsibilities

auth

Register.

Login.

Logout.

Authentication.

Password hashing.

Token/session validation.

user

User profile.

User information.

Timezone.

User-related settings.

task

Task CRUD.

Priority.

Status.

Start/end time.

Deadline.

Location.

Category.

Complete/restore.

category

Category CRUD.

List categories.

Find tasks by category.

reminder

Create/update/delete reminders.

Enable/disable reminders.

Track reminder status.

Process reminder times.

notification

Create notifications.

Send notifications.

Track notification status.

Notification history.

Recommended structure:

internal/notification/
├── handler.go
├── service.go
├── repository.go
├── model.go
└── channels/
├── interface.go
├── zalo.go
├── messenger.go
└── desktop.go

Zalo is the preferred channel. Messenger and Desktop are secondary.

note

Create/update/delete notes.

List notes.

Get notes for a task.

Support independent notes.

recurring

Daily/weekly/monthly/yearly recurrence.

Repeat interval.

Repeat days.

Start/end date.

Enable/disable recurrence.

schedule

This is a core Remember Me feature.

Responsible for: - Detecting task conflicts. - Finding free time. -
Checking schedules before creating/updating tasks. - Recommending which
task should be prioritised. - Supporting rescheduling recommendations.

Do not create a conflicts database table in the current design.
Conflicts are calculated from task schedule data.

9. Database Entities

Current main entities:

users
categories
tasks
reminders
notifications
notification_settings
notes
recurring_tasks

General relationships:

Users 1 ---- N Tasks
Users 1 ---- N Categories
Users 1 ---- N Notes
Users 1 ---- N NotificationSettings
Users 1 ---- N Notifications

Categories 1 ---- N Tasks

Tasks 1 ---- N Reminders
Tasks 1 ---- N Notes
Tasks 1 ---- 0..1 RecurringTasks

Reminders 1 ---- N Notifications

Task is the central entity.

Do not create separate tables for: - conflicts - priority
recommendations - calendar - deadlines - completed tasks - search -
filters

These are handled using existing task data and backend logic.

10. Database Rules

users

Store: - id - name - email - phone number - password hash - timezone -
create date - update time

Never store plain-text passwords.

categories

Each category belongs to a user.

tasks

Current approved fields: - id - user_id - task_name - description -
start_time - end_time - deadline - priority - status - category_id -
location - create_at - update_date

Do not invent additional database fields unless the approved
specification requires them.

reminders

A reminder belongs to a task.

notifications

A notification may reference a user, task, and reminder.

Possible status values: - Pending - Sent - Failed

notification_settings

Stores the user's notification channel settings.

Example:

Zalo -> enabled
Messenger -> disabled
Desktop -> enabled

notes

A note belongs to a user and may optionally belong to a task.

task_id can be NULL.

recurring_tasks

Stores the recurrence rule for a task.

Current relationship:

Task 1 ---- 0..1 RecurringTask

11. Database Migration Structure

Use migration files instead of one large SQL file:

database/
├── migrations/
│ ├── 001_create_users.sql
│ ├── 002_create_categories.sql
│ ├── 003_create_tasks.sql
│ ├── 004_create_reminders.sql
│ ├── 005_create_notifications.sql
│ ├── 006_create_notification_settings.sql
│ ├── 007_create_notes.sql
│ └── 008_create_recurring_tasks.sql
│
└── seeds/
└── development.sql

If a new database field is needed later, create a new migration instead
of silently changing an old migration.

12. API Structure

Base prefix:

/api

Main groups:

/api/users
/api/tasks
/api/categories
/api/reminders
/api/notifications
/api/notification-settings
/api/notes
/api/recurring
/api/schedule
/api/calendar
/api/deadlines

Core endpoints:

POST /api/users/register
POST /api/users/login
POST /api/users/logout
GET /api/users/me
PUT /api/users/me
PUT /api/users/password

POST /api/tasks
GET /api/tasks
GET /api/tasks/{id}
PUT /api/tasks/{id}
DELETE /api/tasks/{id}
PATCH /api/tasks/{id}/complete
PATCH /api/tasks/{id}/status
PATCH /api/tasks/{id}/priority

POST /api/categories
GET /api/categories
PUT /api/categories/{id}
DELETE /api/categories/{id}

POST /api/tasks/{id}/reminders
GET /api/tasks/{id}/reminders
PUT /api/reminders/{id}
DELETE /api/reminders/{id}
PATCH /api/reminders/{id}/status

POST /api/notifications/send
GET /api/notifications
GET /api/notifications/{id}
PATCH /api/notifications/{id}/status

GET /api/notification-settings
POST /api/notification-settings
PUT /api/notification-settings/{id}
DELETE /api/notification-settings/{id}

POST /api/notes
GET /api/notes
GET /api/notes/{id}
PUT /api/notes/{id}
DELETE /api/notes/{id}
GET /api/tasks/{id}/notes

POST /api/tasks/{id}/recurring
GET /api/tasks/{id}/recurring
PUT /api/recurring/{id}
DELETE /api/recurring/{id}
PATCH /api/recurring/{id}/status

GET /api/schedule/conflicts
GET /api/schedule/free-time
POST /api/schedule/check
GET /api/schedule/recommendation

AI endpoints are future functionality and should not be implemented
during the structure-only phase.

13. Documentation

Use:

docs/
├── brainstorm/
├── database/
│ ├── ERD.png
│ ├── database-design.md
│ └── relationships.md
├── api/
│ ├── api-brainstorm.md
│ └── api-documentation.md
├── architecture/
│ └── system-architecture.md
└── testing/

Keep the original brainstorm as project documentation.

Do not silently rewrite project requirements.

14. AI_CONTEXT.md

Keep AI_CONTEXT.md at the project root.

It should contain: - Project purpose. - Technology stack. -
Architecture. - Database structure. - API structure. - Coding rules. -
Current progress. - Important decisions. - Known limitations. - Things
the AI must not change.

The AI agent should read this file before making major changes.

15. Environment

Use environment variables for backend configuration.

Example:

APP_ENV=development
SERVER_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_NAME=remember_me
DB_USER=postgres
DB_PASSWORD=

JWT_SECRET=

Never commit real secrets.

Commit .env.example, not .env.

16. Coding Rules

Do not change the technology stack without explicit approval.

Do not introduce Docker.

Do not connect Svelte directly to PostgreSQL.

Keep business logic in Golang.

Keep database operations in the repository/database layer.

Keep HTTP handling in handlers.

Organise backend code by feature.

Avoid unnecessary abstractions.

Do not create duplicate entities.

Do not create database tables for backend-calculated features.

Check the current specification before changing the database.

Check existing API endpoints before creating a new endpoint.

Do not delete existing functionality without approval.

Do not rename database fields without checking the current
specification.

Prefer simple, maintainable code.

Do not over-engineer the application.

Preserve compatibility with future mobile clients.

18. Development Phases Status

| Phase | Status | Notes |
|---|---|---|
| 1. Project structure | ✅ Complete | |
| 2. Database / ERD | ✅ Complete | |
| 3. PostgreSQL migrations | ✅ Complete | |
| 4. Backend foundation | ✅ Complete | |
| 5. Database connection | ✅ Complete | pgxpool wired, ping verified |
| 6. Authentication | ✅ Complete | register/login/logout + JWT + bcrypt + auth middleware |
| 7. User | ⬜ Pending | stub exists; full impl in Phase 7 |
| 8. Category | ⬜ Pending | |
| 9. Task | ⬜ Pending | |
| 10. Reminder | ⬜ Pending | |
| 11. Notification | ⬜ Pending | |
| 12. Notes | ⬜ Pending | |
| 13. Recurring Tasks | ⬜ Pending | |
| 14. Schedule / Conflict | ⬜ Pending | |
| 15. Calendar | ⬜ Pending | |
| 16. Frontend integration | ⬜ Pending | |
| 17. Tauri integration | ⬜ Pending | |
| 18. Testing | ⬜ Pending | |
| 19. Future AI / Mobile | ⬜ Pending | |

19. Development Phases

Build in this order:

1. Project structure
   ↓
2. Database / ERD
   ↓
3. PostgreSQL migrations
   ↓
4. Backend foundation
   ↓
5. Database connection
   ↓
6. Authentication
   ↓
7. User
   ↓
8. Category
   ↓
9. Task
   ↓
10. Reminder
    ↓
11. Notification
    ↓
12. Notes
    ↓
13. Recurring Tasks
    ↓
14. Schedule / Conflict
    ↓
15. Calendar
    ↓
16. Frontend integration
    ↓
17. Tauri integration
    ↓
18. Testing
    ↓
19. Future AI / Mobile

20. Structure-Only Agent Task

When first creating the project structure:

Read this document.

Inspect the existing project directory.

Do not overwrite existing code without checking it.

Create only missing base directories and essential files.

Do not implement business features yet.

Do not create unnecessary placeholder files.

Do not install unrelated dependencies.

Do not change the technology stack.

Do not create Docker configuration.

Report what was created.

Report any conflict between the existing project and this
specification.

Ask for approval before making architectural changes.

The first task is structure only.

Do not implement: - Authentication - Task CRUD - Notifications - AI -
Mobile - Business logic

until the user explicitly asks for the next phase.

21. Target Initial Structure

After the structure-only phase:

remember-me/
├── frontend/
│ ├── src/
│ ├── src-tauri/
│ ├── package.json
│ └── vite.config.ts
│
├── backend/
│ ├── cmd/
│ │ └── server/
│ │ └── main.go
│ ├── internal/
│ ├── middleware/
│ ├── config/
│ ├── database/
│ ├── router/
│ ├── migrations/
│ ├── go.mod
│ └── go.sum
│
├── database/
│ ├── migrations/
│ └── seeds/
│
├── docs/
│ ├── brainstorm/
│ ├── database/
│ ├── api/
│ ├── architecture/
│ └── testing/
│
├── scripts/
├── AI_CONTEXT.md
├── README.md
├── .env.example
└── .gitignore

This is the target structure. Keep it simple and expand it only when
implementation requires it.
