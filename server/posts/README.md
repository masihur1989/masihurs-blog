# api documentaions for /v1/posts endpoint

This document includes all the api documentations for the posts endpoint. 

## endpoints 

- GET /posts
- POST /posts
- GET /posts/:postID
- DELETE /posts/:postID
- PUT /posts/:postID
- PATCH /posts/:postID/postview

### GET /posts

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "posts": [
        {
            "id": 1,
            "title": "Java Class",
            "body": "Java Class Syntext",
            "user_id": 1,
            "category_id": 4,
            "post_view": 0,
            "active": true,
            "created_at": "2019-11-04T01:18:29Z",
            "updated_at": "2019-11-04T01:18:29Z"
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

### GET /posts/:postID
- Parameters: 
  - postID:
    - Integer
    - Path

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "id": 1,
    "title": "Java Class",
    "body": "Java Class Syntext",
    "user_id": 1,
    "category_id": 4,
    "post_view": 0,
    "active": true,
    "created_at": "2019-11-04T01:18:29Z",
    "updated_at": "2019-11-04T01:18:29Z"
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
    "error": "No post found with the ID: {postID}"
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

### POST /posts
- Request Body:
```json
{
    "title": "Java Class",
    "body": "Java Class Syntext",
    "user_id": 1,
    "category_id": 4,
    "post_view": 0,
    "active": false
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

### PUT /posts/:postID
- Parameters: 
  - postID:
    - Integer
    - Path
- Request Body:
```json
{
    "title": "Java Class",
    "body": "Java Class Syntext",
    "user_id": 1,
    "category_id": 4,
    "post_view": 0,
    "active": false
}
```

#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "id": 1,
    "title": "Java Class",
    "body": "Java Class Syntext",
    "user_id": 1,
    "category_id": 4,
    "post_view": 0,
    "active": false,
    "created_at": "2019-11-04T01:18:29Z",
    "updated_at": "2019-11-04T02:08:09Z"
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
    "error": "No post found with the ID: {postID}"
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

### DELETE /posts/:postID
- Parameters: 
  - postID:
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
    "error": "No post found with the ID: {postID}"
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


### PATCH /posts/:postID/postview
- Parameters: 
  - postID:
    - Integer
    - Path
- Request Body:
```json
{
    "post_view": 0 // current post_view count
}
```


#### OK
1. StatusCode: `200 OK` 
2. Response: 
```json
{
    "message": "updated"
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
    "error": "No post found with the ID: {postID}"
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

