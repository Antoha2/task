package auth

import (
	"context"

	pb "github.com/antoha2/task/pkg/auth/protoAPI"
	"google.golang.org/grpc/grpclog"
)

func (a *authService) ParseToken(ctx context.Context, token string) (int, error) {

	request := &pb.ParseTokenRequest{
		Token: token,
	}
	response, err := a.Client.ParseToken(ctx, request)
	if err != nil {
		grpclog.Fatalf("ParseToken() - fail to dial: %v", err)
		return 0, err
	}
	return int(response.Id), err
}

/* func (a *authService) ParseToken(ctx context.Context, token string) (int, error) {

	request := &pb.ParseTokenRequest{
		Token: token,
	}
	response, err := a.Client.ParseToken(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("ParseToken() - fail to dial: %v", err)
		return 0, err
	}
	return int(response.Id), fmt.Errorf(response.Err)
} */

func (a *authService) GetRoles(ctx context.Context, Id int) ([]string, error) {

	request := &pb.GetRolesRequest{
		Id: int32(Id),
	}
	response, err := a.Client.GetRoles(ctx, request)
	if err != nil {
		grpclog.Fatalf("GetRoles() - fail to dial: %v", err)
		return nil, err
	}
	return response.Roles, err
}
