FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -v ./cmd/apiserver
EXPOSE 8080
CMD ["/app/apiserver"]