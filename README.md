<div align="center">

# 🔐 Login-Go

**Premium authentication system built with Go, MySQL, and native HTML templates.**

![Go](https://img.shields.io/badge/Go-1.26.2-00ADD8?style=flat-square\&logo=go\&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-MariaDB-4479A1?style=flat-square\&logo=mysql\&logoColor=white)
![Auth](https://img.shields.io/badge/Auth-Authentication-8A2BE2?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-00b894?style=flat-square)

<br>

> 🚀 A clean and modern authentication project built from scratch using Go's standard HTTP server, MySQL, bcrypt password hashing, sessions, and server-side HTML templates.

</div>

---

## ✨ Features

* 🔐 **User Login**
* 📝 **User Registration**
* 🏠 **Authenticated Dashboard**
* 🚪 **Logout**
* 👤 **User profile information**
* 🛡️ **Authentication middleware**
* 🔑 **bcrypt password hashing**
* 🗄️ **MySQL / MariaDB database**
* 🎭 **Role-based user model** (`member` / `admin`)
* 🍪 **HTTP-only session cookie**
* 🛡️ **Security headers**
* 📦 Clean separation between handlers, services, models, routes, and templates
* 🎨 Server-side HTML templates
* ⚡ Lightweight — no heavy web framework required

---

## 🧱 Project Architecture

The project follows a simple layered structure:

```text
login-go/
│
├── config/
│   └── database.go          # Database connection & environment config
│
├── database/
│   └── schema.sql           # MySQL database schema
│
├── handlers/
│   └── auth.go              # Login, register, dashboard & logout handlers
│
├── middleware/
│   └── auth.go              # Authentication middleware & security headers
│
├── models/
│   └── user.go              # User model
│
├── routes/
│   └── routes.go            # Application routes
│
├── services/
│   └── auth_service.go      # Authentication & database logic
│
├── templates/
│   ├── login.html           # Login page
│   ├── register.html        # Registration page
│   └── dashboard.html       # User dashboard
│
├── go.mod
├── go.sum
└── main.go
```

---

## 🛠️ Tech Stack

| Technology              | Purpose                 |
| ----------------------- | ----------------------- |
| **Go**                  | Backend / HTTP server   |
| **MySQL / MariaDB**     | User database           |
| **bcrypt**              | Password hashing        |
| **HTML Templates**      | Server-side UI          |
| **net/http**            | HTTP server & routing   |
| **database/sql**        | Database access         |
| **go-sql-driver/mysql** | MySQL driver            |
| **securecookie**        | Secure cookie utilities |

The project currently uses Go `1.26.2` and dependencies including `go-sql-driver/mysql`, `gorilla/securecookie`, and `golang.org/x/crypto`.

---

## 🔐 Authentication Flow

```text
                ┌──────────────┐
                │    Client    │
                └──────┬───────┘
                       │
                       ▼
                ┌──────────────┐
                │    Login     │
                │ /login       │
                └──────┬───────┘
                       │
                       ▼
              ┌─────────────────┐
              │  Auth Handler   │
              └────────┬────────┘
                       │
                       ▼
              ┌─────────────────┐
              │  Auth Service   │
              │   + bcrypt      │
              └────────┬────────┘
                       │
                       ▼
              ┌─────────────────┐
              │ MySQL / MariaDB │
              └────────┬────────┘
                       │
                       ▼
              ┌─────────────────┐
              │ Session Cookie  │
              └────────┬────────┘
                       │
                       ▼
              ┌─────────────────┐
              │    Dashboard    │
              └─────────────────┘
```

---

## 🛡️ Security

### Password Hashing

Passwords are hashed using **bcrypt** instead of storing plaintext passwords. The project currently uses bcrypt with a cost factor of `12`.

### HTTP-Only Session Cookie

Authentication state is stored through a `session_token` cookie configured with:

* `HttpOnly`
* `SameSite=Lax`
* `Path=/`
* 7-day lifetime

### Authentication Middleware

Protected requests are checked by `AuthMiddleware`. The middleware reads the session cookie, extracts the user ID, and places the authenticated user ID into the request context.

### Security Headers

The project also adds headers such as:

```text
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 0
```

---

## 🗄️ Database

The included schema creates a `makmur_mandiri` database with a `users` table.

```text
users
├── id
├── name
├── email
├── password_hash
├── role
├── last_login
├── created_at
└── updated_at
```

The email field is unique and user roles currently support:

```text
member
admin
```

---

## 🚀 Installation

### 1. Clone Repository

```bash
git clone https://github.com/ManuelKy08/login-go.git
cd login-go
```

### 2. Install Dependencies

```bash
go mod download
```

or:

```bash
go mod tidy
```

### 3. Setup Database

Make sure MySQL or MariaDB is running.

Import the schema:

```bash
mysql -u root -p < database/schema.sql
```

The schema creates:

```text
Database: makmur_mandiri
Table: users
```

### 4. Configure Environment

Set your database environment variables:

```bash
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=your_password
export DB_NAME=makmur_mandiri
export APP_PORT=8080
```

The application reads database configuration from environment variables.

### 5. Run

```bash
go run .
```

You should see:

```text
🚀 Server berjalan di http://localhost:8080
```

Open:

```text
http://localhost:8080
```

The application routes include `/login`, `/register`, `/logout`, `/dashboard`, and `/static/`.

---

## 🌐 Routes

| Method | Endpoint     | Description                      |
| ------ | ------------ | -------------------------------- |
| `GET`  | `/`          | Redirect based on authentication |
| `GET`  | `/login`     | Login page                       |
| `POST` | `/login`     | Process login                    |
| `GET`  | `/register`  | Registration page                |
| `POST` | `/register`  | Create new account               |
| `GET`  | `/dashboard` | Authenticated dashboard          |
| `GET`  | `/logout`    | Logout                           |
| `GET`  | `/static/*`  | Static assets                    |

---

## 📚 What This Project Demonstrates

This project is mainly built as a learning and portfolio project for understanding:

* Go web development
* HTTP handlers
* Routing with `net/http`
* Authentication architecture
* Password hashing
* Database integration
* Session management
* Middleware
* Server-side templates
* Environment-based configuration
* Basic web security practices
* Layered application structure

---

## ⚠️ Development Status

> **This project is intended for learning and development purposes.**

Before using it in a production environment, additional hardening should be considered, including:

* Strong cryptographically signed session tokens
* Session rotation after authentication
* CSRF protection
* Rate limiting / brute-force protection
* Secure cookie configuration with `Secure`
* Input validation
* Production error handling
* Secret management
* Database connection lifecycle management
* Comprehensive authentication testing

---

## 🔮 Roadmap

* [ ] Improve session token security
* [ ] Add CSRF protection
* [ ] Add rate limiting
* [ ] Add password reset
* [ ] Add email verification
* [ ] Add role-based authorization
* [ ] Add account profile management
* [ ] Add session expiration validation
* [ ] Add automated tests
* [ ] Add Docker support
* [ ] Add production deployment guide

---

## 🎯 Project Goal

**Login-Go** is a small authentication system created to understand how authentication works internally without relying on a large web framework.

> Learn the fundamentals.
> Build it yourself.
> Secure it properly. 🔐

---

<div align="center">

### ⭐ If this project helps you learn Go, consider giving it a star!

**Built with ❤️ using Go**

**Created by [ManuelKy08](https://github.com/ManuelKy08)**

</div>
