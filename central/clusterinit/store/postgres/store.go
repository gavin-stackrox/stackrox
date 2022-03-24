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
	baseTable  = "clusterinitbundles"
	countStmt  = "SELECT COUNT(*) FROM clusterinitbundles"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM clusterinitbundles WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM clusterinitbundles WHERE Id = $1"
	deleteStmt  = "DELETE FROM clusterinitbundles WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM clusterinitbundles"
	getIDsStmt  = "SELECT Id FROM clusterinitbundles"
	getManyStmt = "SELECT serialized FROM clusterinitbundles WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM clusterinitbundles WHERE Id = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.InitBundleMeta)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.InitBundleMeta, bool, error)
	Upsert(ctx context.Context, obj *storage.InitBundleMeta) error
	UpsertMany(ctx context.Context, objs []*storage.InitBundleMeta) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.InitBundleMeta, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.InitBundleMeta) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableClusterinitbundles(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists clusterinitbundles (
    Id varchar,
    Name varchar,
    CreatedAt timestamp,
    CreatedBy_Id varchar,
    CreatedBy_AuthProviderId varchar,
    IsRevoked bool,
    ExpiresAt timestamp,
    serialized bytea,
    PRIMARY KEY(Id)
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

	createTableClusterinitbundlesAttributes(ctx, db)
}

func createTableClusterinitbundlesAttributes(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists clusterinitbundles_Attributes (
    clusterinitbundles_Id varchar,
    idx integer,
    Key varchar,
    Value varchar,
    PRIMARY KEY(clusterinitbundles_Id, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (clusterinitbundles_Id) REFERENCES clusterinitbundles(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists clusterinitbundlesAttributes_idx on clusterinitbundles_Attributes using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoClusterinitbundles(ctx context.Context, tx pgx.Tx, obj *storage.InitBundleMeta) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		pgutils.NilOrTime(obj.GetCreatedAt()),
		obj.GetCreatedBy().GetId(),
		obj.GetCreatedBy().GetAuthProviderId(),
		obj.GetIsRevoked(),
		pgutils.NilOrTime(obj.GetExpiresAt()),
		serialized,
	}

	finalStr := "INSERT INTO clusterinitbundles (Id, Name, CreatedAt, CreatedBy_Id, CreatedBy_AuthProviderId, IsRevoked, ExpiresAt, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, CreatedAt = EXCLUDED.CreatedAt, CreatedBy_Id = EXCLUDED.CreatedBy_Id, CreatedBy_AuthProviderId = EXCLUDED.CreatedBy_AuthProviderId, IsRevoked = EXCLUDED.IsRevoked, ExpiresAt = EXCLUDED.ExpiresAt, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetCreatedBy().GetAttributes() {
		if err := insertIntoClusterinitbundlesAttributes(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from clusterinitbundles_Attributes where clusterinitbundles_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetCreatedBy().GetAttributes()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoClusterinitbundlesAttributes(ctx context.Context, tx pgx.Tx, obj *storage.UserAttribute, clusterinitbundles_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		clusterinitbundles_Id,
		idx,
		obj.GetKey(),
		obj.GetValue(),
	}

	finalStr := "INSERT INTO clusterinitbundles_Attributes (clusterinitbundles_Id, idx, Key, Value) VALUES($1, $2, $3, $4) ON CONFLICT(clusterinitbundles_Id, idx) DO UPDATE SET clusterinitbundles_Id = EXCLUDED.clusterinitbundles_Id, idx = EXCLUDED.idx, Key = EXCLUDED.Key, Value = EXCLUDED.Value"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromClusterinitbundles(ctx context.Context, tx pgx.Tx, objs ...*storage.InitBundleMeta) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"name",

		"createdat",

		"createdby_id",

		"createdby_authproviderid",

		"isrevoked",

		"expiresat",

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

			pgutils.NilOrTime(obj.GetCreatedAt()),

			obj.GetCreatedBy().GetId(),

			obj.GetCreatedBy().GetAuthProviderId(),

			obj.GetIsRevoked(),

			pgutils.NilOrTime(obj.GetExpiresAt()),

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

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"clusterinitbundles"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromClusterinitbundlesAttributes(ctx, tx, obj.GetId(), obj.GetCreatedBy().GetAttributes()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromClusterinitbundlesAttributes(ctx context.Context, tx pgx.Tx, clusterinitbundles_Id string, objs ...*storage.UserAttribute) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"clusterinitbundles_id",

		"idx",

		"key",

		"value",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			clusterinitbundles_Id,

			idx,

			obj.GetKey(),

			obj.GetValue(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"clusterinitbundles_attributes"}, copyCols, pgx.CopyFromRows(inputRows))

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
	createTableClusterinitbundles(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.InitBundleMeta) error {
	conn, release := s.acquireConn(ctx, ops.Get, "InitBundleMeta")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromClusterinitbundles(ctx, tx, objs...); err != nil {
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

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.InitBundleMeta) error {
	conn, release := s.acquireConn(ctx, ops.Get, "InitBundleMeta")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoClusterinitbundles(ctx, tx, obj); err != nil {
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

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.InitBundleMeta) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "InitBundleMeta")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.InitBundleMeta) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "InitBundleMeta")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "InitBundleMeta")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "InitBundleMeta")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.InitBundleMeta, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "InitBundleMeta")

	conn, release := s.acquireConn(ctx, ops.Get, "InitBundleMeta")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.InitBundleMeta
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
func (s *storeImpl) Delete(ctx context.Context, id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "InitBundleMeta")

	conn, release := s.acquireConn(ctx, ops.Remove, "InitBundleMeta")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.InitBundleMetaIDs")

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
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.InitBundleMeta, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "InitBundleMeta")

	conn, release := s.acquireConn(ctx, ops.GetMany, "InitBundleMeta")
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
	resultsByID := make(map[string]*storage.InitBundleMeta)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.InitBundleMeta{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.InitBundleMeta, 0, len(resultsByID))
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "InitBundleMeta")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "InitBundleMeta")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.InitBundleMeta) error) error {
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
		var msg storage.InitBundleMeta
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

func dropTableClusterinitbundles(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS clusterinitbundles CASCADE")
	dropTableClusterinitbundlesAttributes(ctx, db)

}

func dropTableClusterinitbundlesAttributes(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS clusterinitbundles_Attributes CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableClusterinitbundles(ctx, db)
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
