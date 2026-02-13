# Chirpy-Go API Documentation

This is a backend server built in Go that mimics basic Twitter functionality. It provides a RESTful API for user management, authentication, and posting "chirps."

## Base URL
All requests should be made to:
`http://localhost:8080` (or your configured port)

---

## Content Type
For all `POST` and `PUT` requests, include the following header:
`Content-Type: application/json`

---

## Authentication
Certain endpoints require a Bearer Token. Include the token in the request header:
`Authorization: Bearer <your_access_token>`

---

## API Endpoints

### 1. Health Check
Checks if the server is running.
* **URL:** `/api/healthz`
* **Method:** `GET`
* **Response:** `200 OK` (Plain text: `OK`)

---

### 2. Chirps

#### Create a Chirp
Posts a new chirp. The body must be under 140 characters.
* **URL:** `/api/chirps`
* **Method:** `POST`
* **Auth Required:** Yes
* **Request Body:**
    ```json
    {
      "body": "Hello world!"
    }
    ```
* **Success Response (201 Created):**
    ```json
    {
      "id": "uuid-string",
      "created_at": "timestamp",
      "updated_at": "timestamp",
      "body": "Hello world!",
      "user_id": "uuid-string"
    }
    ```

#### Get All Chirps
* **URL:** `/api/chirps`
* **Method:** `GET`
* **Success Response (200 OK):**
    ```json
    [
      {
        "id": "uuid-1",
        "body": "First chirp",
        "user_id": "user-uuid"
      },
      {
        "id": "uuid-2",
        "body": "Second chirp",
        "user_id": "user-uuid"
      }
    ]
    ```

#### Get Single Chirp
* **URL:** `/api/chirps/{chirpID}`
* **Method:** `GET`
* **Success Response (200 OK):** JSON object of the chirp.

---

### 3. Users

#### Create User
* **URL:** `/api/users`
* **Method:** `POST`
* **Request Body:**
    ```json
    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```
* **Success Response (201 Created):**
    ```json
    {
      "id": "uuid-string",
      "created_at": "timestamp",
      "updated_at": "timestamp",
      "email": "user@example.com",
      "is_chirpy_red": false
    }
    ```

#### Update User
Update email or password for the authenticated user.
* **URL:** `/api/users`
* **Method:** `PUT`
* **Auth Required:** Yes
* **Request Body:**
    ```json
    {
      "email": "newemail@example.com",
      "password": "newpassword"
    }
    ```

---

### 4. Authentication

#### Login
* **URL:** `/api/login`
* **Method:** `POST`
* **Request Body:**
    ```json
    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```
* **Success Response (200 OK):**
    ```json
    {
      "id": "uuid-string",
      "email": "user@example.com",
      "token": "JWT_ACCESS_TOKEN",
      "refresh_token": "REFRESH_TOKEN"
    }
    ```

#### Refresh Token
Generate a new access token using a refresh token.
* **URL:** `/api/refresh`
* **Method:** `POST`
* **Auth Required:** Yes (Use Refresh Token in Header)
* **Success Response (200 OK):**
    ```json
    {
      "token": "NEW_JWT_ACCESS_TOKEN"
    }
    ```

#### Revoke Token
Revokes a refresh token (Logout).
* **URL:** `/api/revoke`
* **Method:** `POST`
* **Auth Required:** Yes (Use Refresh Token in Header)
* **Success Response:** `204 No Content`

---

## Error Handling
If a request fails (e.g., invalid JSON, unauthorized, or validation error), the server returns a JSON error object:

**Response Format:**
```json
{
  "error": "Description of what went wrong"
}
```

## Common Status Codes:
* **400 Bad Request:** `Malformed JSON or validation failure`
* **401 Unauthorized** `Missing or invalid Bearer token`
* **404 Not Found** `Resource does not exist`
* **500 Internal Server Error** `Server-side issue`