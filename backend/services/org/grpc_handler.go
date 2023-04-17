package org

import (
	pb "github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/go_pb_files"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// plain old grpc client

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

func getAccountGivenAccountId(id *pb.AccountId) (*pb.Account, error) {
	account, err := grpcOrganizationClient.GetAccountInOrganizationBasedOnId(ctx, id)

	if err != nil {
		return nil, err
	}
	return account, nil
}

func inviteAccountToOrganization(request *pb.InviteAccountRequest) (*pb.InviteAccountResponse, error) {
	response, err := grpcOrganizationClient.InviteAccountToOrganization(ctx, request)

	if err != nil {
		return nil, err
	}
	return response, nil
}
