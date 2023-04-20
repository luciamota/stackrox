package inputtypes

import (
	"strings"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/search/postgres/aggregatefunc"
)

// SortOption is the sort option input type.
type SortOption struct {
	Field       *string
	AggregateBy *AggregateBy
	Reversed    *bool
}

// AggregateBy is the sort option input type.
type AggregateBy struct {
	AggregateFunc *string
	Distinct      *bool
}

// AsV1SortOption converts the sort option to proto.
func (s *SortOption) AsV1SortOption() *v1.SortOption {
	if s == nil {
		return nil
	}
	return &v1.SortOption{
		Field: func() string {
			if s.Field == nil {
				return ""
			}
			return *s.Field
		}(),
		AggregateBy: s.AggregateBy.AsV1AggregateBy(),
		Reversed: func() bool {
			if s.Reversed == nil {
				return false
			}
			return *s.Reversed
		}(),
	}
}

// AsV1AggregateBy converts the aggregation to proto.
func (a *AggregateBy) AsV1AggregateBy() *v1.AggregateBy {
	if a == nil {
		return nil
	}
	return &v1.AggregateBy{
		AggrFunc: aggregatefunc.GetAggrFunc(strings.ToLower(*a.AggregateFunc)).Proto(),
		Distinct: *a.Distinct,
	}
}
