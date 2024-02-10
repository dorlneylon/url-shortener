FROM golang:1.21

WORKDIR /url-shortener
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 3000
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

CMD go run ./cmd/main.go -config="./config/config.yaml"