FROM golang:tip-alpine3.24 AS build
WORKDIR /build

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GOOSE_DBSTRING
ENV GOOSE_DRIVER=postgres
ENV GOOSE_MIGRATION_DIR=data/migrations
RUN goose up

RUN templ generate
RUN CGO_ENABLED=0 go build -o astral ./cmd/web

FROM alpine:3.22
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /build/astral .
EXPOSE 8090
CMD ["./astral"]
