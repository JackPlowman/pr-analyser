FROM golang:1.22.6

ENV GO111MODULE="on"

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

ENTRYPOINT ["go", "run", "main.go"]
