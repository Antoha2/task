package auth

import (
	"context"

	pb "github.com/antoha2/task/pkg/auth/protoAPI"
	"google.golang.org/grpc"
)

type AuthService interface {
	ParseToken(ctx context.Context, token string) (int, error)
	GetRoles(ctx context.Context, id int) ([]string, error)
}

type authService struct {
	Client pb.TaskServiceClient
}

func NewAuthService(conn *grpc.ClientConn) *authService {
	client := pb.NewTaskServiceClient(conn)
	return &authService{Client: client}
}
