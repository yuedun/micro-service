# micro-service
go语言微服务

生成*.pb.go代码
```shell
$ protoc protos/method.proto --go_out=plugins=grpc:.
$ protoc --proto_path=. --micro_out=. --go_out=. proto/example/example.proto
```

# Micro Service

This is the Micro service

Generated with

```
micro new micro --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.micro
- Type: srv
- Alias: micro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./micro-srv
```

Build a docker image
```
make docker
```
启动服务和测试
```shell
go run main.go
```
```shell
micro call go.micro.srv.micro-service Example.Call '{"name": "John"}'
```