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
	baseTable  = "simpleaccessscopes"
	countStmt  = "SELECT COUNT(*) FROM simpleaccessscopes"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM simpleaccessscopes WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM simpleaccessscopes WHERE Id = $1"
	deleteStmt  = "DELETE FROM simpleaccessscopes WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM simpleaccessscopes"
	getIDsStmt  = "SELECT Id FROM simpleaccessscopes"
	getManyStmt = "SELECT serialized FROM simpleaccessscopes WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM simpleaccessscopes WHERE Id = ANY($1::text[])"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.SimpleAccessScope)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.SimpleAccessScope, bool, error)
	Upsert(ctx context.Context, obj *storage.SimpleAccessScope) error
	UpsertMany(ctx context.Context, objs []*storage.SimpleAccessScope) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.SimpleAccessScope, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.SimpleAccessScope) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableSimpleaccessscopes(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes (
    Id varchar,
    Name varchar UNIQUE,
    Description varchar,
    Rules_IncludedClusters text[],
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

	createTableSimpleaccessscopesIncludedNamespaces(ctx, db)
	createTableSimpleaccessscopesClusterLabelSelectors(ctx, db)
	createTableSimpleaccessscopesNamespaceLabelSelectors(ctx, db)
}

func createTableSimpleaccessscopesIncludedNamespaces(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes_IncludedNamespaces (
    simpleaccessscopes_Id varchar,
    idx integer,
    ClusterName varchar,
    NamespaceName varchar,
    PRIMARY KEY(simpleaccessscopes_Id, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (simpleaccessscopes_Id) REFERENCES simpleaccessscopes(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists simpleaccessscopesIncludedNamespaces_idx on simpleaccessscopes_IncludedNamespaces using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableSimpleaccessscopesClusterLabelSelectors(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes_ClusterLabelSelectors (
    simpleaccessscopes_Id varchar,
    idx integer,
    PRIMARY KEY(simpleaccessscopes_Id, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (simpleaccessscopes_Id) REFERENCES simpleaccessscopes(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists simpleaccessscopesClusterLabelSelectors_idx on simpleaccessscopes_ClusterLabelSelectors using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableSimpleaccessscopesClusterLabelSelectorsRequirements(ctx, db)
}

func createTableSimpleaccessscopesClusterLabelSelectorsRequirements(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes_ClusterLabelSelectors_Requirements (
    simpleaccessscopes_Id varchar,
    simpleaccessscopes_ClusterLabelSelectors_idx integer,
    idx integer,
    Key varchar,
    Op integer,
    Values text[],
    PRIMARY KEY(simpleaccessscopes_Id, simpleaccessscopes_ClusterLabelSelectors_idx, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (simpleaccessscopes_Id, simpleaccessscopes_ClusterLabelSelectors_idx) REFERENCES simpleaccessscopes_ClusterLabelSelectors(simpleaccessscopes_Id, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists simpleaccessscopesClusterLabelSelectorsRequirements_idx on simpleaccessscopes_ClusterLabelSelectors_Requirements using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func createTableSimpleaccessscopesNamespaceLabelSelectors(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes_NamespaceLabelSelectors (
    simpleaccessscopes_Id varchar,
    idx integer,
    PRIMARY KEY(simpleaccessscopes_Id, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (simpleaccessscopes_Id) REFERENCES simpleaccessscopes(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists simpleaccessscopesNamespaceLabelSelectors_idx on simpleaccessscopes_NamespaceLabelSelectors using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx, db)
}

func createTableSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists simpleaccessscopes_NamespaceLabelSelectors_Requirements (
    simpleaccessscopes_Id varchar,
    simpleaccessscopes_NamespaceLabelSelectors_idx integer,
    idx integer,
    Key varchar,
    Op integer,
    Values text[],
    PRIMARY KEY(simpleaccessscopes_Id, simpleaccessscopes_NamespaceLabelSelectors_idx, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (simpleaccessscopes_Id, simpleaccessscopes_NamespaceLabelSelectors_idx) REFERENCES simpleaccessscopes_NamespaceLabelSelectors(simpleaccessscopes_Id, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists simpleaccessscopesNamespaceLabelSelectorsRequirements_idx on simpleaccessscopes_NamespaceLabelSelectors_Requirements using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoSimpleaccessscopes(ctx context.Context, tx pgx.Tx, obj *storage.SimpleAccessScope) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		obj.GetDescription(),
		obj.GetRules().GetIncludedClusters(),
		serialized,
	}

	finalStr := "INSERT INTO simpleaccessscopes (Id, Name, Description, Rules_IncludedClusters, serialized) VALUES($1, $2, $3, $4, $5) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, Description = EXCLUDED.Description, Rules_IncludedClusters = EXCLUDED.Rules_IncludedClusters, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetRules().GetIncludedNamespaces() {
		if err := insertIntoSimpleaccessscopesIncludedNamespaces(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from simpleaccessscopes_IncludedNamespaces where simpleaccessscopes_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetRules().GetIncludedNamespaces()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRules().GetClusterLabelSelectors() {
		if err := insertIntoSimpleaccessscopesClusterLabelSelectors(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from simpleaccessscopes_ClusterLabelSelectors where simpleaccessscopes_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetRules().GetClusterLabelSelectors()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRules().GetNamespaceLabelSelectors() {
		if err := insertIntoSimpleaccessscopesNamespaceLabelSelectors(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from simpleaccessscopes_NamespaceLabelSelectors where simpleaccessscopes_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetRules().GetNamespaceLabelSelectors()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSimpleaccessscopesIncludedNamespaces(ctx context.Context, tx pgx.Tx, obj *storage.SimpleAccessScope_Rules_Namespace, simpleaccessscopes_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		simpleaccessscopes_Id,
		idx,
		obj.GetClusterName(),
		obj.GetNamespaceName(),
	}

	finalStr := "INSERT INTO simpleaccessscopes_IncludedNamespaces (simpleaccessscopes_Id, idx, ClusterName, NamespaceName) VALUES($1, $2, $3, $4) ON CONFLICT(simpleaccessscopes_Id, idx) DO UPDATE SET simpleaccessscopes_Id = EXCLUDED.simpleaccessscopes_Id, idx = EXCLUDED.idx, ClusterName = EXCLUDED.ClusterName, NamespaceName = EXCLUDED.NamespaceName"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSimpleaccessscopesClusterLabelSelectors(ctx context.Context, tx pgx.Tx, obj *storage.SetBasedLabelSelector, simpleaccessscopes_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		simpleaccessscopes_Id,
		idx,
	}

	finalStr := "INSERT INTO simpleaccessscopes_ClusterLabelSelectors (simpleaccessscopes_Id, idx) VALUES($1, $2) ON CONFLICT(simpleaccessscopes_Id, idx) DO UPDATE SET simpleaccessscopes_Id = EXCLUDED.simpleaccessscopes_Id, idx = EXCLUDED.idx"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetRequirements() {
		if err := insertIntoSimpleaccessscopesClusterLabelSelectorsRequirements(ctx, tx, child, simpleaccessscopes_Id, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from simpleaccessscopes_ClusterLabelSelectors_Requirements where simpleaccessscopes_Id = $1 AND simpleaccessscopes_ClusterLabelSelectors_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, simpleaccessscopes_Id, idx, len(obj.GetRequirements()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSimpleaccessscopesClusterLabelSelectorsRequirements(ctx context.Context, tx pgx.Tx, obj *storage.SetBasedLabelSelector_Requirement, simpleaccessscopes_Id string, simpleaccessscopes_ClusterLabelSelectors_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		simpleaccessscopes_Id,
		simpleaccessscopes_ClusterLabelSelectors_idx,
		idx,
		obj.GetKey(),
		obj.GetOp(),
		obj.GetValues(),
	}

	finalStr := "INSERT INTO simpleaccessscopes_ClusterLabelSelectors_Requirements (simpleaccessscopes_Id, simpleaccessscopes_ClusterLabelSelectors_idx, idx, Key, Op, Values) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT(simpleaccessscopes_Id, simpleaccessscopes_ClusterLabelSelectors_idx, idx) DO UPDATE SET simpleaccessscopes_Id = EXCLUDED.simpleaccessscopes_Id, simpleaccessscopes_ClusterLabelSelectors_idx = EXCLUDED.simpleaccessscopes_ClusterLabelSelectors_idx, idx = EXCLUDED.idx, Key = EXCLUDED.Key, Op = EXCLUDED.Op, Values = EXCLUDED.Values"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSimpleaccessscopesNamespaceLabelSelectors(ctx context.Context, tx pgx.Tx, obj *storage.SetBasedLabelSelector, simpleaccessscopes_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		simpleaccessscopes_Id,
		idx,
	}

	finalStr := "INSERT INTO simpleaccessscopes_NamespaceLabelSelectors (simpleaccessscopes_Id, idx) VALUES($1, $2) ON CONFLICT(simpleaccessscopes_Id, idx) DO UPDATE SET simpleaccessscopes_Id = EXCLUDED.simpleaccessscopes_Id, idx = EXCLUDED.idx"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetRequirements() {
		if err := insertIntoSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx, tx, child, simpleaccessscopes_Id, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from simpleaccessscopes_NamespaceLabelSelectors_Requirements where simpleaccessscopes_Id = $1 AND simpleaccessscopes_NamespaceLabelSelectors_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, simpleaccessscopes_Id, idx, len(obj.GetRequirements()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx context.Context, tx pgx.Tx, obj *storage.SetBasedLabelSelector_Requirement, simpleaccessscopes_Id string, simpleaccessscopes_NamespaceLabelSelectors_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		simpleaccessscopes_Id,
		simpleaccessscopes_NamespaceLabelSelectors_idx,
		idx,
		obj.GetKey(),
		obj.GetOp(),
		obj.GetValues(),
	}

	finalStr := "INSERT INTO simpleaccessscopes_NamespaceLabelSelectors_Requirements (simpleaccessscopes_Id, simpleaccessscopes_NamespaceLabelSelectors_idx, idx, Key, Op, Values) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT(simpleaccessscopes_Id, simpleaccessscopes_NamespaceLabelSelectors_idx, idx) DO UPDATE SET simpleaccessscopes_Id = EXCLUDED.simpleaccessscopes_Id, simpleaccessscopes_NamespaceLabelSelectors_idx = EXCLUDED.simpleaccessscopes_NamespaceLabelSelectors_idx, idx = EXCLUDED.idx, Key = EXCLUDED.Key, Op = EXCLUDED.Op, Values = EXCLUDED.Values"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromSimpleaccessscopes(ctx context.Context, tx pgx.Tx, objs ...*storage.SimpleAccessScope) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"name",

		"description",

		"rules_includedclusters",

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

			obj.GetDescription(),

			obj.GetRules().GetIncludedClusters(),

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

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromSimpleaccessscopesIncludedNamespaces(ctx, tx, obj.GetId(), obj.GetRules().GetIncludedNamespaces()...); err != nil {
			return err
		}
		if err = s.copyFromSimpleaccessscopesClusterLabelSelectors(ctx, tx, obj.GetId(), obj.GetRules().GetClusterLabelSelectors()...); err != nil {
			return err
		}
		if err = s.copyFromSimpleaccessscopesNamespaceLabelSelectors(ctx, tx, obj.GetId(), obj.GetRules().GetNamespaceLabelSelectors()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSimpleaccessscopesIncludedNamespaces(ctx context.Context, tx pgx.Tx, simpleaccessscopes_Id string, objs ...*storage.SimpleAccessScope_Rules_Namespace) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"simpleaccessscopes_id",

		"idx",

		"clustername",

		"namespacename",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			simpleaccessscopes_Id,

			idx,

			obj.GetClusterName(),

			obj.GetNamespaceName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes_includednamespaces"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFromSimpleaccessscopesClusterLabelSelectors(ctx context.Context, tx pgx.Tx, simpleaccessscopes_Id string, objs ...*storage.SetBasedLabelSelector) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"simpleaccessscopes_id",

		"idx",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			simpleaccessscopes_Id,

			idx,
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes_clusterlabelselectors"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {

		if err = s.copyFromSimpleaccessscopesClusterLabelSelectorsRequirements(ctx, tx, simpleaccessscopes_Id, idx, obj.GetRequirements()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSimpleaccessscopesClusterLabelSelectorsRequirements(ctx context.Context, tx pgx.Tx, simpleaccessscopes_Id string, simpleaccessscopes_ClusterLabelSelectors_idx int, objs ...*storage.SetBasedLabelSelector_Requirement) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"simpleaccessscopes_id",

		"simpleaccessscopes_clusterlabelselectors_idx",

		"idx",

		"key",

		"op",

		"values",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			simpleaccessscopes_Id,

			simpleaccessscopes_ClusterLabelSelectors_idx,

			idx,

			obj.GetKey(),

			obj.GetOp(),

			obj.GetValues(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes_clusterlabelselectors_requirements"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFromSimpleaccessscopesNamespaceLabelSelectors(ctx context.Context, tx pgx.Tx, simpleaccessscopes_Id string, objs ...*storage.SetBasedLabelSelector) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"simpleaccessscopes_id",

		"idx",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			simpleaccessscopes_Id,

			idx,
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes_namespacelabelselectors"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {

		if err = s.copyFromSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx, tx, simpleaccessscopes_Id, idx, obj.GetRequirements()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx context.Context, tx pgx.Tx, simpleaccessscopes_Id string, simpleaccessscopes_NamespaceLabelSelectors_idx int, objs ...*storage.SetBasedLabelSelector_Requirement) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"simpleaccessscopes_id",

		"simpleaccessscopes_namespacelabelselectors_idx",

		"idx",

		"key",

		"op",

		"values",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			simpleaccessscopes_Id,

			simpleaccessscopes_NamespaceLabelSelectors_idx,

			idx,

			obj.GetKey(),

			obj.GetOp(),

			obj.GetValues(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"simpleaccessscopes_namespacelabelselectors_requirements"}, copyCols, pgx.CopyFromRows(inputRows))

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
	createTableSimpleaccessscopes(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.SimpleAccessScope) error {
	conn, release := s.acquireConn(ctx, ops.Get, "SimpleAccessScope")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromSimpleaccessscopes(ctx, tx, objs...); err != nil {
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

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.SimpleAccessScope) error {
	conn, release := s.acquireConn(ctx, ops.Get, "SimpleAccessScope")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoSimpleaccessscopes(ctx, tx, obj); err != nil {
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

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.SimpleAccessScope) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "SimpleAccessScope")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.SimpleAccessScope) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "SimpleAccessScope")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "SimpleAccessScope")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "SimpleAccessScope")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.SimpleAccessScope, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "SimpleAccessScope")

	conn, release := s.acquireConn(ctx, ops.Get, "SimpleAccessScope")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.SimpleAccessScope
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "SimpleAccessScope")

	conn, release := s.acquireConn(ctx, ops.Remove, "SimpleAccessScope")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.SimpleAccessScopeIDs")

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
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.SimpleAccessScope, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "SimpleAccessScope")

	conn, release := s.acquireConn(ctx, ops.GetMany, "SimpleAccessScope")
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
	resultsByID := make(map[string]*storage.SimpleAccessScope)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.SimpleAccessScope{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.SimpleAccessScope, 0, len(resultsByID))
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "SimpleAccessScope")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "SimpleAccessScope")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.SimpleAccessScope) error) error {
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
		var msg storage.SimpleAccessScope
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

func dropTableSimpleaccessscopes(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes CASCADE")
	dropTableSimpleaccessscopesIncludedNamespaces(ctx, db)
	dropTableSimpleaccessscopesClusterLabelSelectors(ctx, db)
	dropTableSimpleaccessscopesNamespaceLabelSelectors(ctx, db)

}

func dropTableSimpleaccessscopesIncludedNamespaces(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes_IncludedNamespaces CASCADE")

}

func dropTableSimpleaccessscopesClusterLabelSelectors(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes_ClusterLabelSelectors CASCADE")
	dropTableSimpleaccessscopesClusterLabelSelectorsRequirements(ctx, db)

}

func dropTableSimpleaccessscopesClusterLabelSelectorsRequirements(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes_ClusterLabelSelectors_Requirements CASCADE")

}

func dropTableSimpleaccessscopesNamespaceLabelSelectors(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes_NamespaceLabelSelectors CASCADE")
	dropTableSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx, db)

}

func dropTableSimpleaccessscopesNamespaceLabelSelectorsRequirements(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS simpleaccessscopes_NamespaceLabelSelectors_Requirements CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableSimpleaccessscopes(ctx, db)
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
