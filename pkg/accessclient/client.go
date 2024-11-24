package accessclient

import (
	"fmt"

	"github.com/waryataw/auth/pkg/accessv1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AccessClient Client for GRPC connect to Auth service.
type AccessClient struct {
	AccessClient accessv1.AccessClient
	Conn         *grpc.ClientConn
}

// New User GRPC Client constructor.
func New(address string) (*AccessClient, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Access server: %w", err)
	}

	client := &AccessClient{
		AccessClient: accessv1.NewAccessClient(conn),
		Conn:         conn,
	}

	return client, nil
}
