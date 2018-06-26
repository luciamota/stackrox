package mappings

import (
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/search"
)

// OptionsMap is exposed for e2e test
var OptionsMap = map[string]*v1.SearchField{
	search.Cluster:    search.NewStringField("policy.scope.cluster"),
	search.Namespace:  search.NewStringField("policy.scope.namespace"),
	search.LabelKey:   search.NewStringField("policy.scope.label.key"),
	search.LabelValue: search.NewStringField("policy.scope.label.value"),

	search.PolicyID:    search.NewStringField("policy.id"),
	search.Enforcement: search.NewEnforcementField("policy.enforcement"),
	search.PolicyName:  search.NewStringField("policy.name"),
	search.Description: search.NewStringField("policy.description"),
	search.Category:    search.NewStringField("policy.categories"),
	search.Severity:    search.NewSeverityField("policy.severity"),
}
