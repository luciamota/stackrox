// Code generated by pg-bindings generator. DO NOT EDIT.
package conversion

import (
	"github.com/lib/pq"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations/m_208_to_m_209_policy_updates_for_4_6/schema"
)

// ConvertPolicyFromProto converts a `*storage.Policy` to Gorm model
func ConvertPolicyFromProto(obj *storage.Policy) (*schema.Policies, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &schema.Policies{
		ID:                 obj.GetId(),
		Name:               obj.GetName(),
		Description:        obj.GetDescription(),
		Disabled:           obj.GetDisabled(),
		Categories:         pq.Array(obj.GetCategories()).(*pq.StringArray),
		LifecycleStages:    pq.Array(convertEnumSliceToIntArray(obj.GetLifecycleStages())).(*pq.Int32Array),
		Severity:           obj.GetSeverity(),
		EnforcementActions: pq.Array(convertEnumSliceToIntArray(obj.GetEnforcementActions())).(*pq.Int32Array),
		LastUpdated:        nilOrTime(obj.GetLastUpdated()),
		SORTName:           obj.GetSORTName(),
		SORTLifecycleStage: obj.GetSORTLifecycleStage(),
		SORTEnforcement:    obj.GetSORTEnforcement(),
		Serialized:         serialized,
	}
	return model, nil
}

// ConvertPolicyToProto converts Gorm model `Policies` to its protobuf type object
func ConvertPolicyToProto(m *schema.Policies) (*storage.Policy, error) {
	var msg storage.Policy
	if err := msg.UnmarshalVTUnsafe(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
