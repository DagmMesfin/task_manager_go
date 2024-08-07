# Task Manager REST API Documentation

## Overview

The Task Management REST API is a backend application built using the Go programming language and the Gin framework. This API enables users to manage tasks through basic CRUD (Create, Read, Update, Delete) operations. The API is designed to be simple, efficient, and scalable, providing a foundational backend service for task management applications.

## Features

- **User Registration and Authentication**: Secure user registration and login using JWT tokens.
- **Task Management**: Create, read, update, and delete tasks.
- **User Role Management**: Admin and non-admin roles with different access levels.

## Technologies Used

- **Go**: A statically typed, compiled programming language known for its simplicity and performance.
- **Gin Framework**: A lightweight and fast HTTP web framework for Go, ideal for building RESTful APIs.
- **MongoDB**: A NoSQL database known for its flexibility and scalability, used for storing task and user data.
- **JWT (JSON Web Tokens)**: Used for securing API endpoints by providing token-based authentication.

## MongoDB Integration

- **MongoDB Driver**: Utilizes the official MongoDB Go driver for database interactions.
- **Task Collection**: Stores tasks in a MongoDB collection, ensuring efficient and scalable data management.
- **User Collection**: Stores user data, including encrypted passwords and role information.

## Endpoints

### Task Endpoints

#### Create a Task

**Endpoint:** `POST /tasks`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Request Body:**
```json
{
  "title": "Task 10",
  "description": "Lost Task",
  "due_date": "2024-07-31T10:04:14+03:00",
  "status": "Complete"
}
```

**Description:** This endpoint allows the client to add a new task.

**Response:**
```json
{
  "message": "Task added successfully"
}
```

---

#### Get a Task

**Endpoint:** `GET /tasks/:id`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Description:** This endpoint retrieves a specific task identified by the provided ID.

**Response:**
```json
{
  "id": "66b337085f620f085a4985e5",
  "title": "Task Title",
  "description": "Task Description",
  "due_date": "2024-07-31T10:04:14+03:00",
  "status": "Pending"
}
```

---

#### Get all Tasks

**Endpoint:** `GET /tasks`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Description:** This endpoint retrieves a list of tasks.

**Response:**
```json
{
  "tasks": [
    {
      "id": "66b337085f620f085a4985e5",
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-07-31T10:04:14+03:00",
      "status": "Pending"
    }
  ]
}
```

---

#### Update a Task

**Endpoint:** `PUT /tasks/:id`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Request Body:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "due_date": "2024-07-31T10:04:14+03:00",
  "status": "Pending"
}
```

**Description:** This endpoint is used to update a specific task identified by its ID.

**Response:**
```json
{
  "message": "Task updated"
}
```

---

#### Delete a Task

**Endpoint:** `DELETE /tasks/:id`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Description:** The `DELETE` request is used to delete a specific task identified by the `id` parameter from the server.

**Response:**
```json
{
  "message": "Task deleted"
}
```

---

### User Authentication Endpoints

#### Register User

**Endpoint:** `POST /register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "userpassword",
  "isadmin": true
}
```

**Description:** This endpoint allows the client to register a new user.

**Response:**
```json
{
  "message": "User registered successfully"
}
```

---

#### User Login

**Endpoint:** `POST /login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "userpassword"
}
```

**Description:** This API endpoint is used to authenticate a user by providing their email and password.

**Response:**
```json
{
  "message": "User logged in successfully",
  "token": "string"
}
```

---

#### Delete User

**Endpoint:** `DELETE /users/:id`

**Request Headers:**
```
Authorization: Bearer <jwt-token>
```

**Description:** This endpoint deletes a user identified by the `id` parameter.

**Response:**
```json
{
  "message": "User deleted"
}
```

---

## Authentication

The API uses JWT (JSON Web Token) for authentication and authorization.

- **Bearer Token:** Include the JWT token in the `Authorization` header of the requests to endpoints that require authentication.

**Example:**
```
Authorization: Bearer <jwt-token>
```

## Environment Variables

- **JWT_SECRET:** Secret key for signing JWT tokens.
- **MONGO_URI:** Connection string for the MongoDB database.

## Running the API

1. **Install Dependencies:** Ensure you have Go and MongoDB installed.
2. **Clone the Repository:**
   ```bash
   git clone https://github.com/DagmMesfin/task_manager_go.git
   ```
3. **Navigate to the Project Directory:**
   ```bash
   cd task_manager_go
   ```
4. **Set Environment Variables:**
   ```bash
   export JWT_SECRET=<your_jwt_secret>
   export MONGO_URI=<your_mongo_uri>
   ```
5. **Run the Application:**
   ```bash
   go run main.go
   ```