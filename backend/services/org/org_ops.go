package org

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	"net/http"
)

// middle man between router and the grpc client
// modifying requests given by router to support grcp client call
// responds to router with raw response

type OrganizationModel struct {
	Id                 string
	Arn                string
	MasterAccountArn   string
	MasterAccountId    string
	MasterAccountEmail string
}

type AccountModel struct {
	Id              string
	Arn             string
	Email           string
	Name            string
	Status          string
	JoinedMethod    string
	JoinedTimestamp string
}

func (org *OrganizationModel) GetOrganization() (*OrganizationModel, int, error) {
	organization, err := getOrganization()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if organization != nil {
		org.Id = organization.Id
		org.Arn = organization.Arn
		org.MasterAccountArn = organization.MasterAccountArn
		org.MasterAccountEmail = organization.MasterAccountEmail
		org.MasterAccountId = organization.MasterAccountId
	}

	return org, http.StatusOK, nil
}

func (acc *AccountModel) GetAccountGivenAccountId(accountId string) (*AccountModel, int, error) {
	account, err := getAccountGivenAccountId(&pb.AccountId{
		AccountId: accountId,
	})

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if account != nil {
		acc.Name = account.Name
		acc.Id = account.Id
		acc.Email = account.Email
		acc.Arn = account.Arn
		acc.Status = account.Status
		acc.JoinedMethod = account.JoinedMethod
		acc.JoinedTimestamp = account.JoinedTimestamp
	}

	return acc, http.StatusOK, nil
}

func GetAccountsInOrganization() ([]*AccountModel, int, error) {
	listOfAccounts, err := getAccountsInOrganization()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var accounts []*AccountModel
	if listOfAccounts != nil {
		for _, account := range listOfAccounts {
			accounts = append(accounts, &AccountModel{
				Id:              account.Id,
				Arn:             account.Arn,
				Email:           account.Email,
				Name:            account.Name,
				Status:          account.Status,
				JoinedMethod:    account.JoinedMethod,
				JoinedTimestamp: account.JoinedTimestamp,
			})
		}
	}
	return accounts, http.StatusOK, nil
}

func InviteAccountToOrganization(model InviteAccountRequest) (*InviteAccountResponse, int, error) {
	response, err := inviteAccountToOrganization(&pb.InviteAccountRequest{
		EmailId:   &model.EmailId,
		AccountId: &model.AccountId,
	})

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if response != nil {
		accountResponse := InviteAccountResponse{
			Id:                  response.Id,
			Arn:                 response.Arn,
			RequestedTimestamp:  response.RequestedTimestamp,
			ExpirationTimestamp: response.ExpirationTimestamp,
			State:               response.State,
			Action:              response.Action,
		}
		for _, val := range response.Parties {
			var partyPair PartiesPair
			partyPair.Id = val.Id
			partyPair.Type = val.Type
			accountResponse.Parties = append(accountResponse.Parties, &partyPair)
		}
		return &accountResponse, http.StatusOK, nil
	}
	return &InviteAccountResponse{}, http.StatusOK, nil
}
