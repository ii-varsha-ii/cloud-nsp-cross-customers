package org

type InviteAccountRequest struct {
	EmailId   string `json:"email_id"`
	AccountId string `json:"account_id"`
}

type PartiesPair struct {
	Id   string
	Type string
}

type InviteAccountResponse struct {
	Id                  string
	Arn                 string
	Parties             []*PartiesPair
	State               string
	RequestedTimestamp  string
	ExpirationTimestamp string
	Action              string
}
