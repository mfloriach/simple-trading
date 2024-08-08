# Simple Trading Platform
Simple platform for trading written in golang, grpc and time base database. 

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
```
$ docker compose up -d
```
[API](http://localhost:8081)