FROM golang:tip-alpine3.24 AS base
WORKDIR /build

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate

RUN go build -o astral ./cmd/web

EXPOSE 8090

CMD [ "/build/astral" ]
