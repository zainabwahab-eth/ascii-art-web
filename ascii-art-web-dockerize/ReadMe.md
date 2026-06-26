# ASCII-Art-Web-Dockerize

## Description

A continuation of the ASCII-Art-Web project — same Go server, same three banner styles, but now fully containerized with Docker. Instead of needing Go installed locally to run the project, anyone with Docker can build and run it in one command.

## Authors

- Zainab Wahab

## Usage

**Clone the repository:**
```bash
git clone <repo-url>
cd ascii-art-web-dockerize
```

**Build the image:**
```bash
docker build -t ascii-art-web .
```

**Run a container from it:**
```bash
docker run -p 8080:8080 ascii-art-web
```

**Open in your browser:**
```
http://localhost:8080
```

**Clean up unused images/containers (garbage collection):**
```bash
docker system prune
```

## Implementation Details

### Dockerfile

```dockerfile
FROM golang:1.22

LABEL version="1.0"

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o ascii-art-web .

EXPOSE 8080

CMD ["./ascii-art-web"]
```

- `FROM golang:1.22` — uses the official Go image as the base, so the build environment matches exactly what the app needs.
- `WORKDIR /app` — sets where commands run inside the container.
- `COPY go.mod ./` then `COPY . .` — copies dependency files first to take advantage of Docker's build cache, then copies the rest of the project.
- `RUN go build` — compiles the binary inside the container at build time.
- `EXPOSE 8080` — documents which port the app listens on.
- `LABEL` — attaches metadata (maintainer, version, description) to the image, satisfying the requirement to apply metadata to Docker objects.
- `CMD` — the command that runs when a container starts from this image.

### Image vs Container

- **Image** — the built artifact from the Dockerfile, created with `docker build`. It's the blueprint.
- **Container** — a running instance of that image, created with `docker run`. Multiple containers can run from the same image.

This project produces exactly one Dockerfile, one image, and one container — no extra services or databases are needed since the app is a self-contained Go server.

### Garbage Collection

Every `docker build` and every `docker run` leaves artifacts behind — old images, stopped containers, dangling layers. `docker system prune` removes anything not currently in use, which keeps the local Docker environment from filling up with unused objects over time.

## What I Learned

This was my first real hands-on project with Docker, and most of the learning happened before I even got to writing the Dockerfile — I had to actually get Docker installed and working on Ubuntu, which meant adding Docker's official GPG key and repository manually since `docker-desktop` depends on packages that aren't in Ubuntu's default repos.

Once it was running, the core idea that stuck with me is the difference between an **image** (the blueprint) and a **container** (a running instance of that blueprint) — the same mental model as a class versus an object. I also learned that `docker run` both pulls an image if it's missing locally and creates a new container every time you run it, whereas `docker start` just restarts an existing container with its original config — you can't change environment variables on a container that already exists, you have to remove it and run a new one.

Writing the actual Dockerfile taught me about layer caching — copying `go.mod` before the rest of the project so Docker doesn't rebuild dependencies every time application code changes — and `LABEL` for attaching metadata directly to the image, which is what the project required for documenting the build.
