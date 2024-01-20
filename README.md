# Telemedicine App API

This is the API for the Telemedicine App. It is built using the Go programming language and the Fiber framework.

## Getting Started

To get started, follow the instructions below.

1. Clone the repository

```bash
git clone https://github.com/aaqyaar/telemedicine-api
```

2. Copy the `.env.example` file to `.env` and fill in the required values

```bash
cp .env.example .env
```

3. Install the dependencies using `go get .`
4. Run `go run .` to start the server

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/) (optional)

## Built With

- [x] Go & Fiber
- [x] PostgreSQL
- [x] RabbitMQ
- [x] Redis (for caching)

## Authors

- **Abdi Zamed Mohamed** (Team Lead)
- **Abdullahi Abdi Ahmed**
- **Ancaam Awil**
- **Shukri Dahir**

## API Documentation

The API documentation is available at [Postman](Postman)

### Basic Endpoints

| Method | Endpoint                     | Description                 |
| ------ | ---------------------------- | --------------------------- |
| `POST` | /api/v1/auth/register        | Register a new user         |
| `POST` | /api/v1/auth/login           | Login a user                |
| `POST` | /api/v1/auth/logout          | Logout a user               |
| `POST` | /api/v1/auth/refresh         | Refresh a user's token      |
| `POST` | /api/v1/auth/forgot-password | Send a password reset email |
| `POST` | /api/v1/auth/reset-password  | Reset a user's password     |

## Contributing

Contributions are welcome. Feel free to open a pull request or an issue.
