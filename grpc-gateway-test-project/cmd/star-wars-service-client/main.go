package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
	proto "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/starwars/proto"
	grpc "google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println(conn)
	client := proto.NewStarWarsServiceClient(conn)
	fmt.Println(client)
	stream, err := client.ListCharacters(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		character, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListCharacters(_) = _, %v", client, err)
		}
		log.Println(character.String())
	}
}
