FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.sum ./

COPY . ./

RUN go mod download

RUN go build -o /main server/cmd/main.go

EXPOSE 8080

CMD [ "/main" ]