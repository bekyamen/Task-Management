# Full-Stack microservices Task Management System

A robust, containerized Task Management system built with modern best practices, microservices architecture, and clean patterns.

## Tech Stack
- Frontend: React (Vite + TypeScript + Tailwind CSS)
- Backend: Go (Gin, GORM, Clean Architecture)
- DB: PostgreSQL
- Cache: Redis
- Gateway: Nginx

## Steps to Run
Requirements: Docker & Docker Compose installed.

1. Ensure no processes are holding ports `80`, `5432`, `6379`, `8080` on your host.
2. In the root of this project, run:
```bash
docker-compose up --build
```
3. Open a browser and navigate to `http://localhost:8080/`
4. The database is heavily automated. When `docker-compose up` is fired, the PostgreSQL container runs `db-init/init.sql` automatically to provision `task_db` and `user_db`.
5. Enjoy the polished web UI!

## Service Endpoints
- `user-service`: Internal `8080`, handles authentication logic (JWT, BCrypt) mapped to `/api/auth/`
- `task-service`: Internal `8080`, handles Task CRUD and Redis cache invalidation, mapped to `/api/tasks/`
- `frontend/Nginx`: Host port `8080`, serves the static files and proxies to `/api/auth` and `/api/tasks`
