# Create Tag

<h2><span class="color-post">POST</spam> <code>/api/v1/tag</code></h2>

This route is used to create a new tag for the authenticated user. Tags help categorize or label cards.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
```plaintext
Content-Type: application/json
```

### Body (JSON)
The request body must contain the tag details. 
Example:

```json
{
  "name": "Bug",
  "color": "FF0000"
}
```

## Responses

### `201` Created
The tag was successfully created. The response includes the tag's unique ID.

Example response:

```json
{
  "id": 1
}
```

---

### `400` Bad Request
The request was invalid. 

---

### `401` Unauthorized
The user is not authenticated, or the minban_token cookie is missing or invalid.

---

### `500` Internal Server Error
An unexpected error occurred while creating the state.




