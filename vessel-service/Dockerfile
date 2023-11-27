FROM golang:1.21-alpine AS builder
RUN mkdir /app
# ADD . /app
COPY . /app
WORKDIR /app

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]