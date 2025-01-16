# Use the official Go image as the base image
FROM golang:1.16-alpine as builder

# Set the working directory to /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o weather-api ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/weather-api .

EXPOSE 8080

CMD ["./go-weather-api"]