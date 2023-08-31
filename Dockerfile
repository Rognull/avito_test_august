FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest

ENV SERV_PORT=8080
ENV SERV_DBUSER=kirill
ENV SERV_DBPASS=password
ENV SERV_DBHOST=postgres
ENV SERV_DBPORT=5433
ENV SERV_DBNAME=test

COPY --from=builder /app/main /main
EXPOSE 8080

CMD ["./main"]