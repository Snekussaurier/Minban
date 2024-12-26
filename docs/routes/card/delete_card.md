# Delete Card

<h2><span class="color-delete">DELETE</span> <code>/api/v1/card/:card_id</code></h2>

This route is used to delete an existing card for the authenticated user. Once deleted, the card is permanently removed from the database.

---

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
_No special headers required._

### Parameters
`card_id` (string, required): The unique identifier of the card to be deleted.

## Responses

### `204` No Content
The request was successful, and the card was deleted. No content is returned in the response.

---

### `400` Bad Request
The request was invalid.

---

### `401` Unauthorized
The user is not authenticated, or the `minban_token` cookie is missing or invalid.

---

### `500` Internal Server Error
An unexpected error occurred while deleting the card.
