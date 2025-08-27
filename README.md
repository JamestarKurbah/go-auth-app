# Auth App

## Overview
The `go-auth-app` is a simple authentication API built using Go. It provides endpoints for user login and token generation, utilizing the Gin framework for routing and handling HTTP requests.

## Project Structure
```
auth-app
├── src
│   ├── main.go            # Entry point of the application
│   ├── handlers           # Contains HTTP handler functions
│   │   └── handler.go
│   ├── routes             # Sets up application routes
│   │   └── routes.go
│   ├── models             # Defines data models
│   │   └── model.go
│   └── tests              # Contains unit tests for handlers
│       └── handler_test.go
├── go.mod                 # Module definition and dependencies
├── go.sum                 # Dependency checksums
└── README.md              # Project documentation
```


## Setup Instructions

### Local Development
1. **Clone the repository:**
    ```
    git clone https://github.com/JamestarKurbah/go-auth-app
    cd go-auth-app
    ```

2. **Install dependencies:**
    ```
    go mod tidy
    ```

3. **Run the application:**
    ```
    go run src/main.go
    ```

### Using Docker & Docker Compose
1. **Build and start the app and PostgreSQL database:**
    ```
    docker-compose up --build
    ```

2. **Database:**
    - The database service uses PostgreSQL and creates a database named `goauth`.
    - Default credentials: `postgres`/`postgres`.
    - The app connects using the `DATABASE_URL` environment variable.
    - Make sure the `users` table exists:
       ```sql
       CREATE TABLE users (
          id SERIAL PRIMARY KEY,
          username VARCHAR(255) UNIQUE NOT NULL,
          password_hash TEXT NOT NULL
       );
       ```

## API Endpoints

### Register a new user
`POST /register`
Request body (JSON):
```
{
   "username": "yourusername",
   "password": "yourpassword"
}
```

### Login
`POST /login`
Request body (JSON):
```
{
   "username": "yourusername",
   "password": "yourpassword"
}
```
Response:
```
{
   "token": "<JWT token>"
}
```

### Validate Token
`GET /token`
Header:
```
Authorization: Bearer <JWT token>
```
Response:
```
{
   "message": "Token is valid"
}
```

## Testing Endpoints with curl (PowerShell)

**Register:**
```
curl -Method POST -Uri http://localhost:8080/register -Body '{"username":"yourusername","password":"yourpassword"}' -ContentType "application/json"
```

**Login:**
```
curl -Method POST -Uri http://localhost:8080/login -Body '{"username":"yourusername","password":"yourpassword"}' -ContentType "application/json"
```

**Validate Token:**
```
curl -Method GET -Uri http://localhost:8080/token -Headers @{Authorization="Bearer <your_token_here>"}
```

4. **Access the API:**
   The API will be available at `http://localhost:8080`.

## Usage
- **Login Endpoint:** `POST /login`
  - Request body should contain user credentials.
  
- **Token Generation Endpoint:** `POST /token`
  - Request body should contain necessary information for token generation.

## Testing
To run the tests, use the following command:
```
go test ./src/tests
```

## Contributing
Feel free to submit issues or pull requests for improvements or bug fixes. 

## License
This project is licensed under the MIT License.