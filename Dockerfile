FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -o server -ldflags="-w -s" ./cmd/main.go

FROM scratch

COPY --from=builder /app/server /server

EXPOSE 9000

CMD [ "/server" ]