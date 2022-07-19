FROM golang:1.17.0-alpine3.14 as builder

WORKDIR /src

COPY go.mod /src/
COPY go.sum /src/
RUN go mod download
COPY cmd/ /src/cmd/

RUN go mod download \
    && GOOS=linux go build -v -o bin/sms github.com/cheald/sentry-mattermost-integration/cmd/sms

FROM alpine:3.14.2

COPY --from=builder /src/bin/sms /usr/bin/go-sms

ENV GIN_MODE=release
ENV SMS_PORT=1323

EXPOSE 1323

CMD ["/usr/bin/go-sms"]
