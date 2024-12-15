# ------------------------------------------------------------------------------
# General
# ------------------------------------------------------------------------------

build:
    echo TODO

test:
    echo TODO

# ------------------------------------------------------------------------------
# Go
# ------------------------------------------------------------------------------

go-format:
    go fmt ./...

go-staticcheck:
    go staticcheck ./...

go-vet:
    go vet ./...

# ------------------------------------------------------------------------------
# Docker Commands
# ------------------------------------------------------------------------------

# Build the Docker image
docker-build:
    docker build -t jackplowman/github-pr-analyser:latest .

# Run the analyser in a Docker container, used for testing the github action docker image
docker-run:
    docker run \
      --rm jackplowman/github-pr-analyser:latest

# ------------------------------------------------------------------------------
# Prettier
# ------------------------------------------------------------------------------

prettier-check:
    prettier . --check

prettier-format:
    prettier . --check --write

# ------------------------------------------------------------------------------
# Justfile
# ------------------------------------------------------------------------------

format:
    just --fmt --unstable

format-check:
    just --fmt --check --unstable
