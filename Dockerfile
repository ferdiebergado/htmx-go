FROM golang:1.22.6-bookworm AS build

ARG APP_ENV
ARG APP_PORT

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

FROM bookworm:12.6

WORKDIR /usr/src/app

COPY --from=build /usr/local/bin/app .

COPY ./templates ./templates

EXPOSE ${APP_PORT}

CMD ["./app"]
