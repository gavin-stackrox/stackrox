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
	baseTable  = "singlekey"
	countStmt  = "SELECT COUNT(*) FROM singlekey"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM singlekey WHERE Key = $1)"

	getStmt     = "SELECT serialized FROM singlekey WHERE Key = $1"
	deleteStmt  = "DELETE FROM singlekey WHERE Key = $1"
	walkStmt    = "SELECT serialized FROM singlekey"
	getIDsStmt  = "SELECT Key FROM singlekey"
	getManyStmt = "SELECT serialized FROM singlekey WHERE Key = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM singlekey WHERE Key = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.TestSingleKeyStruct)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, key string) (bool, error)
	Get(ctx context.Context, key string) (*storage.TestSingleKeyStruct, bool, error)
	Upsert(ctx context.Context, obj *storage.TestSingleKeyStruct) error
	UpsertMany(ctx context.Context, objs []*storage.TestSingleKeyStruct) error
	Delete(ctx context.Context, key string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.TestSingleKeyStruct, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.TestSingleKeyStruct) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableSinglekey(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists singlekey (
    Key varchar,
    Name varchar UNIQUE,
    StringSlice text[],
    Bool bool,
    Uint64 integer,
    Int64 integer,
    Float numeric,
    Labels jsonb,
    Timestamp timestamp,
    Enum integer,
    Enums int[],
    Embedded_Embedded varchar,
    Oneofstring varchar,
    Oneofnested_Nested varchar,
    Oneofnested_Nested2_Nested2 varchar,
    Bytess varchar,
    serialized bytea,
    PRIMARY KEY(Key)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists singlekey_Key on singlekey using hash(Key)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableSinglekeyNested(ctx, db)
}

func createTableSinglekeyNested(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists singlekey_Nested (
    singlekey_Key varchar,
    idx integer,
    Nested varchar,
    Nested2_Nested2 varchar,
    PRIMARY KEY(singlekey_Key, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (singlekey_Key) REFERENCES singlekey(Key) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists singlekeyNested_idx on singlekey_Nested using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoSinglekey(ctx context.Context, tx pgx.Tx, obj *storage.TestSingleKeyStruct) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetKey(),
		obj.GetName(),
		obj.GetStringSlice(),
		obj.GetBool(),
		obj.GetUint64(),
		obj.GetInt64(),
		obj.GetFloat(),
		obj.GetLabels(),
		pgutils.NilOrTime(obj.GetTimestamp()),
		obj.GetEnum(),
		obj.GetEnums(),
		obj.GetEmbedded().GetEmbedded(),
		obj.GetOneofstring(),
		obj.GetOneofnested().GetNested(),
		obj.GetOneofnested().GetNested2().GetNested2(),
		obj.GetBytess(),
		serialized,
	}

	finalStr := "INSERT INTO singlekey (Key, Name, StringSlice, Bool, Uint64, Int64, Float, Labels, Timestamp, Enum, Enums, Embedded_Embedded, Oneofstring, Oneofnested_Nested, Oneofnested_Nested2_Nested2, Bytess, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) ON CONFLICT(Key) DO UPDATE SET Key = EXCLUDED.Key, Name = EXCLUDED.Name, StringSlice = EXCLUDED.StringSlice, Bool = EXCLUDED.Bool, Uint64 = EXCLUDED.Uint64, Int64 = EXCLUDED.Int64, Float = EXCLUDED.Float, Labels = EXCLUDED.Labels, Timestamp = EXCLUDED.Timestamp, Enum = EXCLUDED.Enum, Enums = EXCLUDED.Enums, Embedded_Embedded = EXCLUDED.Embedded_Embedded, Oneofstring = EXCLUDED.Oneofstring, Oneofnested_Nested = EXCLUDED.Oneofnested_Nested, Oneofnested_Nested2_Nested2 = EXCLUDED.Oneofnested_Nested2_Nested2, Bytess = EXCLUDED.Bytess, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetNested() {
		if err := insertIntoSinglekeyNested(ctx, tx, child, obj.GetKey(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from singlekey_Nested where singlekey_Key = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetKey(), len(obj.GetNested()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSinglekeyNested(ctx context.Context, tx pgx.Tx, obj *storage.TestSingleKeyStruct_Nested, singlekey_Key string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		singlekey_Key,
		idx,
		obj.GetNested(),
		obj.GetNested2().GetNested2(),
	}

	finalStr := "INSERT INTO singlekey_Nested (singlekey_Key, idx, Nested, Nested2_Nested2) VALUES($1, $2, $3, $4) ON CONFLICT(singlekey_Key, idx) DO UPDATE SET singlekey_Key = EXCLUDED.singlekey_Key, idx = EXCLUDED.idx, Nested = EXCLUDED.Nested, Nested2_Nested2 = EXCLUDED.Nested2_Nested2"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromSinglekey(ctx context.Context, tx pgx.Tx, objs ...*storage.TestSingleKeyStruct) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"key",

		"name",

		"stringslice",

		"bool",

		"uint64",

		"int64",

		"float",

		"labels",

		"timestamp",

		"enum",

		"enums",

		"embedded_embedded",

		"oneofstring",

		"oneofnested_nested",

		"oneofnested_nested2_nested2",

		"bytess",

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

			obj.GetKey(),

			obj.GetName(),

			obj.GetStringSlice(),

			obj.GetBool(),

			obj.GetUint64(),

			obj.GetInt64(),

			obj.GetFloat(),

			obj.GetLabels(),

			pgutils.NilOrTime(obj.GetTimestamp()),

			obj.GetEnum(),

			obj.GetEnums(),

			obj.GetEmbedded().GetEmbedded(),

			obj.GetOneofstring(),

			obj.GetOneofnested().GetNested(),

			obj.GetOneofnested().GetNested2().GetNested2(),

			obj.GetBytess(),

			serialized,
		})

		// Add the id to be deleted.
		deletes = append(deletes, obj.GetKey())

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

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"singlekey"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromSinglekeyNested(ctx, tx, obj.GetKey(), obj.GetNested()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSinglekeyNested(ctx context.Context, tx pgx.Tx, singlekey_Key string, objs ...*storage.TestSingleKeyStruct_Nested) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"singlekey_key",

		"idx",

		"nested",

		"nested2_nested2",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			singlekey_Key,

			idx,

			obj.GetNested(),

			obj.GetNested2().GetNested2(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"singlekey_nested"}, copyCols, pgx.CopyFromRows(inputRows))

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
	createTableSinglekey(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.TestSingleKeyStruct) error {
	conn, release := s.acquireConn(ctx, ops.Get, "TestSingleKeyStruct")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromSinglekey(ctx, tx, objs...); err != nil {
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

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.TestSingleKeyStruct) error {
	conn, release := s.acquireConn(ctx, ops.Get, "TestSingleKeyStruct")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoSinglekey(ctx, tx, obj); err != nil {
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

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.TestSingleKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "TestSingleKeyStruct")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.TestSingleKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "TestSingleKeyStruct")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "TestSingleKeyStruct")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, key string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "TestSingleKeyStruct")

	row := s.db.QueryRow(ctx, existsStmt, key)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, key string) (*storage.TestSingleKeyStruct, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "TestSingleKeyStruct")

	conn, release := s.acquireConn(ctx, ops.Get, "TestSingleKeyStruct")
	defer release()

	row := conn.QueryRow(ctx, getStmt, key)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.TestSingleKeyStruct
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
func (s *storeImpl) Delete(ctx context.Context, key string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "TestSingleKeyStruct")

	conn, release := s.acquireConn(ctx, ops.Remove, "TestSingleKeyStruct")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, key); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.TestSingleKeyStructIDs")

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
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.TestSingleKeyStruct, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "TestSingleKeyStruct")

	conn, release := s.acquireConn(ctx, ops.GetMany, "TestSingleKeyStruct")
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
	resultsByID := make(map[string]*storage.TestSingleKeyStruct)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.TestSingleKeyStruct{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetKey()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.TestSingleKeyStruct, 0, len(resultsByID))
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "TestSingleKeyStruct")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "TestSingleKeyStruct")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.TestSingleKeyStruct) error) error {
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
		var msg storage.TestSingleKeyStruct
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

func dropTableSinglekey(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS singlekey CASCADE")
	dropTableSinglekeyNested(ctx, db)

}

func dropTableSinglekeyNested(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS singlekey_Nested CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableSinglekey(ctx, db)
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
