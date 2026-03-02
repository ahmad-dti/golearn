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

### Development with live reload

The `golive` service mounts your source and uses [cosmtrek/air](https://github.com/cosmtrek/air) to rebuild automatically on changes:

```bash
docker compose up golive
```

Edit Go files locally and the container will rebuild/restart the server within a second.

