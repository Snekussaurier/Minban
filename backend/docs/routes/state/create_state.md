# Create State

<h2><span class="color-post">POST</spam> <code>/api/v1/state</code></h2>

This route is used to create a new state or column for the authenticated user. States represent workflow stages, and cards are assigned to these states.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
```plaintext
Content-Type: application/json
```

### Body (JSON)
The request body must contain the state details.Example:

```json
{
  "name": "In Progress",
  "position": 2,
  "color": "00FF00"
}
```

## Responses

### `201` Created
The state was successfully created. The response includes the state's unique ID.

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




