generate:
	protoc -I/usr/local/include -I. \
	  -I${GOPATH}/src \
	  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:pkg/example \
	  --grpc-gateway_out=logtostderr=true:pkg/example \
	  --swagger_out=logtostderr=true:docs/ \
	  proto/example_service.proto
	protoc -I/usr/local/include -I. \
	  -I${GOPATH}/src \
	  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:pkg/starwars \
	  --grpc-gateway_out=logtostderr=true:pkg/starwars \
	  --swagger_out=logtostderr=true:docs/ \
	  proto/starwars_service.proto
