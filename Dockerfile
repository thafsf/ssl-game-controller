FROM node:16.15-alpine3.15 AS build_node
WORKDIR /tmp/ssl-game-controller
COPY public public
COPY src src
COPY babel.config.js .
COPY package.json .
COPY vue.config.js .
COPY yarn.lock .
RUN yarn install
RUN yarn build

FROM golang:1.18-alpine3.15 AS build_go
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-game-controller
COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY go.mod .
COPY go.sum .
COPY --from=build_node /tmp/ssl-game-controller/internal/app/ui/dist internal/app/ui/dist
RUN go install -v ./cmd/ssl-game-controller

# Start fresh from a smaller image
FROM alpine:3.16
COPY --from=build_go /go/bin/ssl-game-controller /app/ssl-game-controller
RUN mkdir -p config && chown -R 1000: config
USER 1000
EXPOSE 8081 10007 10008 10009 10011 10107 10108 10111
ENTRYPOINT ["/app/ssl-game-controller"]
CMD []
