FROM golang:1.25.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

FROM alpine:3.19

WORKDIR /app

COPY --from=builder app/server .

COPY configs/config.yml ./configs/config.yml

EXPOSE 8000

CMD ["./server"]
