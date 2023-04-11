package main

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
