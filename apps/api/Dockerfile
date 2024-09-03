FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum /app/
COPY /apps/api/go.mod /apps/api/go.sum /app/apps/api/

RUN go mod download 

COPY /apps/api /app/apps/api

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /temp/api /app/apps/api/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY /apps/api/configs/config.yaml /app/config.yaml
COPY --from=builder /temp/api /app/api

EXPOSE 8000

CMD ["/app/api", "--config", "/app/config.yaml"]
