package vpc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/common"
	"net/http"
)

const (
	vpcApiPrefix = "/vpc/api"
)

func CreateVpcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling CreateVPCHandler")

	var createVpcRequest CreateVpcRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&createVpcRequest); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("exception--- while creating account in an organization. %v", err))
		return
	}
	defer r.Body.Close()
	if vpc, status, err := CreateVpc(createVpcRequest); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while creating vpc. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, vpc)
		return
	}
}

func GetVpcBasedOnRegionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling GetVpcBasedOnRegion")

	var region = r.URL.Query().Get("region")
	var vpcId = r.URL.Query().Get("vpc_id")

	if vpc, status, err := GetVpcBasedOnRegion(region, &vpcId); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while creating vpc. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, vpc)
		return
	}

}

func initializeMuxRoutes() {
	httpRouter = mux.NewRouter()

	vpcRouter := httpRouter.PathPrefix(vpcApiPrefix).Subrouter()

	vpcRouter.HandleFunc("/create", CreateVpcHandler).Methods("POST")
	vpcRouter.HandleFunc("/get", GetVpcBasedOnRegionHandler).Methods("GET").Queries("region", "vpc_id")
}
