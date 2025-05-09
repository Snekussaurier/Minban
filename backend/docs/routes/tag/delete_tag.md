# Delete Tag

<h2><span class="color-delete">DELETE</span> <code>/api/v1/tag/:tag_id</code></h2>

This route is used to delete an existing tag for the authenticated user. Once deleted, the tag is permanently removed from the database.

---

## Request

### Authentication
🔒 Authentication is required. A valid `minban_token` cookie must be present to authorize the request.

### Headers
_No special headers required._

### Parameters
`tag_id` (integer, required): The unique identifier of the tag to be deleted.

## Responses

### `204` No Content
The request was successful, and the tag was deleted. No content is returned in the response.

---

### `400` Bad Request
The request was invalid.

---

### `401` Unauthorized
The user is not authenticated, or the `minban_token` cookie is missing or invalid.

---

### `500` Internal Server Error
An unexpected error occurred while deleting the card.
