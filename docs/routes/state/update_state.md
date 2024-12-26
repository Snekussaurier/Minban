# Update State

<h2><span class="color-patch">PATCH</span> <code>/api/v1/state/:state_id</code></h2>

This route is used to update an existing state (or column) for the authenticated user. The state's details, including its name, position, or color, can be modified.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
```plaintext
Content-Type: application/json
```

### Parameters
`card_id` (string, required): The unique identifier of the card to be updated.

### Body (JSON)
The request body must contain the state details. 
Example:

```json
{
  "name": "In Progress",
  "position": 2,
  "color": "00FF00"
}
```

## Responses

### `204` No Content
The request was successful, and the state was updated. No content is returned in the response.

---

### `400` Bad Request
The request was invalid.

---

### `401` Unauthorized
The user is not authenticated, or the minban_token cookie is missing or invalid.

---
### `500` Internal Server Error
An unexpected error occurred while updating the state.