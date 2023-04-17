package org

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
	ctx                        = context.Background()
	httpServerHost             = "localhost"
	httpServerPort             = 8000
	gRPCOrganizationClientHost = "localhost"
	gRPCOrganizationClientPort = 50051
	httpRouter                 *mux.Router
)

var (
	grpcOrganizationClient pb.OrganizationServiceClient
)

func initializeHTTPRouter() error {
	initializeMuxRoutes()
	if httpRouter == nil {
		return fmt.Errorf("http router not initialized")
	}
	return nil
}

func initializeGrpcOrganizationClient() error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCOrganizationClientHost, gRPCOrganizationClientPort), grpc.WithInsecure())

	if err != nil {
		return fmt.Errorf("organization grpc client initailization failed: %v", err)
	}

	grpcOrganizationClient = pb.NewOrganizationServiceClient(conn)
	return nil
}

func initialize() error {
	if err := initializeGrpcOrganizationClient(); err != nil {
		return fmt.Errorf("gRPC client initialization error. %v", err)
	}

	if err := initializeHTTPRouter(); err != nil {
		return fmt.Errorf("Http router initilization error, %v", err)
	}
	return nil
}

func RunOrgServer() {
	if err := initialize(); err != nil {
		log.Panic(err)
	}

	fmt.Printf("Server running at %s:%d", httpServerHost, httpServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", httpServerHost, httpServerPort), httpRouter))
}
