package userclient

import (
	"fmt"

	"github.com/waryataw/auth/pkg/userv1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UserClient Client for GRPC connect to Auth service.
type UserClient struct {
	UserServiceClient userv1.UserServiceClient
	Conn              *grpc.ClientConn
}

// New User GRPC Client constructor.
func New(address string) (*UserClient, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Auth server: %w", err)
	}

	client := &UserClient{
		UserServiceClient: userv1.NewUserServiceClient(conn),
		Conn:              conn,
	}

	return client, nil
}
