# Go JWT Authentication API

A lightweight and secure authentication service built with Golang, Gin, MongoDB, and JWT, featuring role-based access control, Dockerized deployment, and CI/CD integration.

## Features
- User Signup & Login with hashed passwords (bcrypt)
- JWT Access & Refresh Tokens for authentication
- Role-based Access Control (ADMIN, USER)
- MongoDB for persistent storage
- Protected Routes via JWT middleware
- Docker & docker-compose for local and production setups
- GitHub Actions CI/CD with:
    - Go linting & vulnerability scanning
    - Docker image build & security scan
    - Automatic push to Docker Hub

## Project Structure

```bash
.
├── main.go                # Application entry point
├── controllers/           # Business logic for each route
├── database/              # MongoDB connection handling
├── helpers/               # Token generation & role checks
├── middleware/            # JWT authentication middleware
├── models/                # MongoDB schemas
├── routes/                # Route definitions
├── Dockerfile             # Multi-stage build
├── docker-compose.yml     # App + MongoDB + Mongo Express
└── .github/workflows/     # GitHub Actions pipeline
```

## Quick Start

1. Copy the docker-compose.yml file from this repository.
2. Run:

```bash
docker compose up
```

This will automatically pull the latest Docker image from Docker Hub, start the authentication service, MongoDB, and Mongo Express.
- API: http://localhost:9000
- Mongo Express: http://localhost:8081

## API Endpoints

| Method| Endpoint| Auth Required | Role Required | Description |
| :------------ | :--------- | :------ | :------ | ------: |
| POST | /users/signup | No | - | Register a new user |
| POST | /users/login | NO | - | Login and get token |
| GET | /users| Yes | ADMIN | List all users (paginated) |
| GET | /users/:user_id | Yes | ADMIN | Get user by ID |
| GET | /api-1 or /api-2 | Yes | USER/ADMIN | Example protected API |

## Screenshots

Signup

<img alt="signup" src="/images/signup.png">

Login

<img alt="login" src="/images/login.png">

## CI/CD Pipeline (GitHub Actions + Docker Hub)
- Linting & Static Code Analysis: golangci-lint for Go best practices
- Security Scans:
    - gosec → Go code security vulnerabilities
    - Trivy → Docker image OS/library vulnerability scanning
- Automated Docker Builds:
    - Builds tagged image with both :latest and commit SHA
    - Pushes images to Docker Hub
- Artifact Uploads:
    - SARIF vulnerability scan results uploaded to GitHub Security tab
- Build Cache:
    - Go module caching for faster CI runs

    ## Acknowledgements

The core application logic for JWT authentication was implemented by following the excellent tutorial series by Akhil Sharma: [https://youtube.com/playlist?list=PL5dTjWUk_cPY7Q2VTnMbbl8n-H4YDI5wF&si=bZZ_o5_rszljKIJY](https://youtube.com/playlist?list=PL5dTjWUk_cPY7Q2VTnMbbl8n-H4YDI5wF&si=bZZ_o5_rszljKIJY)

Dockerization, CI/CD pipeline (GitHub Actions with security scanning), and other production readiness improvements were implemented independently.
