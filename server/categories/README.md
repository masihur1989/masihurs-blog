# api documentaions for /v1/categories endpoint

This document includes all the api documentations for the categories endpoint. 

## endpoints 

- GET /categories
- POST /categories
- GET /categories/:categoryID
- DELETE /categories/:categoryID
- PUT /categories/:categoryID

### GET /categories

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "categories": [
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

### GET /categories/:categoryID
- Parameters: 
  - categoryID:
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
    "error": "No category found with the ID: {categoryID}"
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

### POST /categories
- Request Body:
```json
{
    "name": "java",
    "active": true
}
```

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
    "error": "json: cannot unmarshal number into Go struct field Category.name of type string"
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

### PUT /categories/:categoryID
- Parameters: 
  - categoryID:
    - Integer
    - Path
- Request Body:
```json
{
    "name": "java",
    "active": true
}
```

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
    "error": "No category found with the ID: {categoryID}"
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

### DELETE /categories/:categoryID
- Parameters: 
  - categoryID:
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
    "error": "No category found with the ID: {categoryID}"
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

