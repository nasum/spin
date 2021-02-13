FROM golang:1.14.3-alpine as dev

WORKDIR /app

COPY . .

RUN go get -u github.com/cosmtrek/air

FROM node:14-alpine as frontend_build

WORKDIR /build

copy . .

RUN yarn install \
    && yarn run build

FROM golang:1.14.3-alpine as backend_build

WORKDIR /build

COPY . .
COPY --from=frontend_build /build/dist ./dist

RUN go get -u github.com/cosmtrek/air

RUN apk add --no-cache git \
    && go build -o spin

FROM alpine

WORKDIR /app

COPY --from=backend_build /build .
COPY --from=frontend_build /dist .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/spin

CMD ["./spin"]