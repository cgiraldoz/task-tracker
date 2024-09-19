FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o task-tracker main.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/task-tracker /build/task-tracker
CMD ["./task-tracker"]
