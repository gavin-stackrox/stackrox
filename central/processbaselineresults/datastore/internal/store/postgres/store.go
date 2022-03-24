// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

const (
	baseTable  = "processwhitelistresults"
	countStmt  = "SELECT COUNT(*) FROM processwhitelistresults"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM processwhitelistresults WHERE DeploymentId = $1)"

	getStmt     = "SELECT serialized FROM processwhitelistresults WHERE DeploymentId = $1"
	deleteStmt  = "DELETE FROM processwhitelistresults WHERE DeploymentId = $1"
	walkStmt    = "SELECT serialized FROM processwhitelistresults"
	getIDsStmt  = "SELECT DeploymentId FROM processwhitelistresults"
	getManyStmt = "SELECT serialized FROM processwhitelistresults WHERE DeploymentId = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM processwhitelistresults WHERE DeploymentId = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.ProcessBaselineResults)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, deploymentId string) (bool, error)
	Get(ctx context.Context, deploymentId string) (*storage.ProcessBaselineResults, bool, error)
	Upsert(ctx context.Context, obj *storage.ProcessBaselineResults) error
	UpsertMany(ctx context.Context, objs []*storage.ProcessBaselineResults) error
	Delete(ctx context.Context, deploymentId string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.ProcessBaselineResults, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.ProcessBaselineResults) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableProcesswhitelistresults(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists processwhitelistresults (
    DeploymentId varchar,
    ClusterId varchar,
    Namespace varchar,
    serialized bytea,
    PRIMARY KEY(DeploymentId)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableProcesswhitelistresultsBaselineStatuses(ctx, db)
}

func createTableProcesswhitelistresultsBaselineStatuses(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists processwhitelistresults_BaselineStatuses (
    processwhitelistresults_DeploymentId varchar,
    idx integer,
    ContainerName varchar,
    BaselineStatus integer,
    AnomalousProcessesExecuted bool,
    PRIMARY KEY(processwhitelistresults_DeploymentId, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (processwhitelistresults_DeploymentId) REFERENCES processwhitelistresults(DeploymentId) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists processwhitelistresultsBaselineStatuses_idx on processwhitelistresults_BaselineStatuses using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoProcesswhitelistresults(ctx context.Context, tx pgx.Tx, obj *storage.ProcessBaselineResults) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetDeploymentId(),
		obj.GetClusterId(),
		obj.GetNamespace(),
		serialized,
	}

	finalStr := "INSERT INTO processwhitelistresults (DeploymentId, ClusterId, Namespace, serialized) VALUES($1, $2, $3, $4) ON CONFLICT(DeploymentId) DO UPDATE SET DeploymentId = EXCLUDED.DeploymentId, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetBaselineStatuses() {
		if err := insertIntoProcesswhitelistresultsBaselineStatuses(ctx, tx, child, obj.GetDeploymentId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from processwhitelistresults_BaselineStatuses where processwhitelistresults_DeploymentId = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetDeploymentId(), len(obj.GetBaselineStatuses()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoProcesswhitelistresultsBaselineStatuses(ctx context.Context, tx pgx.Tx, obj *storage.ContainerNameAndBaselineStatus, processwhitelistresults_DeploymentId string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		processwhitelistresults_DeploymentId,
		idx,
		obj.GetContainerName(),
		obj.GetBaselineStatus(),
		obj.GetAnomalousProcessesExecuted(),
	}

	finalStr := "INSERT INTO processwhitelistresults_BaselineStatuses (processwhitelistresults_DeploymentId, idx, ContainerName, BaselineStatus, AnomalousProcessesExecuted) VALUES($1, $2, $3, $4, $5) ON CONFLICT(processwhitelistresults_DeploymentId, idx) DO UPDATE SET processwhitelistresults_DeploymentId = EXCLUDED.processwhitelistresults_DeploymentId, idx = EXCLUDED.idx, ContainerName = EXCLUDED.ContainerName, BaselineStatus = EXCLUDED.BaselineStatus, AnomalousProcessesExecuted = EXCLUDED.AnomalousProcessesExecuted"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromProcesswhitelistresults(ctx context.Context, tx pgx.Tx, objs ...*storage.ProcessBaselineResults) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"deploymentid",

		"clusterid",

		"namespace",

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

			obj.GetDeploymentId(),

			obj.GetClusterId(),

			obj.GetNamespace(),

			serialized,
		})

		// Add the id to be deleted.
		deletes = append(deletes, obj.GetDeploymentId())

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

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"processwhitelistresults"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromProcesswhitelistresultsBaselineStatuses(ctx, tx, obj.GetDeploymentId(), obj.GetBaselineStatuses()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromProcesswhitelistresultsBaselineStatuses(ctx context.Context, tx pgx.Tx, processwhitelistresults_DeploymentId string, objs ...*storage.ContainerNameAndBaselineStatus) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"processwhitelistresults_deploymentid",

		"idx",

		"containername",

		"baselinestatus",

		"anomalousprocessesexecuted",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			processwhitelistresults_DeploymentId,

			idx,

			obj.GetContainerName(),

			obj.GetBaselineStatus(),

			obj.GetAnomalousProcessesExecuted(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"processwhitelistresults_baselinestatuses"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableProcesswhitelistresults(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.ProcessBaselineResults) error {
	conn, release := s.acquireConn(ctx, ops.Get, "ProcessBaselineResults")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromProcesswhitelistresults(ctx, tx, objs...); err != nil {
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

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.ProcessBaselineResults) error {
	conn, release := s.acquireConn(ctx, ops.Get, "ProcessBaselineResults")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoProcesswhitelistresults(ctx, tx, obj); err != nil {
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

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.ProcessBaselineResults) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "ProcessBaselineResults")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.ProcessBaselineResults) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "ProcessBaselineResults")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ProcessBaselineResults")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, deploymentId string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ProcessBaselineResults")

	row := s.db.QueryRow(ctx, existsStmt, deploymentId)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, deploymentId string) (*storage.ProcessBaselineResults, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ProcessBaselineResults")

	conn, release := s.acquireConn(ctx, ops.Get, "ProcessBaselineResults")
	defer release()

	row := conn.QueryRow(ctx, getStmt, deploymentId)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.ProcessBaselineResults
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, deploymentId string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ProcessBaselineResults")

	conn, release := s.acquireConn(ctx, ops.Remove, "ProcessBaselineResults")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, deploymentId); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.ProcessBaselineResultsIDs")

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
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.ProcessBaselineResults, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ProcessBaselineResults")

	conn, release := s.acquireConn(ctx, ops.GetMany, "ProcessBaselineResults")
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
	resultsByID := make(map[string]*storage.ProcessBaselineResults)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.ProcessBaselineResults{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetDeploymentId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.ProcessBaselineResults, 0, len(resultsByID))
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "ProcessBaselineResults")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "ProcessBaselineResults")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.ProcessBaselineResults) error) error {
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
		var msg storage.ProcessBaselineResults
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableProcesswhitelistresults(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS processwhitelistresults CASCADE")
	dropTableProcesswhitelistresultsBaselineStatuses(ctx, db)

}

func dropTableProcesswhitelistresultsBaselineStatuses(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS processwhitelistresults_BaselineStatuses CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableProcesswhitelistresults(ctx, db)
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
