# Chirpy

Chirpy is a backend social media API built with **Go** and **PostgreSQL**. It features user authentication, a "Chirpy Red" subscription system via webhooks, and advanced filtering/sorting for posts.

## 🚀 Features

* **User Authentication**: JWT-based access tokens and secure refresh tokens stored in the database.
* **Chirp Management**: Create, read, and delete "Chirps" (posts).
* **Filtering**: Fetch chirps by a specific author using `?author_id=` query parameters.
* **Sorting**: Sort chirps in ascending or descending order using the `?sort=` query parameter.
* **Profanity Filter**: Automatic cleaning of "bad words" from chirp content.
* **Polka Integration**: Webhook handler to upgrade users to "Chirpy Red" status.

## 🛠️ Technology Stack

* **Language**: [Go](https://go.dev/)
* **Database**: [PostgreSQL](https://www.postgresql.org/)
* **Tooling**: [SQLC](https://sqlc.dev/) (Type-safe SQL) and [Goose](https://github.com/pressly/goose) (Migrations)

## 🚦 Getting Started

### 1. Prerequisites
Ensure you have Go and PostgreSQL installed on your system.

### 2. Configuration
Create a `.env` file in the root directory and add the following:

PORT=8080
DB_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=your_super_secret_key
PLATFORM=dev
POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e

### 3. Installation & Setup

# Install dependencies
go mod download

# Run migrations (if using Goose)
goose -dir sql/schema postgres "YOUR_DB_URL" up

# Start the server
go run .

## 📖 API Documentation

### Chirps
* `GET /api/chirps`: Get all chirps.
    * Optional: `?author_id={uuid}`
    * Optional: `?sort=asc|desc`
* `POST /api/chirps`: Create a chirp (Requires JWT).
* `DELETE /api/chirps/{id}`: Delete a chirp (Requires JWT, owner only).

### Users
* `POST /api/users`: Create a user.
* `POST /api/login`: Login and receive Access/Refresh tokens.
* `POST /api/polka/webhooks`: Upgrade user to Chirpy Red (Requires Polka API Key).