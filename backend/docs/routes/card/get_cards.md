# Get Cards

<h2><span class="color-get">GET</spam> <code>/api/v1/cards</code></h2>

This route is used to retrieve all cards associated with the authenticated user. Each card will include its associated tags.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
_No special headers required._

### Body
_No body is required._

## Responses

### `200` OK

The request was successful, and a list of cards associated with the authenticated user is returned. Each card includes its associated tags.

Example response:

```json
[
  {
    "id": 1,
    "title": "Fix Bug #123",
    "description": "Resolve the issue causing app crashes.",
    "state_id": 2,
    "position": 1,
    "tags": [
      {
        "id": 1,
        "name": "Bug",
        "color": "FF0000"
      }
    ]
  },
  {
    "id": 2,
    "title": "Add New Feature",
    "description": "Implement the new user dashboard.",
    "state_id": 1,
    "position": 2,
    "tags": [
      {
        "id": 2,
        "name": "Feature",
        "color": "00FF00"
      }
    ]
  }
]
```

---

### `401` Unauthorized

The user is not authenticated, or the `minban_token` cookie is missing or invalid.

Example response:

```json
{
  "error": "Unauthorized"
}
```

---

### `500` Internal Server Error

An unexpected error occurred while querying the database.

Example response:

``` json
{
  "error": "Failed to retrieve cards"
}
```