# ------------------------------------------------------------------------------
# General
# ------------------------------------------------------------------------------

build:
    echo TODO

test:
    go test -coverprofile=coverage.out ./...

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

# ------------------------------------------------------------------------------
# gitleaks
# ------------------------------------------------------------------------------

gitleaks-detect:
    gitleaks detect --source . > /dev/null

# ------------------------------------------------------------------------------
# Git Hooks
# ------------------------------------------------------------------------------

# Install pre commit hook to run on all commits
install-git-hooks:
    cp -f githooks/pre-commit .git/hooks/pre-commit
    cp -f githooks/post-commit .git/hooks/post-commit
    chmod ug+x .git/hooks/*
