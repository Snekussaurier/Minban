# Create Card

<h2><span class="color-post">POST</spam> <code>/api/v1/card</code></h2>

This route is used to create a new card for the authenticated user. A card must belong to an existing state and can optionally have associated tags.

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
```plaintext
Content-Type: application/json
```

### Body (JSON)
The request body must contain the card details. Example:

```json
{
  "title": "Implement Authentication",
  "description": "Add login functionality to the application",
  "state_id": 1,
  "position": 1,
  "tags": [
    {
      "id": 2
    }
  ]
}
```

## Responses

### `201` Created
The card was successfully created. The response includes the card's unique ID.

Example response:

```json
{
  "id": "b7e9a101-6c58-4f21-8228-c1a2bb3fcf38"
}
```

---

### `400` Bad Request
The request was invalid. 

```json
{
  "error": "State with ID: 1 not found"
}
```

!!! failure "Reasons for an invalid request"
    - Missing required fields (title, state_id, position).
    - Invalid state_id or tag ID (state or tag does not exist or does not belong to the user).

---

### `401` Unauthorized
The user is not authenticated, or the minban_token cookie is missing or invalid.

---

### `500` Internal Server Error
An unexpected error occurred while creating the card.




