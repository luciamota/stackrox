package store

import (
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"github.com/boltdb/bolt"
)

const benchmarkBucket = "benchmarks"

// Store provides storage functionality for alerts.
type Store interface {
	GetBenchmark(id string) (*v1.Benchmark, bool, error)
	GetBenchmarks(request *v1.GetBenchmarksRequest) ([]*v1.Benchmark, error)
	AddBenchmark(benchmark *v1.Benchmark) (string, error)
	UpdateBenchmark(benchmark *v1.Benchmark) error
	RemoveBenchmark(id string) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) Store {
	bolthelper.RegisterBucket(db, benchmarkBucket)
	return &storeImpl{
		DB: db,
	}
}
