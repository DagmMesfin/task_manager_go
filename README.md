
# Task Manager REST API

## Overview
This Task Manager REST API allows users to manage tasks, including creating, updating, retrieving, and deleting tasks. The API is designed to be used with various clients and provides a set of endpoints for task management.

# Task Management REST API

## Overview

The Task Management REST API is a backend application built using the Go programming language and the Gin framework. This API enables users to manage tasks through basic CRUD (Create, Read, Update, Delete) operations. The API is designed to be simple, efficient, and scalable, providing a foundational backend service for task management applications.

## Features

- **Create a Task**: Add new tasks with details such as title, description, due date, and status.
- **Read Tasks**: Retrieve a list of all tasks or specific task details by ID.
- **Update a Task**: Modify existing task details.
- **Delete a Task**: Remove tasks from the system.

## Technologies Used

- **Go**: A statically typed, compiled programming language known for its simplicity and performance.
- **Gin Framework**: A lightweight and fast HTTP web framework for Go, ideal for building RESTful APIs.
- **MongoDB**: A NoSQL database known for its flexibility and scalability, used for storing task data.

## MongoDB Integration

- **MongoDB Driver**: Utilizes the official MongoDB Go driver for database interactions.
- **Task Collection**: Stores tasks in a MongoDB collection, ensuring efficient and scalable data management.
- **Database Operations**: Implements CRUD operations using MongoDB to handle task data, providing robust data storage and retrieval capabilities.

## Base URL
`https://localhost:8080/`

## Endpoints

### Create a Task
- **URL:** `/tasks`
- **Method:** `POST`
- **Request Body:** JSON
- **Response:** JSON

### Retrieve All Tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Response:** JSON

### Retrieve a Single Task
- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Response:** JSON

### Update a Task
- **URL:** `/tasks/{id}`
- **Method:** `PUT`
- **Request Body:** JSON
- **Response:** JSON

### Delete a Task
- **URL:** `/tasks/{id}`
- **Method:** `DELETE`
- **Response:** JSON

## Error Handling
Each endpoint provides appropriate HTTP status codes and messages for errors.

## Sample Request
```json
{
  "title": "New Task",
  "description": "Task details",
  "dueDate": "2024-08-01"
}
```

## Sample Response
```json
{
  "id": "1",
  "title": "New Task",
  "description": "Task details",
  "dueDate": "2024-08-01",
  "status": "pending"
}
```

For detailed information, visit the [API documentation](https://documenter.getpostman.com/view/37336034/2sA3kdAcpR).
