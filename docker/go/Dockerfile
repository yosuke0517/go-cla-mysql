FROM golang:1.17-alpine3.14 as dev

LABEL maintainer="https://github.com/yosuke0517"

WORKDIR /go/app

COPY . .

ARG AIR_VERSION=v1.27.3
ARG DLV_VERSION=v1.7.1
ARG MIGRATE_VERSION=v4.15.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git vim curl make && \
  go install github.com/cosmtrek/air@${AIR_VERSION} && \
  go install github.com/go-delve/delve/cmd/dlv@${DLV_VERSION} && \
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@${MIGRATE_VERSION}

RUN set -eux && \
  go build -o go-cla-mysql ./main.go

ENV CGO_ENABLED 0

FROM alpine:3.14

WORKDIR /app

COPY --from=build /go/app/go-cla-mysql .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/go-cla-mysql

CMD ["./go-cla-mysql"]

FROM golang:1.15.7-alpine as builder

ENV ROOT=/go/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY ../.. ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/app/binary"]