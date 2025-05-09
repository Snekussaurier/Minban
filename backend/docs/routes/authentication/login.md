# Login 

<h2><span class="color-post">POST</spam> <code>/api/v1/login</code></h2>

This route is used to authenticate a user by validating their credentials. On success, it sets an `minban_token` cookie containing the JWT token for further authenticated requests.

## Request

### Authentication
🔓 None

### Headers
``` plaintext
Content-Type: application/json
```

### Body (JSON)
``` json
{
  "username": "<username>",
  "password": "<passowrd>"
}
```

## Responses

### `204` No content

The request was successful, and the user is authenticated. A secure `minban_token` cookie is set, containing the JWT token for future requests.

---

### `400` Bad request

The request was invalid, usually because the provided username or password is incorrect. The response body will look like:

```json
{
  "error": "Invalid username or password"
}
```

---

### `500` Internal server error

An unexpected error occurred on the server, typically when generating the JWT token fails. The response body will look like:

```json
{
  "error": "Failed to generate token"
}
```