FROM golang:1.21

WORKDIR /usr/src/app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/app cmd/app/main.go

EXPOSE 8080

CMD ["/wallet"]