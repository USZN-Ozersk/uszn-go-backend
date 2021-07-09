FROM golang:latest as build
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -v ./cmd/apiserver

FROM scratch
COPY --from=build /app/apiserver /app/
#COPY config/config-prod.toml /app/
WORKDIR /app
EXPOSE 8080
CMD ["/app/apiserver", "-config-path", "config-prod.toml"]