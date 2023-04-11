package main

import (
	"net/http"
)

func GetOrganization() (*OrganizationModel, int, error) {
	organization, err := getOrganization()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if organization != nil {
		orgObj := OrganizationModel{
			Id:                 organization.Id,
			Arn:                organization.Arn,
			MasterAccountId:    organization.MasterAccountId,
			MasterAccountArn:   organization.MasterAccountArn,
			MasterAccountEmail: organization.MasterAccountEmail,
		}
		return &orgObj, http.StatusOK, nil
	}

	return &OrganizationModel{}, http.StatusOK, nil
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
