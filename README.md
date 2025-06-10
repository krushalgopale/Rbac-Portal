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
- PostgreSQL database integration
- JWT-based authentication
