# Logout

<h2><span class="color-post">POST</spam> <code>/api/v1/logout</code></h2>

This route is used to log out a user by invalidating their `minban_token` cookie. Once the cookie is cleared, the user will no longer be authenticated for future requests.

## Request

### Authentication
🔒 No authentication

### Headers
_No special headers required._

### Body
_No body is required._

## Responses

### `204` No Content

The request was successful, and the `minban_token` cookie has been cleared. The user is now logged out.