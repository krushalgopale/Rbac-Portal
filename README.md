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

