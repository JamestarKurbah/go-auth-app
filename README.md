# Auth App

## Overview
The `auth-app` is a simple authentication API built using Go. It provides endpoints for user login and token generation, utilizing the Gin framework for routing and handling HTTP requests.

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
1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd auth-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run src/main.go
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