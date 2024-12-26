# Update Card

<h2><span class="color-patch">PATCH</spam> <code>/api/v1/card/:card_id</code></h2>

This route is used to update an existing card for the authenticated user. The card's details, including its title, description, state, position, or tags, can be modified.

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
The request body must contain the card details. Example:

```json
{
  "title": "Update Authentication",
  "description": "Refactor login functionality for better security",
  "state_id": 2,
  "position": 3,
  "tags": [
    {
      "id": 1
    },
    {
      "id": 3
    }
  ]
}
```

## Responses

### `204` No Content
The request was successful, and the card was updated. No content is returned in the response.

---

### `400` Bad Request
The request was invalid.

---

### `401` Unauthorized
The user is not authenticated, or the minban_token cookie is missing or invalid.

---
### `500` Internal Server Error
An unexpected error occurred while updating the card.