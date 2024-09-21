FROM golang:alpine AS builder

LABEL org.opencontainers.image.description="Task Tracke: A simple task tracker application"
LABEL org.opencontainers.image.source=https://github.com/cgiraldoz/task-tracker
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o task-tracker main.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/task-tracker /build/task-tracker
CMD ["./task-tracker"]
