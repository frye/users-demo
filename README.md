# User Profile REST API

A simple RESTful API that provides user profile information including ID, full name, and emoji.

## Features

- Get all user profiles
- Get a specific user profile by ID
- Create new user profiles
- Update existing user profiles
- Delete user profiles

## API Endpoints

- GET `/api/v1/users` - Get all users
- GET `/api/v1/users/:id` - Get a specific user by ID
- POST `/api/v1/users` - Create a new user
- PUT `/api/v1/users/:id` - Update an existing user
- DELETE `/api/v1/users/:id` - Delete a user

## Data Model

Each user profile contains:
- `id`: String identifier
- `fullName`: User's full name
- `emoji`: An emoji representing the user

## Getting Started

### Prerequisites

- Go 1.16 or newer

### Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod tidy` to ensure dependencies are correctly installed

### Running the API

```
go run main.go
```

The API will start on `http://localhost:8080`

## Example Usage

### Get all users
```
curl http://localhost:8080/api/v1/users
```

### Get user by ID
```
curl http://localhost:8080/api/v1/users/1
```

### Create a new user
```
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"id":"4", "fullName":"Alice Cooper", "emoji":"🎭"}'
```

### Update a user
```
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"fullName":"John Smith", "emoji":"😎"}'
```

### Delete a user
```
curl -X DELETE http://localhost:8080/api/v1/users/1
```
