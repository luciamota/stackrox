// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableClusterHealthStatusStmt holds the create statement for table `cluster_health_status`.
	CreateTableClusterHealthStatusStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists cluster_health_status (
                   Id varchar,
                   SensorHealthStatus integer,
                   CollectorHealthStatus integer,
                   OverallHealthStatus integer,
                   AdmissionControlHealthStatus integer,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// ClusterHealthStatusSchema is the go schema for table `cluster_health_status`.
	ClusterHealthStatusSchema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("cluster_health_status")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ClusterHealthStatus)(nil)), "cluster_health_status")
		globaldb.RegisterTable(schema)
		return schema
	}()
)
