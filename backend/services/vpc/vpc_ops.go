package vpc

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	"net/http"
)

func CreateVpc(req CreateVpcRequest) (*Vpc, int, error) {

	requestObj := pb.CreateVpcRequest{
		Region: req.Region,
		Name:   req.Name,
		IpCidr: req.IpCidr,
	}

	var listOfSubnetsPb []*pb.CreateSubnetRequest
	for _, subnet := range req.Subnets {

		listOfSubnetsPb = append(listOfSubnetsPb, &pb.CreateSubnetRequest{
			Name:             subnet.Name,
			AvailabilityZone: subnet.AvailabilityZone,
			CidrBlock:        subnet.CidrBlock,
			VpcId:            subnet.VpcId,
			Access:           subnet.Access,
			Region:           subnet.Region,
		})
	}

	requestObj.Subnets = listOfSubnetsPb

	response, err := createVpcWithSubnets(&requestObj)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	//print(response)
	var listOfSubnets []*SubnetModel

	responseVpc := &Vpc{
		Id:        response.Id,
		CidrBlock: response.CidrBlock,
		Name:      response.Name,
		Region:    response.Region,
	}
	for _, subnet := range response.Subnets {
		listOfSubnets = append(listOfSubnets, &SubnetModel{
			Id:               subnet.Id,
			Name:             subnet.Name,
			AvailabilityZone: subnet.AvailabilityZone,
			CidrBlock:        subnet.CidrBlock,
			VpcId:            subnet.VpcId,
			RouteTableId:     subnet.RouteTableId,
		})
	}
	responseVpc.Subnets = listOfSubnets
	return responseVpc, http.StatusOK, nil

}

func GetVpcBasedOnRegion(region string, vpc *string) ([]*Vpc, int, error) {
	regionVpcRequest := pb.RegionAndVpcRequest{
		Region: region,
		VpcId:  vpc,
	}

	response, err := getVpcBasedOnRegion(&regionVpcRequest)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var listOfVpc []*Vpc

	for _, vpc := range response.Vpc {
		var tempVpc = &Vpc{
			Id:        vpc.Id,
			CidrBlock: vpc.CidrBlock,
			Name:      vpc.Name,
			Region:    vpc.Region,
		}

		for _, subnet := range vpc.Subnets {
			tempVpc.Subnets = append(tempVpc.Subnets, &SubnetModel{
				Id:               subnet.Id,
				Name:             subnet.Name,
				AvailabilityZone: subnet.AvailabilityZone,
				CidrBlock:        subnet.CidrBlock,
				VpcId:            subnet.VpcId,
				RouteTableId:     subnet.RouteTableId,
			})
		}
		listOfVpc = append(listOfVpc, tempVpc)
	}

	return listOfVpc, http.StatusOK, nil
}
