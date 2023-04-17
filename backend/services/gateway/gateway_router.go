package gateway

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/common"
	"net/http"
)

const (
	gatewayApiPrefix = "/gateway/api"
)

func GetTransitGatewayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling GetTransitGateway")

	var region = r.URL.Query().Get("region")

	var tgw *TransitGatewayModel
	if tgw, status, err := tgw.GetTransitGateway(region); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting transit gateway. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, tgw)
		return
	}

}

func initializeMuxRoutes() {
	httpRouter = mux.NewRouter()

	gatewayRouter := httpRouter.PathPrefix(gatewayApiPrefix).Subrouter()

	gatewayRouter.HandleFunc("/get", GetTransitGatewayHandler).Methods("GET")
}
