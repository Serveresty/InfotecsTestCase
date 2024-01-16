FROM golang:1.21 AS builder

WORKDIR /usr/src/app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/app ./cmd/app/main.go

COPY --from=builder /usr/src/app/bin/app /

EXPOSE 8080

CMD ["/wallet"]