# api documentaions for /v1/users endpoint

This document includes all the api documentations for the users endpoint. 

## endpoints 

- GET /users
- POST /users
- GET /users/:userID
- DELETE /users/:userID
- PATCH /users/:userID/forgotpassword

### GET /users

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "users": [
        {
            "id": 1,
            "name": "java",
            "active": false,
            "created_at": "2019-11-03T03:26:12Z",
            "updated_at": "2019-11-03T04:31:00Z"
        }
    ]
}
```

#### Internal Server Error
1. StatusCode: `500 Internal Server Error` 
2. Response: 
```json
{
    "error": ""
}
```

### GET /users/:userID
- Parameters: 
  - userID:
    - Integer
    - Path

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "id": 1,
    "name": "java",
    "active": false,
    "created_at": "2019-11-03T03:26:12Z",
    "updated_at": "2019-11-03T04:31:00Z"
}
```

#### Bad Request
1. StatusCode: `400 Bad Request` 
2. Response: 
```json
{
    "error": "Invalid id Passed"
}
```

#### Not Found
1. StatusCode: `404 Not Found` 
2. Response: 
```json
{
    "error": "No user found with the ID: {userID}"
}
```

#### Internal Server Error
1. StatusCode: `500 Internal Server Error` 
2. Response: 
```json
{
    "error": ""
}
```

### POST /users

#### CREATED
1. StatusCode: `201 Created` 
2. Response: 
```json
{
    "message": "success"
}
```

#### Bad Request
1. StatusCode: `400 Bad Request` 
2. Response: 
```json
{
    "error": "json: cannot unmarshal number into Go struct field user.name of type string"
}
```

#### Internal Server Error
1. StatusCode: `500 Internal Server Error` 
2. Response: 
```json
{
    "error": ""
}
```

### PATCH /users/:userID/forgotpassword
- Parameters: 
  - userID:
    - Integer
    - Path

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "id": 1,
    "name": "java",
    "active": false,
    "created_at": "2019-11-03T03:26:12Z",
    "updated_at": "2019-11-03T04:31:00Z"
}
```

#### Bad Request
1. StatusCode: `400 Bad Request` 
2. Response: 
```json
{
    "error": "Invalid id Passed"
}
```

#### Not Found
1. StatusCode: `404 Not Found` 
2. Response: 
```json
{
    "error": "No user found with the ID: {userID}"
}
```
```json
{
    "error": "Cannot use the same password for new password!"
}
```
```json
{
    "error": "Wrong Password"
}
```

#### Internal Server Error
1. StatusCode: `500 Internal Server Error` 
2. Response: 
```json
{
    "error": ""
}
```

### DELETE /users/:userID
- Parameters: 
  - userID:
    - Integer
    - Path

#### ACCEPTED
1. StatusCode: `202 Accepted` 
2. Response: 
```json
{
    "message": "deleted"
}
```

#### Bad Request
1. StatusCode: `400 Bad Request` 
2. Response: 
```json
{
    "error": "Invalid id Passed"
}
```

#### Not Found
1. StatusCode: `404 Not Found` 
2. Response: 
```json
{
    "error": "No user found with the ID: {userID}"
}
```

#### Internal Server Error
1. StatusCode: `500 Internal Server Error` 
2. Response: 
```json
{
    "error": ""
}
```