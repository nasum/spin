FROM golang:1.14.3-alpine as build

WORKDIR /build

copy . .

RUN apk add --no-cache git \
    && go build -o app

FROM alpine

WORKDIR /app

COPY --from=build /build .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]