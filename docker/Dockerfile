FROM golang:1.23-alpine AS build
LABEL authors="Andrei venski"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o authservice ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=build /app/authservice /app/authservice
COPY --from=build /app/migrations /app/migrations
COPY .env /app/.env

EXPOSE 8080

CMD ["./authservice"]
