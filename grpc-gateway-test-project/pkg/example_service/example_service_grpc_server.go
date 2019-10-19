package exampleservice

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"log"
	"net"

	gw "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/example_service/proto"
	pb "github.com/zoidbergwill/experiments/grpc-gateway-test-project/pkg/example_service/proto"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint        = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	grpcGatewayServerEndpoint = flag.String("grpc-gateway-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func actuallyRunGRPCServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("Starting http server on :8081")
	return http.ListenAndServe(":8081", mux)
}

// RunGRPCServer should have a comment
func RunGRPCServer() {
	flag.Parse()
	defer glog.Flush()

	if err := actuallyRunGRPCServer(); err != nil {
		glog.Fatal(err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// Echo implements example.YourServiceServer
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.StringMessage{Value: in.GetValue()}, nil
}

// RunGRPCGatewayServer should have a comment
func RunGRPCGatewayServer() {
	fmt.Println("Starting http server on :8080")
	lis, err := net.Listen("tcp", *grpcGatewayServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
