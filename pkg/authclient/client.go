package authclient

import (
	"fmt"

	"github.com/waryataw/auth/pkg/authv1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AuthClient Client for GRPC connect to Auth service.
type AuthClient struct {
	authServiceClient authv1.AuthServiceClient
	conn              *grpc.ClientConn
}

// New Auth GRPC Client constructor.
func New(config GRPCClientConfig) (*AuthClient, error) {
	conn, err := grpc.NewClient(
		config.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Auth server: %w", err)
	}

	client := &AuthClient{
		authServiceClient: authv1.NewAuthServiceClient(conn),
		conn:              conn,
	}

	return client, nil
}
