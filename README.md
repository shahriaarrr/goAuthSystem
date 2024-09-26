# Go JWT Authentication System

This is a simple JWT (JSON Web Token) authentication system built using Golang. The project structure follows the MVC architecture with controllers, models, middleware, and initializers for better organization.

## Features

- **JWT Authentication**: Secure login with JWT.
- **Controllers**: Handles user authentication (login, signup) and protected routes.
- **Middleware**: JWT-based authorization for protecting endpoints.
- **Environment Configurations**: Uses `.env` file for managing secrets.
- **Golang Modules**: Clean module management using `go.mod` and `go.sum`.

## Project Structure

```bash
├── controllers       # Handlers for authentication routes
├── initializers      # Initialization files for database and environment variables
├── middleware        # Middleware to protect routes with JWT
├── models            # Database models for the application
├── .env              # Environment variables (DO NOT share this file publicly)
├── .gitignore        # Ignore unnecessary files for git tracking
├── Authorization.postman_collection.json   # Postman collection for testing API
├── go.mod            # Go module file
├── go.sum            # Dependencies for the Go module
├── LICENSE           # License for the project
├── main.go           # Entry point of the application
└── README.md         # Project documentation (this file)
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.19 or later installed on your machine.
- [Postman](https://www.postman.com/) or [cURL](https://curl.se/) to test the API endpoints.
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/shahriaarrr/goAuthSystem.git
   cd go-jwt-auth
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file and add your JWT secret key and database credentials:
   ```bash
   PORT=3000
   DB="host=localhost user={PUT_YOUR_USERNAME} password={PUT_YOUR_PASSWORD} dbname={PUT_YOUR_DB_NAME} port=5432 sslmode=disable"
   SECRET={YOUR_SECRET_KEY}
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

### Usage

1. **Signup**: Create a new user by sending a POST request to `/signup` with `email` and `password`.
2. **Login**: Authenticate an existing user by sending a POST request to `/login`.
3. **Protected Route**: Access a protected route by sending a request to `/validate` with a valid JWT in the Authorization header.

### Postman Collection

The project includes a Postman collection file `Authorization.postman_collection.json` which contains pre-configured requests for testing the API.

To use the collection:

1. Open Postman.
2. Import the collection by selecting the Authorization.postman_collection.json file from the project.
3. You will see three endpoints: signup, login, and validate.

## API Documentation

The following endpoints are available for the authentication system. You can use [Postman](https://www.postman.com/) or `curl` to interact with the API.

### Signup

**Endpoint**: `POST /signup`  
**Description**: Register a new user by providing an email and password.

**Request**:
```json
POST http://localhost:3000/signup
Content-Type: application/json

{
    "email": "example@example.com",
    "password": "yourpassword"
}
```

**Response**:
```json
{
    "message": "User created successfully"
}
```

### Login

**Endpoint**: `POST /login`  
**Description**: Authenticate a user and receive a JWT token.

**Request**:
```json
POST http://localhost:3000/login
Content-Type: application/json

{
    "email": "example@example.com",
    "password": "yourpassword"
}
```

**Response**:
```json
{
    "message": "welcome back :)"
}
```
*Note: Upon successful login, a JWT token will be set in the response as an HTTP-only cookie.*

### Validate Token

**Endpoint**: `GET /validate`  
**Description**: Validate a JWT token and access protected routes.

**Request**:
```http
GET http://localhost:3000/validate
Authorization: Bearer your_jwt_token
```

**Response**:
```json
{
    "message": User_data
}
```

## Contributing

Feel free to open an issue or submit a pull request if you have suggestions for improving the project. Contributions are welcome!
