FROM docker.io/golang:1.20 as builder

RUN mkdir /app

COPY cmd /app/cmd
COPY sql /app/sql
COPY templates /app/templates

WORKDIR /app

RUN go mod init github.com/calvarado2004/go-testing/go-webapp &&  go mod tidy

RUN CGO_ENABLED=0 go build -o goWebAppTesting ./cmd/web

RUN chmod +x /app/goWebAppTesting

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app /app

CMD [ "/app/goWebAppTesting"]