package main

import (
	"fmt"
	"github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/common"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World."))
}

// Routes to enable ---------------------------------------------

func GetOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling GetOrganizationHandler")

	if organization, status, err := GetOrganization(); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting organization. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, organization)
		return
	}
}

func GetAccountsGivenOrgIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling Get all accounts given organization Id::")

	if accounts, status, err := GetAccountsInOrganization(); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting accounts in an organization. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, accounts)
		return
	}
}

func CreateAccountGivenOrgIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("::Create Account given organization Id::")
}

func InviteAccountToOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("::Invite Account given organization Id::")
}

func GetAccountsGivenAccountIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("::Get accounts given account Id::")
}

// ---------------------------------------------------------------------------

func initializeMuxRoutes() {
	httpRouter = mux.NewRouter()

	httpRouter.HandleFunc("/", HelloWorldHandler).Methods("GET")

	org_router := httpRouter.PathPrefix("/org/api/").Subrouter()

	org_router.HandleFunc("/get", GetOrganizationHandler).Methods("GET")

	account_router := httpRouter.PathPrefix("/account/api/").Subrouter()

	account_router.HandleFunc("/create", CreateAccountGivenOrgIdHandler).Methods("POST")
	account_router.HandleFunc("/invite", InviteAccountToOrganizationHandler).Methods("POST")
	account_router.HandleFunc("/get/{account_id}", GetAccountsGivenAccountIdHandler).Methods("GET")
	account_router.HandleFunc("/all", GetAccountsGivenOrgIdHandler).Methods("GET")
}
