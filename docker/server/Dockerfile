FROM golang:1.22-alpine3.19

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/spf13/cobra-cli@latest

CMD ["air", "-c", ".air.server.toml"]