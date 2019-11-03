# api documentaions for /v1/tags endpoint

This document includes all the api documentations for the tags endpoint. 

## ednpoints 

- GET /tags
- POST /tags
- GET /tags/:tagID
- DELETE /tags/:tagID
- PUT /tags/:tagID

### GET /tags

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "tags": [
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

### GET /tags/:tagID
- Parameters: 
  - tagID:
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
    "error": "No tag found with the ID: {tagID}"
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

### POST /tags

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
    "error": "json: cannot unmarshal number into Go struct field Tag.name of type string"
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

### UPDATE /tags/:tagID
- Parameters: 
  - tagID:
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
    "error": "No tag found with the ID: {tagID}"
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

### DELETE /tags/:tagID
- Parameters: 
  - tagID:
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
    "error": "No tag found with the ID: {tagID}"
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

