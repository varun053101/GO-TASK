# Go User Management API

This project is a simple REST API built using Go and Fiber.
It provides CRUD operations for managing users.
User age is calculated dynamically using date of birth.

---

## Prerequisites

- Go installed
- Docker and Docker Compose installed

---

## Clone the Repository

Clone the repository and move into the project directory:

    git clone https://github.com/varun053101/GO-TASK.git
    cd GO-TASK

---

## Environment Setup

Create a `.env` file in the project root with the following values:

    SERVER_PORT=8080
    DB_HOST=localhost
    DB_PORT=5433
    DB_USER=postgres
    DB_PASSWORD=8904389396
    DB_NAME=go-task

Important:
Make sure the DB_PORT matches the PostgreSQL port exposed in
docker-compose.yml. If the port does not match, the application
will fail to connect to the database.

---

## Running the Application

Step 1: Start PostgreSQL using Docker

    docker compose up -d

Make sure the database container is running before starting the server.

Step 2: Run the Go server

    go run cmd/server/main.go

The server will start on:

    http://localhost:8080

---

## API Endpoints

Create User  
POST /users

Request body:

    {
      "name": "Alice",
      "dob": "1990-05-10"
    }

---

Get User by ID  
GET /users/:id

---

Update User  
PUT /users/:id

Request body:

    {
      "name": "Alice Updated",
      "dob": "1991-03-15"
    }

---

Delete User  
DELETE /users/:id

Returns 204 No Content.

---

List All Users  
GET /users

Returns all users with calculated age.

---

Health Check  
GET /health

Used to verify that the server is running.
