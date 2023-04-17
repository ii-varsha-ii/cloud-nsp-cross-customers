package gateway

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	"net/http"
)

type TransitGatewayModel struct {
	Id           string
	Arn          string
	OwnerId      string
	Description  string
	State        string
	CreationTime string
}

func (tgw *TransitGatewayModel) GetTransitGateway(region string) (*TransitGatewayModel, int, error) {
	regionRequest := pb.RegionRequest{Region: region}
	response, err := getTransitGateway(&regionRequest)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if response.Id != "" {
		tgw.Id = response.Id
		tgw.Arn = response.Arn
		tgw.State = response.State
		tgw.OwnerId = response.OwnerId
		tgw.Description = response.Description
		tgw.CreationTime = response.CreationTime
	}

	return tgw, http.StatusOK, nil
}
