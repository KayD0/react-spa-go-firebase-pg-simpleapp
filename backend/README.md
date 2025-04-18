# Go Backend for User Profile API

This is a Go implementation of the User Profile API backend, replacing the original Python implementation.

## Features

- User authentication using Firebase
- User profile management
- PostgreSQL database integration
- RESTful API endpoints

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Firebase project

## Environment Variables

Copy the `.env.example` file to `.env` and fill in the values:

```bash
cp .env.example .env
```

## Installation

1. Install dependencies:

```bash
go mod download
```

2. Build the application:

```bash
go build -o app
```

## Running the Application

```bash
./app
```

Or directly with Go:

```bash
go run main.go
```

## API Endpoints

- `GET /` - API information
- `POST /api/auth/verify` - Verify authentication token
- `GET /api/profile` - Get user profile
- `PUT /api/profile` - Update user profile

## Development

### Project Structure

- `main.go` - Application entry point
- `controllers/` - API route handlers
- `middleware/` - Middleware functions
- `models/` - Database models
- `services/` - Business logic and external services
