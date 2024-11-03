package authclient

import (
	"net"

	"github.com/pkg/errors"
)

// GRPCClientConfig GRPCConfig GRPC config.
type GRPCClientConfig interface {
	Address() string
}

type grpcClientConfig struct {
	host string
	port string
}

// NewGRPCClientConfig NewGRPCConfig GRPC config constructor.
func NewGRPCClientConfig(host string, port string) (GRPCClientConfig, error) {
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcClientConfig{
		host: host,
		port: port,
	}, nil
}

// Address Joined Host and Port.
func (cfg *grpcClientConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
