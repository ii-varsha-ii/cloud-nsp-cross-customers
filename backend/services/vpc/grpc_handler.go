package vpc

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
)

func createVpcWithSubnets(request *pb.CreateVpcRequest) (*pb.Vpc, error) {
	response, err := grpcVpcClient.CreateVpcWithSubnets(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func getVpcBasedOnRegion(in *pb.RegionAndVpcRequest) (*pb.ListOfVpcs, error) {
	response, err := grpcVpcClient.GetVpcBasedOnRegion(ctx, in)

	if err != nil {
		return nil, err
	}
	return response, nil
}
