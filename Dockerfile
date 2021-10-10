# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

ARG VERSION

RUN go mod download

RUN go build -o /bot -ldflags "-X `go list ./version`.VersionStr=${VERSION}"

CMD [ "/bot" ]