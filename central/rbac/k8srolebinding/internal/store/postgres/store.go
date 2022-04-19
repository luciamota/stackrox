// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/metrics"
	pkgSchema "github.com/stackrox/rox/central/postgres/schema"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac/effectiveaccessscope"
)

const (
	baseTable  = "rolebindings"
	countStmt  = "SELECT COUNT(*) FROM rolebindings"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM rolebindings WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM rolebindings WHERE Id = $1"
	deleteStmt  = "DELETE FROM rolebindings WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM rolebindings"
	getIDsStmt  = "SELECT Id FROM rolebindings"
	getManyStmt = "SELECT serialized FROM rolebindings WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM rolebindings WHERE Id = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log    = logging.LoggerForModule()
	schema = pkgSchema.RolebindingsSchema
)

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.K8SRoleBinding, bool, error)
	Upsert(ctx context.Context, obj *storage.K8SRoleBinding) error
	UpsertMany(ctx context.Context, objs []*storage.K8SRoleBinding) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.K8SRoleBinding, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.K8SRoleBinding) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	pgutils.CreateTable(ctx, db, pkgSchema.CreateTableRolebindingsStmt)

	return &storeImpl{
		db: db,
	}
}

func insertIntoRolebindings(ctx context.Context, tx pgx.Tx, obj *storage.K8SRoleBinding) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		obj.GetNamespace(),
		obj.GetClusterId(),
		obj.GetClusterName(),
		obj.GetClusterRole(),
		obj.GetLabels(),
		obj.GetAnnotations(),
		obj.GetRoleId(),
		serialized,
	}

	finalStr := "INSERT INTO rolebindings (Id, Name, Namespace, ClusterId, ClusterName, ClusterRole, Labels, Annotations, RoleId, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, Namespace = EXCLUDED.Namespace, ClusterId = EXCLUDED.ClusterId, ClusterName = EXCLUDED.ClusterName, ClusterRole = EXCLUDED.ClusterRole, Labels = EXCLUDED.Labels, Annotations = EXCLUDED.Annotations, RoleId = EXCLUDED.RoleId, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetSubjects() {
		if err := insertIntoRolebindingsSubjects(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from rolebindings_Subjects where rolebindings_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetSubjects()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoRolebindingsSubjects(ctx context.Context, tx pgx.Tx, obj *storage.Subject, rolebindings_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		rolebindings_Id,
		idx,
		obj.GetKind(),
		obj.GetName(),
	}

	finalStr := "INSERT INTO rolebindings_Subjects (rolebindings_Id, idx, Kind, Name) VALUES($1, $2, $3, $4) ON CONFLICT(rolebindings_Id, idx) DO UPDATE SET rolebindings_Id = EXCLUDED.rolebindings_Id, idx = EXCLUDED.idx, Kind = EXCLUDED.Kind, Name = EXCLUDED.Name"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromRolebindings(ctx context.Context, tx pgx.Tx, objs ...*storage.K8SRoleBinding) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"name",

		"namespace",

		"clusterid",

		"clustername",

		"clusterrole",

		"labels",

		"annotations",

		"roleid",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			obj.GetId(),

			obj.GetName(),

			obj.GetNamespace(),

			obj.GetClusterId(),

			obj.GetClusterName(),

			obj.GetClusterRole(),

			obj.GetLabels(),

			obj.GetAnnotations(),

			obj.GetRoleId(),

			serialized,
		})

		// Add the id to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.Exec(ctx, deleteManyStmt, deletes)
			if err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"rolebindings"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromRolebindingsSubjects(ctx, tx, obj.GetId(), obj.GetSubjects()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromRolebindingsSubjects(ctx context.Context, tx pgx.Tx, rolebindings_Id string, objs ...*storage.Subject) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"rolebindings_id",

		"idx",

		"kind",

		"name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			rolebindings_Id,

			idx,

			obj.GetKind(),

			obj.GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"rolebindings_subjects"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.K8SRoleBinding) error {
	conn, release, err := s.acquireConn(ctx, ops.Get, "K8SRoleBinding")
	if err != nil {
		return err
	}
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromRolebindings(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.K8SRoleBinding) error {
	conn, release, err := s.acquireConn(ctx, ops.Get, "K8SRoleBinding")
	if err != nil {
		return err
	}
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoRolebindings(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.K8SRoleBinding) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "K8SRoleBinding")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.K8SRoleBinding) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "K8SRoleBinding")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "K8SRoleBinding")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "K8SRoleBinding")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.K8SRoleBinding, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "K8SRoleBinding")

	conn, release, err := s.acquireConn(ctx, ops.Get, "K8SRoleBinding")
	if err != nil {
		return nil, false, err
	}
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.K8SRoleBinding
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func(), error) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Release, nil
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "K8SRoleBinding")

	conn, release, err := s.acquireConn(ctx, ops.Remove, "K8SRoleBinding")
	if err != nil {
		return err
	}
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.K8SRoleBindingIDs")

	rows, err := s.db.Query(ctx, getIDsStmt)
	if err != nil {
		return nil, pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.K8SRoleBinding, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "K8SRoleBinding")

	conn, release, err := s.acquireConn(ctx, ops.GetMany, "K8SRoleBinding")
	if err != nil {
		return nil, nil, err
	}
	defer release()

	rows, err := conn.Query(ctx, getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	resultsByID := make(map[string]*storage.K8SRoleBinding)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.K8SRoleBinding{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.K8SRoleBinding, 0, len(resultsByID))
	for i, id := range ids {
		if result, ok := resultsByID[id]; !ok {
			missingIndices = append(missingIndices, i)
		} else {
			elems = append(elems, result)
		}
	}
	return elems, missingIndices, nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ctx context.Context, ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "K8SRoleBinding")

	conn, release, err := s.acquireConn(ctx, ops.RemoveMany, "K8SRoleBinding")
	if err != nil {
		return err
	}
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.K8SRoleBinding) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.K8SRoleBinding
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

func isInScope(obj *storage.K8SRoleBinding, eas *effectiveaccessscope.ScopeTree) bool {
	if eas.State == effectiveaccessscope.Included {
		return true
	}
	if eas.State == effectiveaccessscope.Excluded {
		return false
	}
	clusterId := obj.GetClusterId()
	cluster := eas.GetClusterByID(clusterId)
	if cluster.State == effectiveaccessscope.Included {
		return true
	}
	if cluster.State == effectiveaccessscope.Excluded {
		return false
	}
	namespaceName := obj.GetNamespace()
	return cluster.Namespaces[namespaceName].State == effectiveaccessscope.Included
}

//// Used for testing

func dropTableRolebindings(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS rolebindings CASCADE")
	dropTableRolebindingsSubjects(ctx, db)

}

func dropTableRolebindingsSubjects(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS rolebindings_Subjects CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableRolebindings(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
