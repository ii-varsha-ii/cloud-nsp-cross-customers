package gateway

import (
	"fmt"
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
)

func getTransitGateway(region *pb.RegionRequest) (*pb.TransitGateway, error) {
	result, err := grpcGatewayClient.GetTransitGateway(ctx, region)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
