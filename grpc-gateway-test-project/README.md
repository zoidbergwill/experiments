# Example GRPC Project

https://github.com/grpc-ecosystem/grpc-gateway

## Quickstart

### Echo Server

In one tab:

```
go run cmd/your-service/main.go
```

In another:

```
# go run cmd/your-service-grpc-server/main.go
```

In a third tab:

```
# no_proxy='*' curl -k http://127.0.0.1:8081/v1/example/echo -X POST -d '{"value": "potato"}'
{"value":"potato"}
```

In the second, you should get:

```
# go run cmd/your-service-grpc-server/main.go
2019/08/29 13:45:18 Received: potato
```
