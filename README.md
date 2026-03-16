# MemberList — Member Management System

A full-stack member management system built with **Vue 3 + Vuetify**, **Go + Gin**, and **MongoDB**.

---

## Tech Stack

| Layer     | Technology                        |
|-----------|-----------------------------------|
| Frontend  | Vue 3, Vuetify 3, Vite, Axios     |
| Backend   | Go, Gin Framework            |
| Database  | MongoDB 7                         |
| Container | Docker + Docker Compose           |

---

## Project Structure

```
memberlist-system/
├── docker-compose.yml
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│   ├── .env.example
│   ├── config/
│   │   └── database.go          # MongoDB connection
│   ├── models/
│   │   └── member.go            # Data structures
│   ├── controllers/
│   │   └── member_controller.go # HTTP handlers
│   ├── services/
│   │   └── member_service.go    # Business logic
│   └── routes/
│       └── routes.go            # Route registration
└── frontend/
    ├── index.html
    ├── vite.config.js
    ├── package.json
    ├── Dockerfile
    ├── nginx.conf
    └── src/
        ├── main.js              # App entry + Vuetify setup
        ├── App.vue              # Root component
        ├── components/
        │   ├── DatePicker.vue  # date picker
        ├── views/
        │   ├── MemberView.vue  # Data table with pagination
        ├── modals/
        │   ├── memberModals.vue  # Add member modal
        └── api/
            └── member.js # Axios API calls
```

---

## Quick Start

### Option 1: Docker Compose (Recommended)

```bash
# Start all services (MongoDB, Backend, Frontend)
docker-compose up --build

# Frontend:  http://localhost:8080
# Backend:   http://localhost:3000
# MongoDB:   localhost:27017
```

---

## API Reference

### GET /api/members

<!-- Retrieve a paginated list of members.

**Query Parameters:**

| Parameter | Type   | Default | Description              |
|-----------|--------|---------|--------------------------|
| `page`    | int    | 1       | Page number (1-indexed)  |
| `limit`   | int    | 10      | Items per page           |

**Example Request:**
```
GET /api/members?page=1&limit=10
``` -->

**Example Response:**
```json
{
  "data": {
    "data": [
      {
        "id": "65a1b2c3d4e5f6789012345",
        "name": "Jane",
        "surname": "Doe",
        "address": "123 Main St, Bangkok, Thailand",
        "date_of_birth": "1990-05-15",
        "age": 34,
        "created_at": "2024-01-15T10:30:00Z"
      }
    ],
    "total": 42,
    "page": 1,
    "limit": 10,
    "total_pages": 5
  }
}
```

---

### POST /api/members

Create a new member.

**Request Body:**
```json
{
  "name": "Jane",
  "surname": "Doe",
  "address": "456 Sukhumvit Rd, Bangkok, Thailand",
  "date_of_birth": "1995-08-22"
}
```

**Example Response:**
```json
{
  "data": {
    "id": "65a1b2c3d4e5f6789012346",
    "name": "Jane",
    "surname": "Doe",
    "address": "456 Sukhumvit Rd, Bangkok, Thailand",
    "date_of_birth": "1995-08-22",
    "age": 28,
    "created_at": "2024-01-15T11:00:00Z"
  }
}
```

---

## Features

- **Member Table** — Vuetify data table with Full Name (with avatar), Address, Date of Birth, and calculated Age
- **Add Member Modal** — Validated form with live age preview
- **Pagination** — Server-side pagination with configurable page sizes (5, 10, 20, 50)

---

## Development Notes

- Age is calculated dynamically on every API response — it's never stored in the DB
- Members are sorted by `created_at` ascending, so new members appear at the end
- After adding a member, the frontend automatically navigates to the last page
- The `date_of_birth` field is stored as a string `"YYYY-MM-DD"` in MongoDB for simplicity and locale-independence
