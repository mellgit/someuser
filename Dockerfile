FROM golang:1.24.2-alpine AS builder

WORKDIR /etc/build-someuser

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . .
RUN go build -o someuser
RUN chmod +x someuser

FROM alpine AS runner

RUN apk update && \
    apk upgrade --no-cache && \
    apk --no-cache add ffmpeg && \
    adduser -D -u 1001 -h /home/app -s /bin/sh app

WORKDIR /home/app

COPY --from=builder --chown=app:app /etc/build-someuser/someuser .
COPY config.yml /home/app/config.yml
COPY .env /home/app/.env

USER app

RUN chmod +x /home/app/someuser

CMD ["./someuser"]

