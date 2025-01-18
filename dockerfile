# Use the official Go image as the base image
FROM golang:1.22.0-alpine as builder

# Set the working directory to /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o weather-api ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/weather-api /app/weather-api
COPY --from=builder /app/cmd/config.json /app/config.json
COPY --from=builder /app/docs /app/docs

CMD ["./weather-api"]