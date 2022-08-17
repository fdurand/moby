package network // import "github.com/fdurand/moby/api/server/router/network"

import (
	"context"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
	"github.com/fdurand/moby/api/types/network"
	"github.com/fdurand/moby/libnetwork"
)

// Backend is all the methods that need to be implemented
// to provide network specific functionality.
type Backend interface {
	FindNetwork(idName string) (libnetwork.Network, error)
	GetNetworks(filters.Args, types.NetworkListConfig) ([]types.NetworkResource, error)
	CreateNetwork(nc types.NetworkCreateRequest) (*types.NetworkCreateResponse, error)
	ConnectContainerToNetwork(containerName, networkName string, endpointConfig *network.EndpointSettings) error
	DisconnectContainerFromNetwork(containerName string, networkName string, force bool) error
	DeleteNetwork(networkID string) error
	NetworksPrune(ctx context.Context, pruneFilters filters.Args) (*types.NetworksPruneReport, error)
}

// ClusterBackend is all the methods that need to be implemented
// to provide cluster network specific functionality.
type ClusterBackend interface {
	GetNetworks(filters.Args) ([]types.NetworkResource, error)
	GetNetwork(name string) (types.NetworkResource, error)
	GetNetworksByName(name string) ([]types.NetworkResource, error)
	CreateNetwork(nc types.NetworkCreateRequest) (string, error)
	RemoveNetwork(name string) error
}
