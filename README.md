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

Live reload via the `air` tool was previously supported but has been removed due to reliability issues. You can simply restart the container when you want to pick up code changes.
