FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/tui

FROM alpine

COPY --from=builder /go/bin/tui /usr/local/bin/tui

EXPOSE 42069

CMD ["/usr/local/bin/tui"]