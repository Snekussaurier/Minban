# Check Authentication Status

<h2><span class="color-get">GET</spam> <code>/api/v1/check-auth</code></h2>

This route is used to verify if a user is currently authenticated by checking the validity of the `minban_token` cookie. It ensures the user is logged in and their session is still active.

## Request

### Authentication
🔒 Authentication required

### Headers
_No special headers required._

### Body
_No body is required._

## Responses

### `200` OK

The request was successful, and the `minban_token` is in the cookie store is not empty and valid.

---

### `401` Unauthorized

The authentication cookie is not present or invalid.