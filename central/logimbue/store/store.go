package store

import (
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"github.com/boltdb/bolt"
)

const logsBucket = "logs"

// Store provides storage functionality for alerts.
type Store interface {
	GetLogs() ([]string, error)
	CountLogs() (count int, err error)
	GetLogsRange() (start int64, end int64, err error)
	AddLog(log string) error
	RemoveLogs(from, to int64) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) Store {
	bolthelper.RegisterBucket(db, logsBucket)
	return &storeImpl{
		DB: db,
	}
}
