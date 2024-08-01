
# Task Manager REST API

## Overview
This Task Manager REST API allows users to manage tasks, including creating, updating, retrieving, and deleting tasks. The API is designed to be used with various clients and provides a set of endpoints for task management.

## Base URL
`https://api.example.com`

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
