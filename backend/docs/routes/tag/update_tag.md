# Update Tag

<h2><span class="color-patch">PATCH</span> <code>/api/v1/tag/:tag_id</code></h2>

This route is used to update an existing tag for the authenticated user. The tag's details, including its name or color, can be modified.

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
The request body must contain the tag details. 
Example:

```json
{
  "name": "Feature",
  "color": "00FF00"
}
```

## Responses

### `204` No Content
The request was successful, and the tag was updated. No content is returned in the response.

---

### `400` Bad Request
The request was invalid.

---

### `401` Unauthorized
The user is not authenticated, or the minban_token cookie is missing or invalid.

---
### `500` Internal Server Error
An unexpected error occurred while updating the state.