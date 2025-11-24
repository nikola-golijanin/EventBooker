# Event Booker API

A RESTful API for managing events and user registrations built with Go and the Gin framework.

## Features

- User authentication with JWT tokens
- Event creation, retrieval, updating, and deletion
- Event registration and cancellation
- Owner-based authorization for event management
- SQLite database for data persistence
- Password hashing with bcrypt

## Tech Stack

- **Go** 1.25.3
- **Gin** - Web framework
- **SQLite** - Database (go-sqlite3)
- **JWT** - Authentication (golang-jwt/jwt)
- **Bcrypt** - Password hashing

## Prerequisites

- Go 1.25.3 or higher

## Installation

1. Clone the repository
```bash
git clone <repository-url>
cd EventBooker
```

2. Install dependencies
```bash
go mod download
```

3. Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Building

```bash
go build -o event-booker.exe
./event-booker.exe
```

## API Endpoints

### Authentication

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/signup` | Create a new user account | No |
| POST | `/api/login` | Login and receive JWT token | No |

### Events

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/events` | Get all events | No |
| GET | `/api/events/:id` | Get event by ID | No |
| POST | `/api/events` | Create a new event | Yes |
| PUT | `/api/events/:id` | Update an event (owner only) | Yes |
| DELETE | `/api/events/:id` | Delete an event (owner only) | Yes |

### Registration

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/events/:id/register` | Register for an event | Yes |
| DELETE | `/api/events/:id/register` | Cancel event registration | Yes |

## Authentication

Protected endpoints require a JWT token in the `Authorization` header:

```
Authorization: <your-jwt-token>
```

Tokens are obtained by logging in via `/api/login` and expire after 2 hours.

## Request Examples

### Signup
```json
POST /api/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

### Login
```json
POST /api/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

### Create Event
```json
POST /api/events
Content-Type: application/json
Authorization: <your-jwt-token>

{
  "name": "Tech Conference 2025",
  "description": "Annual technology conference",
  "location": "San Francisco Convention Center",
  "dateTime": "2025-12-15T09:00:00Z"
}
```

## Testing

The repository includes an `api-colleciton.http` file for testing endpoints with the REST Client extension in VSCode.

## Database Schema

The application uses SQLite with three tables:

**users**
- id (INTEGER, PRIMARY KEY)
- email (TEXT, UNIQUE)
- password (TEXT, hashed)

**events**
- id (INTEGER, PRIMARY KEY)
- name (TEXT)
- description (TEXT)
- location (TEXT)
- dateTime (DATETIME)
- user_id (INTEGER, FOREIGN KEY → users.id)

**registrations**
- id (INTEGER, PRIMARY KEY)
- user_id (INTEGER, FOREIGN KEY → users.id)
- event_id (INTEGER, FOREIGN KEY → events.id)

Database file: `api.db` (auto-created on first run)

## Project Structure

```
.
├── main.go              # Application entry point
├── db/                  # Database initialization and schema
├── models/              # Data models (Event, User)
├── routes/              # HTTP route handlers
├── middlewares/         # Authentication middleware
├── utils/               # JWT and password hashing utilities
└── api-colleciton.http  # API test requests
```

## License

This project is for personal/educational use.
