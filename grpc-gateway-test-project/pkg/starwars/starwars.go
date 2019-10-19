package example

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"log"
	"net"

	gw "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/starwars/proto"
	pb "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/starwars/proto"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint        = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
	grpcGatewayServerEndpoint = flag.String("grpc-gateway-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func actuallyRunGRPCGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterStarWarsServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Printf("Starting http GRPC gateway server on %s\n", *grpcGatewayServerEndpoint)
	return http.ListenAndServe(*grpcGatewayServerEndpoint, mux)
}

// RunGRPCGatewayServer should have a comment
func RunGRPCGatewayServer() {
	flag.Parse()
	defer glog.Flush()

	if err := actuallyRunGRPCGatewayServer(); err != nil {
		glog.Fatal(err)
	}
}

// server is used to implement starwars.StarWarsServer.
type server struct{}

var characters = []*pb.Character{
	{
		Id:        1000,
		Name:      "Luke Skywalker",
		Friends:   []*pb.Character{{Id: 1002}, {Id: 1003}, {Id: 2000}, {Id: 2001}},
		AppearsIn: []pb.Episode{pb.Episode_NEWHOPE, pb.Episode_EMPIRE, pb.Episode_JEDI},
		Species:   pb.Species_HUMAN,
	},
	{
		Id:        1001,
		Name:      "Darth Vader",
		Friends:   []*pb.Character{{Id: 1004}},
		AppearsIn: []pb.Episode{pb.Episode_NEWHOPE, pb.Episode_EMPIRE, pb.Episode_JEDI},
		Species:   pb.Species_HUMAN,
	},
	{
		Id:        1002,
		Name:      "Han Solo",
		Friends:   []*pb.Character{{Id: 1000}, {Id: 1003}, {Id: 2001}},
		AppearsIn: []pb.Episode{pb.Episode_NEWHOPE, pb.Episode_EMPIRE, pb.Episode_JEDI},
		Species:   pb.Species_HUMAN,
	},
	{
		Id:        1003,
		Name:      "Leia Organa",
		Friends:   []*pb.Character{{Id: 1000}, {Id: 1003}, {Id: 2000}, {Id: 2001}},
		AppearsIn: []pb.Episode{pb.Episode_NEWHOPE, pb.Episode_EMPIRE, pb.Episode_JEDI},
		Species:   pb.Species_HUMAN,
	},
	{
		Id:        1004,
		Name:      "Wilhuff Tarkin",
		Friends:   []*pb.Character{{Id: 1001}},
		AppearsIn: []pb.Episode{pb.Episode_NEWHOPE},
		Species:   pb.Species_HUMAN,
	},
}

// ListCharacters implements starwars.StarWars
func (s *server) ListCharacters(in *empty.Empty, stream pb.StarWarsService_ListCharactersServer) error {
	// log.Printf("Received: %v", in.GetValue())
	for _, character := range characters {
		if err := stream.Send(character); err != nil {
			return err
		}
	}
	return nil
}

// RunGRPCServer should have a comment
func RunGRPCServer() {
	fmt.Printf("Starting grpc server on %s\n", *grpcServerEndpoint)
	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStarWarsServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
