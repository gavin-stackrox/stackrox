// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac/resources"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "test_structs"
	storeName = "TestStruct"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.TestStructsSchema
	targetResource = resources.Namespace
)

type storeType = storage.TestStruct

// Store is the interface to interact with the storage for storage.TestStruct
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, key1 string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) ([]string, error)
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, key1 string) (bool, error)

	Get(ctx context.Context, key1 string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStore[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoTestStructs,
		copyFromTestStructs,
		metricsSetAcquireDBConnDuration,
		metricsSetPostgresOperationDurationTime,
		pgSearch.GloballyScopedUpsertChecker[storeType, *storeType](targetResource),
		targetResource,
	)
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetKey1()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoTestStructs(batch *pgx.Batch, obj *storage.TestStruct) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetKey1(),
		obj.GetKey2(),
		obj.GetStringSlice(),
		obj.GetBool(),
		obj.GetUint64(),
		obj.GetInt64(),
		obj.GetFloat(),
		pgutils.EmptyOrMap(obj.GetLabels()),
		pgutils.NilOrTime(obj.GetTimestamp()),
		obj.GetEnum(),
		obj.GetEnums(),
		obj.GetString_(),
		obj.GetInt32Slice(),
		obj.GetOneofnested().GetNested(),
		serialized,
	}

	finalStr := "INSERT INTO test_structs (Key1, Key2, StringSlice, Bool, Uint64, Int64, Float, Labels, Timestamp, Enum, Enums, String_, Int32Slice, Oneofnested_Nested, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT(Key1) DO UPDATE SET Key1 = EXCLUDED.Key1, Key2 = EXCLUDED.Key2, StringSlice = EXCLUDED.StringSlice, Bool = EXCLUDED.Bool, Uint64 = EXCLUDED.Uint64, Int64 = EXCLUDED.Int64, Float = EXCLUDED.Float, Labels = EXCLUDED.Labels, Timestamp = EXCLUDED.Timestamp, Enum = EXCLUDED.Enum, Enums = EXCLUDED.Enums, String_ = EXCLUDED.String_, Int32Slice = EXCLUDED.Int32Slice, Oneofnested_Nested = EXCLUDED.Oneofnested_Nested, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetNested() {
		if err := insertIntoTestStructsNesteds(batch, child, obj.GetKey1(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from test_structs_nesteds where test_structs_Key1 = $1 AND idx >= $2"
	batch.Queue(query, obj.GetKey1(), len(obj.GetNested()))
	return nil
}

func insertIntoTestStructsNesteds(batch *pgx.Batch, obj *storage.TestStruct_Nested, testStructKey1 string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		testStructKey1,
		idx,
		obj.GetNested(),
		obj.GetIsNested(),
		obj.GetInt64(),
		obj.GetNested2().GetNested2(),
		obj.GetNested2().GetIsNested(),
		obj.GetNested2().GetInt64(),
	}

	finalStr := "INSERT INTO test_structs_nesteds (test_structs_Key1, idx, Nested, IsNested, Int64, Nested2_Nested2, Nested2_IsNested, Nested2_Int64) VALUES($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT(test_structs_Key1, idx) DO UPDATE SET test_structs_Key1 = EXCLUDED.test_structs_Key1, idx = EXCLUDED.idx, Nested = EXCLUDED.Nested, IsNested = EXCLUDED.IsNested, Int64 = EXCLUDED.Int64, Nested2_Nested2 = EXCLUDED.Nested2_Nested2, Nested2_IsNested = EXCLUDED.Nested2_IsNested, Nested2_Int64 = EXCLUDED.Nested2_Int64"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromTestStructs(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.TestStruct) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	deletes := make([]string, 0, batchSize)

	copyCols := []string{
		"key1",
		"key2",
		"stringslice",
		"bool",
		"uint64",
		"int64",
		"float",
		"labels",
		"timestamp",
		"enum",
		"enums",
		"string_",
		"int32slice",
		"oneofnested_nested",
		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{
			obj.GetKey1(),
			obj.GetKey2(),
			obj.GetStringSlice(),
			obj.GetBool(),
			obj.GetUint64(),
			obj.GetInt64(),
			obj.GetFloat(),
			pgutils.EmptyOrMap(obj.GetLabels()),
			pgutils.NilOrTime(obj.GetTimestamp()),
			obj.GetEnum(),
			obj.GetEnums(),
			obj.GetString_(),
			obj.GetInt32Slice(),
			obj.GetOneofnested().GetNested(),
			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetKey1())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = deletes[:0]

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"test_structs"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromTestStructsNesteds(ctx, s, tx, obj.GetKey1(), obj.GetNested()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromTestStructsNesteds(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, testStructKey1 string, objs ...*storage.TestStruct_Nested) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"test_structs_key1",
		"idx",
		"nested",
		"isnested",
		"int64",
		"nested2_nested2",
		"nested2_isnested",
		"nested2_int64",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			testStructKey1,
			idx,
			obj.GetNested(),
			obj.GetIsNested(),
			obj.GetInt64(),
			obj.GetNested2().GetNested2(),
			obj.GetNested2().GetIsNested(),
			obj.GetNested2().GetInt64(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"test_structs_nesteds"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return nil
}

// endregion Helper functions

// region Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableTestStructs(ctx, db)
}

func dropTableTestStructs(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_structs CASCADE")
	dropTableTestStructsNesteds(ctx, db)

}

func dropTableTestStructsNesteds(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_structs_nesteds CASCADE")

}

// endregion Used for testing
