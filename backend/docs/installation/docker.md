# Docker Installation

## Using Docker Compose

To set up and run the MinBan backend using **Docker Compose**, create a `docker-compose.yml` file in your project directory with the following configuration:

```yaml
services:
  app:
    image: ghcr.io/snekussaurier/minban-backend:latest
    container_name: minban-backend
    ports:
      - "9916:9916"
    volumes:
      - ./data:/app/data  # Ensure that your local 'data' directory exists
    environment:
      - DATABASE_PATH=/app/data/miniban.db  # Path to the SQLite database
      - USER_NAME=default  # Default username
      - USER_PASSWORD=123  # Default user password
      - JWT_SECRET_KEY=<secure_secret_key>  # Secret key for JWT authentication
```

Save this configuration to a file called docker-compose.yml.

Run the application with the following command:

``` bash
docker-compose up
```

This will start the MinBan backend application on port 9916, and it will persist data in the ./data directory on your local machine.

## Using Docker Run

Alternatively, you can run the MinBan backend using the docker run command:

``` bash
docker run -p 9916:9916 -v ./data:/app/data \
  -e DATABASE_PATH=/app/data/miniban.db \
  -e USER_NAME=snekussaurier \
  -e USER_PASSWORD=123 \
  -e JWT_SECRET_KEY=<secure_secret_key> \
  ghcr.io/snekussaurier/minban-backend:latest
```