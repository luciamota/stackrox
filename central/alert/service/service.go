package service

import (
	"context"

	"bitbucket.org/stack-rox/apollo/central/alert/datastore"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

	GetAlert(ctx context.Context, request *v1.ResourceByID) (*v1.Alert, error)
	ListAlerts(ctx context.Context, request *v1.ListAlertsRequest) (*v1.ListAlertsResponse, error)
	GetAlertsGroup(ctx context.Context, request *v1.ListAlertsRequest) (*v1.GetAlertsGroupResponse, error)
	GetAlertsCounts(ctx context.Context, request *v1.GetAlertsCountsRequest) (*v1.GetAlertsCountsResponse, error)
	GetAlertTimeseries(ctx context.Context, req *v1.ListAlertsRequest) (*v1.GetAlertTimeseriesResponse, error)
}

// New returns a new Service instance using the given DataStore.
func New(datastore datastore.DataStore) Service {
	return &serviceImpl{
		datastore: datastore,
	}
}
