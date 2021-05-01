FROM golang:1.14.3 as dev

WORKDIR /app

COPY . .

RUN go get -u github.com/cosmtrek/air
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /bin/migrate

FROM node:14-alpine as frontend_build

WORKDIR /build

COPY . .

RUN yarn install \
    && yarn run build

FROM golang:1.14.3-alpine as backend_build

WORKDIR /build

COPY . .
COPY --from=frontend_build /build/dist ./dist

RUN apk add --no-cache git \
    && go build -o spin ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=backend_build /build .
COPY --from=frontend_build /dist .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/spin

CMD ["./spin"]
