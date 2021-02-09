
FROM node:14-alpine as frontend_build

WORKDIR /build

copy . .

RUN yarn install \
    && yarn run build

FROM golang:1.14.3-alpine as backend_build

WORKDIR /build

COPY . .
COPY --from=frontend_build /build/dist ./dist

RUN apk add --no-cache git \
    && go build -o app

FROM alpine

WORKDIR /app

COPY --from=backend_build /build .


RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]