package main

import (
	"sync"

	starwars "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/starwars"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		starwars.RunGRPCServer()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		starwars.RunGRPCGatewayServer()
	}()
	wg.Wait()
}
