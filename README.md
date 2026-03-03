# goleam

A minimal Go web application scaffolded for Docker-based development.

## Structure

```
cmd/server/main.go     # application entrypoint
internal/handler       # app-specific packages
build/Dockerfile       # container build instructions
compose.yaml           # Docker Compose configuration
.go.mod/.go.sum        # module definitions
```

## Usage

```bash
# build & run production image via compose
docker compose up --build -d
```

### Development

For local development you can run the production service with:

```bash
# rebuild the image if you've made code changes
docker compose up --build
```

The compose configuration now includes a Postgres database and Redis cache alongside the application. Environment variables are provided automatically to instruct the app how to connect:

```text
DATABASE_URL=postgres://gouser:secret@postgres:5432/golearn?sslmode=disable
REDIS_URL=redis://redis:6379/0
```

You can override these in your own `.env` file or when running `docker compose` if you prefer different credentials.

Live reload via the `air` tool was previously supported but has been removed due to reliability issues. You can simply restart the container when you want to pick up code changes.

### Notes on networking

Because the Go service listens directly on port 8080 and all internal services are on the same Docker network, you **do not need** an additional HTTP gateway such as Nginx for local development. The application can talk directly to Postgres and Redis by service name, and you can publish its port to the host as shown above. An external proxy is only necessary if you want features like TLS termination, load balancing multiple copies, or hosting several unrelated services on the same host.
