🚀 Full-Stack Microservices Task Management System

A containerized Task Management System built using a microservices architecture with modern backend practices such as JWT authentication, bcrypt password hashing, Redis caching, and Docker-based deployment.

🧰 Tech Stack
Frontend: React (Vite + TypeScript + Tailwind CSS)
Backend: Go (Gin, GORM, Clean Architecture)
Database: PostgreSQL
Cache: Redis
Containerization: Docker & Docker Compose
Reverse Proxy (optional extension): Nginx-ready architecture
🏗 System Architecture

The system is divided into independent microservices:

👤 User Service

User registration & login
JWT authentication
Forgot password & reset password flow
Password hashing using bcrypt
Token versioning (for session invalidation)
📋 Task Service
Create, read, update, delete tasks
Redis caching support (performance optimization)
JWT-protected endpoints
🗄 Database (PostgreSQL)
Stores user and task data
Auto-initialized using db-init/init.sql
⚡ Redis
Caching layer
Token version tracking for authentication security

🌐 Frontend
React UI (Vite)
Communicates with backend APIs
▶️ How to Run the Project
📦 Requirements
Docker
Docker Compose
🚀 Start the Application
docker compose up --build
🛑 Stop the Application
docker compose down
🌐 Service URLs

Service	URL
Frontend	http://localhost:5174

User Service	http://localhost:8081

Task Service	http://localhost:8082

PostgreSQL	localhost:5433
Redis	localhost:6379
⚙️ Docker Configuration Notes

PostgreSQL initializes automatically using:

db-init/init.sql
Internal Docker communication:
Database → db:5432
Redis → redis:6379
External ports:
PostgreSQL → 5433
Redis → 6379

🔐 Authentication Flow
1. Register
User is created with bcrypt-hashed password
2. Login
Returns JWT token
3. Forgot Password
Secure random token generated
Token is hashed (SHA-256)
Stored with 15-minute expiration
Reset link logged in backend (dev mode)
4. Reset Password
Token validated
Password updated
Token invalidated
User token_version incremented
🔌 API Endpoints
👤 User Service

Base URL: http://localhost:8081/api/auth

Method	Endpoint	Description
POST	/register	Register user
POST	/login	Login user
POST	/forgot-password	Send reset token
POST	/reset-password	Reset password

📋 Task Service

Base URL: http://localhost:8082/api/tasks

Method	Endpoint	Description
GET	/	Get tasks
POST	/	Create task
PUT	/:id	Update task
DELETE	/:id	Delete task

🧠 Key Features
Microservices architecture
JWT authentication
Secure password reset system
bcrypt password hashing
Redis caching layer
Token versioning (session invalidation)
Dockerized full-stack deployment
🔒 Security Highlights

Passwords hashed using bcrypt
Reset tokens hashed using SHA-256
Tokens expire in 15 minutes
JWT invalidation using token versioning
Redis used for distributed session control



📌 Quick Start Summary
docker compose up --build

Then open:

👉 http://localhost:5174