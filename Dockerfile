FROM golang:1.21-rc-alpine3.17

WORKDIR /app

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN go build -o /app/main /app/main.go

CMD [ "/app/main" ]
