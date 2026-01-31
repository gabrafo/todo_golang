FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todo_app ./cmd
# Final image: keep Go toolchain so air can rebuild inside container
FROM golang:1.25-alpine
RUN apk add --no-cache ca-certificates

WORKDIR /src

COPY --from=builder /src /src
COPY --from=builder /todo_app /usr/local/bin/todo_app

COPY --from=builder /go/bin/air /go/bin/air
ENV PATH="/go/bin:${PATH}"

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
