# Task Management System API Reference

Here is the complete list of APIs in your application, formatted and ready for testing in Postman.

## User Service

These endpoints handle authentication and account creation. No authorization headers are needed here.

### 1. Register User
- **Method:** `POST`
- **URL:** `http://localhost:8081/api/auth/register` *(default port, check your `.env` or Docker Compose if different)*
- **Headers:**
  - `Content-Type`: `application/json`
- **Body (raw JSON):**
  ```json
  {
    "email": "user@example.com",
    "password": "mysecurepassword"
  }
  ```

### 2. Login User
- **Method:** `POST`
- **URL:** `http://localhost:8081/api/auth/login`
- **Headers:**
  - `Content-Type`: `application/json`
- **Body (raw JSON):**
  ```json
  {
    "email": "user@example.com",
    "password": "mysecurepassword"
  }
  ```
> **Note:** A successful login returns a `{ "token": "ey..." }`. Copy this token to use in the Authorization header for all Task Service endpoints.

### 3. Forgot Password
🔹 Endpoint
POST /api/auth/forgot-password


🔹 Request Body
{
  "email": "user@example.com"
}

🔹 Response (Success)
{
  "message": "If an account matches that email, a reset link was sent."
}

🔹 Internal Process (What happens in backend)
Check if user exists by email
Generate secure random token (crypto/rand)
Hash token using SHA-256
Store:
reset_token_hash
reset_token_expires_at (15 minutes)
Log reset URL (for development):
http://localhost:5174/reset-password?token=RAW_TOKEN
🔹 Security Notes
Does NOT reveal whether email exists
Token is:
Random (secure)
Hashed before storage
Token expires after 15 minutes


### 4, Reset Password
🔹 Endpoint
POST /api/auth/reset-password

🔹 Request Body
{
  "token": "abc123xyz...",
  "newPassword": "newpassword123"
}

🔹 Response (Success)
{
  "message": "Password has been successfully reset"
}

---

## Task Service

All Task endpoints require authorization using the JWT token obtained from the Login request.

### 1. Create Task
- **Method:** `POST`
- **URL:** `http://localhost:8082/api/tasks` *(default port, check your configuration)*
- **Headers:**
  - `Content-Type`: `application/json`
  - `Authorization`: `Bearer <YOUR_TOKEN_HERE>`
- **Body (raw JSON):**
  ```json
  {
    "title": "Complete project setup",
    "description": "Initialize repository and basic config",
    "status": "pending"
  }

  ```

### 2. Get All User Tasks
- **Method:** `GET`
- **URL:** `http://localhost:8082/api/tasks`
- **Headers:**
  - `Authorization`: `Bearer <YOUR_TOKEN_HERE>`

### 3. Update Task
- **Method:** `PUT`
- **URL:** `http://localhost:8082/api/tasks/:id` *(Replace `:id` with a real numeric task ID, e.g., `/api/tasks/1`)*
- **Headers:**
  - `Content-Type`: `application/json`
  - `Authorization`: `Bearer <YOUR_TOKEN_HERE>`
- **Body (raw JSON):**
  ```json
  {
    "title": "Complete project setup",
    "description": "Initialize repository and basic config (UPDATED)",
    "status": "in-progress"
  }
  ```

### 4. Delete Task
- **Method:** `DELETE`
- **URL:** `http://localhost:8082/api/tasks/:id` *(Replace `:id` with a real numeric task ID)*
- **Headers:**
  - `Authorization`: `Bearer <YOUR_TOKEN_HERE>`
