// Code generated by make bootstrap_migration generator. DO NOT EDIT.
package m206tom207

import (
	"context"

	"github.com/stackrox/rox/pkg/uuid"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/migrator/migrations/m_206_to_m_207_add_default_policy_edge/schema"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/types"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

const (
	startSeqNum = 206
)

var (
	migration = types.Migration{
		StartingSeqNum: startSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startSeqNum + 1)},
		Run: func(databases *types.Databases) error {
			err := updatePolicyCategoryEdge(databases.GormDB)
			if err != nil {
				return errors.Wrap(err, "inserting policy categpry edge")
			}
			return nil
		},
	}
)

func updatePolicyCategoryEdge(db *gorm.DB) error {
	//This is a map of policy id and category id which belong to the policy
	policy_category_map := map[string][]string {
		"fb8f8732-c31d-496b-8fb1-d5abe6056e27": {"f732f1a5-1515-4e9e-9179-3ab2aefe9ad9","99cfb323-c9d3-4e0c-af64-4d0101659866"},
		"ed8c7957-14de-40bc-aeab-d27ceeecfa7b": {"99cfb323-c9d3-4e0c-af64-4d0101659866","9d924f5d-6679-4449-8154-795449c8e754"},
		"6226d4ad-7619-4a0b-a160-46373cfcee66": {"d2bbe19e-3009-4a0e-a701-a0b621b319a0"},
		"dce17697-1b72-49d2-b18a-05d893cd9368": {"d2bbe19e-3009-4a0e-a701-a0b621b319a0"},
		"a9b9ecf7-9707-4e32-8b62-d03018ed454f"  :{"99cfb323-c9d3-4e0c-af64-4d0101659866"},
	}
	for policy,category := range policy_category_map{
		err := add_policy_category_edge(db,policy,category)
		if err!=nil{
			return err
		}
	}
	return nil
}

func upsertPolicyCategoryEdge(db *gorm.DB,ctx context.Context, edge *storage.PolicyCategoryEdge) error {
	dbEdge, err := schema.ConvertPolicyCategoryEdgeFromProto(edge)
	if err != nil {
		return err
	}
	result := db.WithContext(ctx).Table(schema.PolicyCategoryEdgesTableName).Save(dbEdge)
	if result.RowsAffected != 1 {
		return errors.Errorf("failed to save edge for policy id %s, category id %s: %q",
			edge.GetPolicyId(), edge.GetCategoryId(), result.Error)
	}
	return result.Error
}

func add_policy_category_edge(db *gorm.DB,policyID string, categories []string) error {
	ctx := sac.WithAllAccess(context.Background())
	for _,categoryID := range categories {
		edge := &storage.PolicyCategoryEdge{
			Id:         uuid.NewV4().String(),
			PolicyId:   policyID,
			CategoryId: categoryID,
		}
		err := upsertPolicyCategoryEdge(db, ctx, edge)
		if err != nil{
			return err
		}
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
