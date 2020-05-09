FROM golang:1.14-alpine as build

LABEL Maintainer="me@andreamedda.com"

ARG SERVICE='local'
ARG COMMIT='local'

WORKDIR /build

COPY . .

RUN apk --update --no-cache add ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -ldflags "-s -w -X main.commit=${COMMIT}" -o /app ./cmd/toy

ENTRYPOINT ["/app"]
