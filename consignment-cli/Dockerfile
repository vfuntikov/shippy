FROM golang:1.21-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# FROM alpine:latest AS production
FROM scratch AS production
COPY --from=builder /app .
CMD ["./app"]