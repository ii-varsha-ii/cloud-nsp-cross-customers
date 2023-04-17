package vpc

type CreateVpcRequest struct {
	Name    string                 `json:"name"`
	Region  string                 `json:"region"`
	IpCidr  string                 `json:"ip_cidr"`
	Subnets []*CreateSubnetRequest `json:"subnets"`
}

type CreateSubnetRequest struct {
	Name             string `json:"name"`
	Region           string `json:"region"`
	CidrBlock        string `json:"cidr_block"`
	Access           string `json:"access"`
	VpcId            string `json:"vpc_id"`
	AvailabilityZone string `json:"availability_zone"`
}

type SubnetModel struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	AvailabilityZone string `json:"availability_zone"`
	CidrBlock        string `json:"cidr_block"`
	VpcId            string `json:"vpc_id"`
	RouteTableId     string `json:"route_table_id"`
}

type Vpc struct {
	Id        string         `json:"id"`
	CidrBlock string         `json:"cidr_block"`
	Name      string         `json:"name"`
	Region    string         `json:"region"`
	Subnets   []*SubnetModel `json:"subnets"`
}
