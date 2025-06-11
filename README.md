# RBAC Hospital Portal

This is a simple Golang web application with Role-Based Access Control (RBAC) for managing a hospital portal with two roles: Receptionist and Doctor.

## Features

- Single login API for both Receptionist and Doctor
- Receptionist can:
  - Register a new patient
  - View all patients
  - Update patient details
  - Delete patient records
- Doctor can:
  - View all registered patients
  - Update patient medical information
- JWT-based authentication

## Technologies Used

- **Go Programming Language:** The core language used to develop the application.
- **Gin Web Framework:** A lightweight web framework for building web applications in Go.
- **GORM:** An Object-Relational Mapping (ORM) library used for interacting with the database.
- **PostgreSQL:** The database used to store information and roles.

## Project Structure
The project follows a structured design with directories for controllers, models, middlewares, and routes.
```
├── go.mod
├── go.sum
├── internal
│   ├── controllers
│   │   ├── auth
│   │   │   ├── login_controller.go
│   │   │   ├── login_controller_test.go
│   │   │   ├── register_controller.go
│   │   │   └── register_controller_test.go
│   │   ├── doctor
│   │   │   ├── doctor_controller.go
│   │   │   └── doctor_controller_test.go
│   │   └── receptionist
│   │       ├── receptionist_controller.go
│   │       └── receptionist_controller_test.go
│   ├── database
│   │   ├── database.go
│   │   └── database_test.go
│   ├── middlewares
│   │   ├── auth_middleware.go
│   │   ├── auth_middleware_test.go
│   │   ├── role_middleware.go
│   │   └── role_middleware_test.go
│   ├── models
│   │   ├── patient.go
│   │   ├── patient_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   └── routes
│       └── routes.go
└── main.go
```
## API Documentation

You can explore all the APIs via Postman:
```
https://www.postman.com/docking-module-technologist-38792832/workspace/public/collection/42460310-2d6c996b-7937-4854-aa0f-2bc2b174c3db?action=share&creator=42460310
```

## Installation

### 1. Clone the Repository
```
git clone https://github.com/krushalgopale/Rbac-Portal.git
cd Rbac-Portal
```
### 2. Install Go Dependencies
```
go mod tidy
```
### 3. Run the Server
```
go run main.go
```
### Note: Create a .env file and configure your environment variables (database connection, secret, etc.)

