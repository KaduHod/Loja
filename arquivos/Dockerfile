FROM nginx:alpine

RUN apk add npm
RUN apk add nodejs
RUN apk add vim
RUN apk add logrotate

WORKDIR /app

COPY package* ./
COPY . .

RUN cd /app & npm i


COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["/bin/sh", "-c", "npm run dev > server.log 2>&1 & nginx -g 'daemon off;'"]
