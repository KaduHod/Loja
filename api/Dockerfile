FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /app/loja-api

FROM nginx:alpine

RUN apk add --no-cache go

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/loja-api /usr/bin/loja-api
COPY --from=builder /app /app

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["/bin/sh", "-c", "loja-api & nginx -g 'daemon off;'"]
