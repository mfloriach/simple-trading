# Simple Trading Platform

## Tools
* Buff
* TimeScaleDB
* MongoDB
* GRPC

## Installation
```
$ buff generate
$ docker compose up -d
$ go run server/cmd/migrate/main.go
```

## Run 
```
$ go run server/cmd/grpc/main.go
```

## Docs
[API](http://localhost:8081)