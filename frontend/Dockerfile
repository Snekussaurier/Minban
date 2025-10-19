# Multi-stage build for Dioxus frontend
FROM rust:slim-bookworm AS builder

RUN apt-get update && apt-get install -y pkg-config libssl-dev perl make \
  && cargo install dioxus-cli --locked && rustup target add wasm32-unknown-unknown

WORKDIR /app

# Copy Cargo files
COPY Cargo.toml Cargo.lock ./
COPY Dioxus.toml ./

COPY src ./src
COPY assets ./assets
COPY input.css ./
COPY tailwind.config.js ./
COPY package.json ./

# Install Node.js and npm using nvm, then install dependencies
RUN apt-get update && apt-get install -y bash curl \
    && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash \
    && . "$HOME/.nvm/nvm.sh" \
    && nvm install 22 \
    && nvm use 22 \
    && nvm alias default 22 \
    && npm install

# Build the Dioxus application for web
RUN . "$HOME/.nvm/nvm.sh" && nvm use 22 && dx bundle --platform web

# Production stage - use nginx to serve static files
FROM nginx:alpine

# Copy built assets from builder stage (dx bundle creates dist/public folder)
COPY --from=builder /app/dist/public /usr/share/nginx/html

# Copy custom nginx configuration (optional)
# COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 8080
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
