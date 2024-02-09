FROM golang:1.21

WORKDIR /url-shortener
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 3000

CMD go run ./cmd/main.go -config="./config/config.yaml"