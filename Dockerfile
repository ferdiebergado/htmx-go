FROM golang:1.22.6-alpine3.19 AS build

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=build /app/main .

# COPY ./assets ./assets
COPY ./templates ./templates

EXPOSE 8888

CMD ["./main"]
