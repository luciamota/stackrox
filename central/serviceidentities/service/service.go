package service

import (
	"bitbucket.org/stack-rox/apollo/central/serviceidentities/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	RegisterServiceServer(grpcServer *grpc.Server)
	RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)

	GetServiceIdentities(ctx context.Context, _ *empty.Empty) (*v1.ServiceIdentityResponse, error)
	CreateServiceIdentity(ctx context.Context, request *v1.CreateServiceIdentityRequest) (*v1.CreateServiceIdentityResponse, error)
	GetAuthorities(ctx context.Context, request *empty.Empty) (*v1.Authorities, error)
}

// New returns a new Service instance using the given DataStore.
func New(storage store.Store) Service {
	return &serviceImpl{
		storage: storage,
	}
}
