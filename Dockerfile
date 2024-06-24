FROM golang:latest

COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/main.go

RUN chmod +x wait-for-postgres.sh

CMD ["./app"]