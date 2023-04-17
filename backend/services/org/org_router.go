package org

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/common"
	"net/http"
)

// handles user-side requests and responds to user

const (
	organizationApiPrefix = "/org/api/"
	accountsApiPrefix     = "/account/api/"
	IdRegex               = "{[0-9]*?}"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World."))
}

// Routes to enable ---------------------------------------------

func GetOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling GetOrganizationHandler")

	org := OrganizationModel{}
	if organization, status, err := org.GetOrganization(); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting organization. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, organization)
		return
	}
}

func GetAccountsGivenOrgIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling Get all accounts given organization Id")

	if accounts, status, err := GetAccountsInOrganization(); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting accounts in an organization. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, accounts)
		return
	}
}

func GetAccountsGivenAccountIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling Get account given account Id")

	var accountId = r.URL.Query().Get("account_id")
	acc := AccountModel{}
	if account, status, err := acc.GetAccountGivenAccountId(accountId); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while getting account in an organization. %v", err))
		return
	} else {
		if account == nil {
			common.RespondWithError(w, status, fmt.Sprintf("no accounts associated with the given account id. []"))
			return
		} else {
			common.RespondWithJSON(w, status, account)
			return
		}
	}
}

func InviteAccountToOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling invite account to organization")
	var inviteAccountRequest InviteAccountRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inviteAccountRequest); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("exception--- while creating account in an organization. %v", err))
		return
	}
	defer r.Body.Close()

	if account, status, err := InviteAccountToOrganization(inviteAccountRequest); err != nil {
		common.RespondWithError(w, status, fmt.Sprintf("exception while creating account in an organization. %v", err))
		return
	} else {
		common.RespondWithJSON(w, status, account)
		return
	}
}

// ---------------------------------------------------------------------------

func initializeMuxRoutes() {
	httpRouter = mux.NewRouter()

	httpRouter.HandleFunc("/", HelloWorldHandler).Methods("GET")

	orgRouter := httpRouter.PathPrefix(organizationApiPrefix).Subrouter()
	orgRouter.HandleFunc("/get", GetOrganizationHandler).Methods("GET")

	accountRouter := httpRouter.PathPrefix(accountsApiPrefix).Subrouter()

	accountRouter.HandleFunc("/invite", InviteAccountToOrganizationHandler).Methods("POST")
	accountRouter.HandleFunc("/get", GetAccountsGivenAccountIdHandler).Queries("account_id", IdRegex).Methods("GET")
	accountRouter.HandleFunc("/all", GetAccountsGivenOrgIdHandler).Methods("GET")
}
