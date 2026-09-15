# System Architecture

Status: **structure-only placeholder** — to be documented in Phase 4.

## High-Level Flow

```
Svelte + Tauri  →  REST API  →  Golang Backend  →  PostgreSQL
```

## Architectural Rules

- Frontend must not connect directly to PostgreSQL.
- Business logic lives in the backend.
- Database queries live in the repository layer.
- API handlers stay thin.
- The API must be reusable for a future mobile client.
