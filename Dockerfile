FROM golang:1.21-alpine AS build

ENV CGO_ENABLED=1

WORKDIR /app
COPY . .

RUN apk update && \
    apk add --no-cache build-base && \
    go mod download && \
    go build -o /app/api

FROM alpine:latest

WORKDIR /app
ENV CGO_ENABLED=1

COPY --from=build /app/api .

EXPOSE 8080

CMD ["./api"]
