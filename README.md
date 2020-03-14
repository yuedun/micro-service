# micro-service
go语言微服务

## Getting Started

- [教程](https://micro.mu/docs/go-helloworld.html)
- [用户服务](https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro/part1)

## 依赖

grpc v1.25.1，
v1.27.1会有类型这样的错误`undefined: resolver.BuildOption`,`undefined: resolver.ResolveNowOption`

## Usage
生成*.pb.go代码
```shell
$ protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
```
启动服务和测试
```shell
go run main.go
```
```shell
go run client/main.go
```

### 开启api服务
```shell
go run main.go
```
```shell
go run api/api.go
```
```shell
micro api --handler=api
```
```shell
curl "http://localhost:8080/user/say/hello?name=huohuo"
```
返回
{"id":"123","password":"dslhgfoif40u9b9","username":"huohuo"}

### api层说明：
API模块，它的身份其实就是一个网关或者代理层，它的能力就是让一个单一的入口可以去访问微服务。API工作在我们的软件服务架构的边缘。
![api作用](https://micro.mu/docs/images/api.png)
## 服务注册服务
v2.0版本默认使用的服务注册发现是**mdns**。
Micro内置了mDNS组播系统，这是一种零依赖的服务注册发现机制，它是区别于有注册中心的替代方案。
通过在启动指令中传入--registry=mdns 或者在环境变量中设置MICRO_REGISTRY=mdns。
其实也可以不传，早期版本的go-micro默认注册中心是consul，现在换成了mdns
mDNS（多播DNS）是一种局域网内使用的DNS机制，他的大致原理如下：当有新的节点加入局域网的时候，如果打开了mDNS，就主动向局域网其他所有节点广播，自己提供的服务（域名是什么、ip地址是什么、端口号是什么）, 这样我们任何一个节点都知道局域网提供了什么服务。
所以生产环境需要其他中间件，如**consul**，**etcd**。