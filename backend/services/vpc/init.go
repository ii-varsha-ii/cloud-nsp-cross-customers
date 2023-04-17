package vpc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	"google.golang.org/grpc"
)

var (
	ctx               = context.Background()
	httpServerHost    = "localhost"
	httpServerPort    = 8000
	gRPCVpcClientHost = "localhost"
	gRPCVpcClientPort = 50052
	httpRouter        *mux.Router
)

var (
	grpcVpcClient pb.VpcServiceClient
)

// User selections
//var (
//	vpc_id,
//	subnet_id,
//	vpc_id2,
//	subnet_id2
//)

func initializeHTTPRouter() error {
	initializeMuxRoutes()
	if httpRouter == nil {
		return fmt.Errorf("http router not initialized")
	}
	return nil
}

func initializeGrpcVpcClient() error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCVpcClientHost, gRPCVpcClientPort), grpc.WithInsecure())

	if err != nil {
		return fmt.Errorf("vpc grpc client initailization failed: %v", err)
	}

	grpcVpcClient = pb.NewVpcServiceClient(conn)
	return nil
}

func initialize() error {
	if err := initializeGrpcVpcClient(); err != nil {
		return fmt.Errorf("gRPC client initialization error. %v", err)
	}

	if err := initializeHTTPRouter(); err != nil {
		return fmt.Errorf("Http router initilization error, %v", err)
	}
	return nil
}

func RunVpcServer() {
	if err := initialize(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("Server running at %s:%d", httpServerHost, httpServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", httpServerHost, httpServerPort), httpRouter))
}
