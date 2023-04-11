package main

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func getOrganization() (*pb.Organization, error) {
	in := new(emptypb.Empty)
	result, err := grpcOrganizationClient.GetOrganization(ctx, in)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func getAccountsInOrganization() ([]*pb.Account, error) {
	in := new(emptypb.Empty)
	result, err := grpcOrganizationClient.GetAccountsInOrganization(ctx, in)

	accounts := result.GetAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
