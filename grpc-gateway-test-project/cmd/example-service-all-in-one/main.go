package main

import (
	"sync"

	exampleservice "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/example"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		exampleservice.RunGRPCServer()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		exampleservice.RunGRPCGatewayServer()
	}()
	wg.Wait()
}
