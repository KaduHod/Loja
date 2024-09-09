FROM nginx:alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache go

RUN apk add vim

RUN apk add logrotate

COPY go.mod go.sum ./

RUN go mod tidy

RUN go build -o /usr/bin/loja-api

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["/bin/sh", "-c", "/usr/bin/loja-api > server.log 2>&1 & nginx -g 'daemon off;'"]
