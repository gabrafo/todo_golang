FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todo_app ./cmd

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

COPY --from=builder /todo_app /usr/local/bin/todo_app
EXPOSE 8080
CMD ["/usr/local/bin/todo_app"]
