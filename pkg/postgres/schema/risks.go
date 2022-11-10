// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableRisksStmt holds the create statement for table `risks`.
	CreateTableRisksStmt = &postgres.CreateStmts{
		GormModel: (*Risks)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// RisksSchema is the go schema for table `risks`.
	RisksSchema = func() *walker.Schema {
		schema := GetSchemaForTable("risks")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Risk)(nil)), "risks")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_RISKS, "risk", (*storage.Risk)(nil)))
		RegisterTable(schema, CreateTableRisksStmt)
		return schema
	}()
)

const (
	RisksTableName = "risks"
)

// Risks holds the Gorm model for Postgres table `risks`.
type Risks struct {
	Id               string                  `gorm:"column:id;type:varchar;primaryKey"`
	SubjectNamespace string                  `gorm:"column:subject_namespace;type:varchar;index:risks_sac_filter,type:btree"`
	SubjectClusterId string                  `gorm:"column:subject_clusterid;type:uuid;index:risks_sac_filter,type:btree"`
	SubjectType      storage.RiskSubjectType `gorm:"column:subject_type;type:integer"`
	Score            float32                 `gorm:"column:score;type:numeric"`
	Serialized       []byte                  `gorm:"column:serialized;type:bytea"`
}
