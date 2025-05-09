# Get Tags

<h2><span class="color-get">GET</spam> <code>/api/v1/tags</code></h2>

This route is used to retrieve all tags associated with the authenticated user.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
_No special headers required._

### Body
_No body is required._

## Responses

### `200` OK

The request was successful, and a list of tags associated with the authenticated user is returned. Each tag includes its details.

Example response:

```json
[
  {
    "id": 1,
    "name": "Bug",
    "color": "FF0000"
  },
  {
    "id": 2,
    "name": "Feature",
    "color": "00FF00"
  },
  {
    "id": 3,
    "name": "Enhancement",
    "color": "0000FF"
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
  "error": "Failed to retrieve states"
}
```