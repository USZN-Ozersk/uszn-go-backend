FROM golang:alpine as build
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -v ./cmd/apiserver

FROM scratch
COPY --from=build /app/apiserver /app/
COPY config/config-prod.toml /app/
WORKDIR /app
EXPOSE 8080
CMD ["/app/apiserver", "-config-path", "config-prod.toml"]version: "3.7"
services:
    api:
        build: ./uszn-go-backend
        ports: 
            - "81:81"
        container_name: api
        networks: 
            - app-network
        volumes:
            - certbot-etc:/etc/letsencrypt
            - certbot-var:/var/lib/letsencrypt
        depends_on: 
            - pgdb
    pgdb:
        image: postgres
        restart: unless-stopped
        ports:
            - "5432:5432"
        environment:
            POSTGRES_USER: api
            POSTGRES_PASSWORD: manaraga
            PGDATA: /var/lib/pgsql/data
        container_name: pgdb
        volumes: 
            - "/var/lib/site/data:/var/lib/pgsql/data"
        networks:
            - app-network
    vue:
        build: ./uszn-vue-frontend
        ports: 
            - "80:80"
            - "443:443"
        depends_on: 
            - api
        container_name: vue
        volumes: 
            - "/var/lib/site/upload:/public_html/upload"
            - web-root:/usr/share/nginx/html
            - certbot-etc:/etc/letsencrypt
            - certbot-var:/var/lib/letsencrypt
            - dhparam:/etc/ssl/certs
        networks:
            - app-network
    certbot:
        image: certbot/certbot
        container_name: certbot
        volumes:
            - "web-root:/usr/share/nginx/html"
            - "certbot-etc:/etc/letsencrypt"
            - "certbot-var:/var/lib/letsencrypt"
        depends_on:
            - vue
        networks:
            - app-network
        #command: certonly --webroot --webroot-path=/usr/share/nginx/html --email ksz@ozerskadm.ru --agree-tos --no-eff-email --force-renewal -d usznozersk.ru -d www.usznozersk.ru
        #--staging
volumes:
    certbot-etc:
    certbot-var:
    web-root:
    dhparam:
        driver: local
        driver_opts:
            type: none
            device: /var/lib/site/dhparam
            o: bind
networks:
    app-network:
        driver: bridge
        driver_opts:
            com.docker.network.enable_ipv6: "false"
            com.docker.network.bridge.name: "app-network"