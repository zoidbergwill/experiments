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

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
)

func httpCall() {
	if err := view.Register(
		// Register a few default views.
		ochttp.ClientSentBytesDistribution,
		ochttp.ClientReceivedBytesDistribution,
		ochttp.ClientRoundtripLatencyDistribution,
		// Register a custom view.
		&view.View{
			Name:        "httpclient_latency_by_path",
			TagKeys:     []tag.Key{ochttp.KeyClientPath},
			Measure:     ochttp.ClientRoundtripLatency,
			Aggregation: ochttp.DefaultLatencyDistribution,
		},
	); err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &ochttp.Transport{},
	}

}

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
