# Task Manager REST API
### Overview

The Task Management REST API is a backend application built using the Go programming language and the Gin framework. This API enables users to manage tasks through basic CRUD (Create, Read, Update, Delete) operations. The API is designed to be simple, efficient, and scalable, providing a foundational backend service for task management applications.

### Features

- **Create a Task**: Add new tasks with details such as title, description, due date, and status.
    
- **Read Tasks**: Retrieve a list of all tasks or specific task details by ID.
    
- **Update a Task**: Modify existing task details.
    
- **Delete a Task**: Remove tasks from the system.
    

### Technologies Used

- **Go**: A statically typed, compiled programming language known for its simplicity and performance.
    
- **Gin Framework**: A lightweight and fast HTTP web framework for Go, ideal for building RESTful APIs.
    
- **MongoDB**: A NoSQL database known for its flexibility and scalability, used for storing task data.
    

### MongoDB Integration

- **MongoDB Driver**: Utilizes the official MongoDB Go driver for database interactions.
    
- **Task Collection**: Stores tasks in a MongoDB collection, ensuring efficient and scalable data management.
    
- **Database Operations**: Implements CRUD operations using MongoDB to handle task data, providing robust data storage and retrieval capabilities.
# 📁 Folder: Tasks 


## End-point: Create a Task
### Add Task

This endpoint allows the client to add a new task.

#### Request Body

- `id` (string, required): The unique identifier for the task.
    
- `title` (string, required): The title of the task.
    
- `description` (string, required): The description of the task.
    
- `due_date` (string, required): The due date for the task.
    
- `status` (string, required): The status of the task.
    

#### Response

Upon successful creation of the task, the endpoint returns a status code of 201 and a JSON response with a message indicating the success of the operation.

Example:

``` json
{
    "message": "Task added successfully"
}

 ```
### Method: POST
>```
>localhost:8080/tasks
>```
### Body (**raw**)

```json
{
            "title": "Task 10",
            "description": "Lost Task",
            "due_date": "2024-07-31T10:04:14+03:00",
            "status": "Complete"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get a Task
### GET /tasks/:id

This endpoint retrieves a specific task identified by the provided ID.

#### Request

No request body is required for this endpoint.

- `id` (string, required): The ID of the task to be retrieved.
    

#### Response

The response is in JSON format with the following schema:

``` json
{
    "id": {"type": "string"},
    "title": {"type": "string"},
    "description": {"type": "string"},
    "due_date": {"type": "string"},
    "status": {"type": "string"}
}

 ```
### Method: GET
>```
>localhost:8080/tasks/:id
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get all Tasks
### GET /tasks

This endpoint retrieves a list of tasks.

#### Request

No request body is required for this endpoint.

#### Response

The response will be in JSON format with a 200 status code. The response body will contain an array of tasks, where each task object includes the following properties:

- `id` (string): The unique identifier for the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date for the task.
    
- `status` (string): The status of the task.
    

Example response body:

``` json
{
    "tasks": [
        {
            "id": "",
            "title": "",
            "description": "",
            "due_date": "",
            "status": ""
        }
    ]
}

 ```
### Method: GET
>```
>localhost:8080/tasks
>```
### Body (**raw**)

```json

```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update a Task
### Update Task

This endpoint is used to update a specific task identified by its ID.

#### Request

- Method: PUT
    
- URL: `localhost:8080/tasks/:id`
    
- Headers:
    
    - Content-Type: application/json
        
- Body:
    
    - `id` (string)
        
    - `title` (string)
        
    - `description` (string)
        
    - `due_date` (string)
        
    - `status` (string)
        

#### Response

The response is in JSON format with the following schema:

``` json
{
    "message": "task updated"
}

 ```
### Method: PUT
>```
>localhost:8080/tasks/:id
>```
### Body (**raw**)

```json
{
    "title": "ermi codefordes",
    "description": "Lost cxause",
    "due_date": "2024-07-31T10:04:14.9521823+03:00",
    "status": "Pending"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete a Task
### Delete Task

The `DELETE` request is used to delete a specific task identified by the `id` parameter from the server.

### Response

The response returned from the server has a status code of 200 and a content type of `application/json`. The response body is a JSON object with a `message` field.

``` json
{
    "message": "task deleted"
}

 ```
### Method: DELETE
>```
>localhost:8080/tasks/:id
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Folder: User Auth 


## End-point: User Register
## Register User

This endpoint allows the client to register a new user.

### Request Body

- `email` (string, required): The email address of the user.
    
- `password` (string, required): The password for the user account.
    
- `isadmin` (boolean, required): Indicates whether the user has admin privileges.
    

### Response

The response for this request is a JSON object conforming to the following schema:

``` json
{
    "message": "User registered successfully"
}

 ```

The `userId` property contains the unique identifier for the newly registered user, and the `message` property provides a success message.
### Method: POST
>```
>localhost:8080/register
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer |


### Body (**raw**)

```json
{
    "email": "dagmmesfin92@gmail.com",
    "password": "ka2kasd4",
    "isadmin": false
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: User Login
# Login Endpoint

This API endpoint is used to authenticate a user by providing their email and password.

## Request Body

- `email` (string): The email address of the user.
    
- `password` (string): The password of the user.
    

## Response

Upon a successful request, the server responds with a status code of 200 and a JSON object containing the following fields:

- `message` (string): A message indicating the result of the login attempt.
    
- `token` (string): A token for the authenticated user session.
    

``` json
{
    "message": "User logged in successfully",
    "token": "string"
}

 ```
### Method: POST
>```
>localhost:8080/login
>```
### Body (**raw**)

```json
{
    "email": "dagmmesfin91@gmail.com",
    "password": "ka2kasd4"
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
