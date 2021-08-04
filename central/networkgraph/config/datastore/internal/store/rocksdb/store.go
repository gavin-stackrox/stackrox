// Code generated by rocksdb-bindings generator. DO NOT EDIT.

package rocksdb

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/db"
	"github.com/stackrox/rox/pkg/rocksdb"
	generic "github.com/stackrox/rox/pkg/rocksdb/crud"
)

var (
	log = logging.LoggerForModule()

	bucket = []byte("networkgraphconfig")
)

type Store interface {
	Count() (int, error)
	Exists(id string) (bool, error)
	GetIDs() ([]string, error)
	Get(id string) (*storage.NetworkGraphConfig, bool, error)
	GetMany(ids []string) ([]*storage.NetworkGraphConfig, []int, error)
	UpsertWithID(id string, obj *storage.NetworkGraphConfig) error
	UpsertManyWithIDs(ids []string, objs []*storage.NetworkGraphConfig) error
	Delete(id string) error
	DeleteMany(ids []string) error
	WalkAllWithID(fn func(id string, obj *storage.NetworkGraphConfig) error) error
	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	crud db.Crud
}

func alloc() proto.Message {
	return &storage.NetworkGraphConfig{}
}

// New returns a new Store instance using the provided rocksdb instance.
func New(db *rocksdb.RocksDB) Store {
	globaldb.RegisterBucket(bucket, "NetworkGraphConfig")
	return &storeImpl{
		crud: generic.NewCRUD(db, bucket, nil, alloc, false),
	}
}

// Count returns the number of objects in the store
func (b *storeImpl) Count() (int, error) {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.Count, "NetworkGraphConfig")

	return b.crud.Count()
}

// Exists returns if the id exists in the store
func (b *storeImpl) Exists(id string) (bool, error) {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.Exists, "NetworkGraphConfig")

	return b.crud.Exists(id)
}

// GetIDs returns all the IDs for the store
func (b *storeImpl) GetIDs() ([]string, error) {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.GetAll, "NetworkGraphConfigIDs")

	return b.crud.GetKeys()
}

// Get returns the object, if it exists from the store
func (b *storeImpl) Get(id string) (*storage.NetworkGraphConfig, bool, error) {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.Get, "NetworkGraphConfig")

	msg, exists, err := b.crud.Get(id)
	if err != nil || !exists {
		return nil, false, err
	}
	return msg.(*storage.NetworkGraphConfig), true, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice 
func (b *storeImpl) GetMany(ids []string) ([]*storage.NetworkGraphConfig, []int, error) {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.GetMany, "NetworkGraphConfig")

	msgs, missingIndices, err := b.crud.GetMany(ids)
	if err != nil {
		return nil, nil, err
	}
	objs := make([]*storage.NetworkGraphConfig, 0, len(msgs))
	for _, m := range msgs {
		objs = append(objs, m.(*storage.NetworkGraphConfig))
	}
	return objs, missingIndices, nil
}
// UpsertWithID inserts the object into the DB
func (b *storeImpl) UpsertWithID(id string, obj *storage.NetworkGraphConfig) error {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.Add, "NetworkGraphConfig")

	return b.crud.UpsertWithID(id, obj)
}

// UpsertManyWithIDs batches objects into the DB
func (b *storeImpl) UpsertManyWithIDs(ids []string, objs []*storage.NetworkGraphConfig) error {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.AddMany, "NetworkGraphConfig")

	msgs := make([]proto.Message, 0, len(objs))
	for _, o := range objs {
		msgs = append(msgs, o)
    }

	return b.crud.UpsertManyWithIDs(ids, msgs)
}

// Delete removes the specified ID from the store
func (b *storeImpl) Delete(id string) error {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.Remove, "NetworkGraphConfig")

	return b.crud.Delete(id)
}

// Delete removes the specified IDs from the store
func (b *storeImpl) DeleteMany(ids []string) error {
	defer metrics.SetRocksDBOperationDurationTime(time.Now(), ops.RemoveMany, "NetworkGraphConfig")

	return b.crud.DeleteMany(ids)
}
// WalkAllWithID iterates over all of the objects in the store and applies the closure
func (b *storeImpl) WalkAllWithID(fn func(id string, obj *storage.NetworkGraphConfig) error) error {
	return b.crud.WalkAllWithID(func(id []byte, msg proto.Message) error {
		return fn(string(id), msg.(*storage.NetworkGraphConfig))
	})
}

// AckKeysIndexed acknowledges the passed keys were indexed
func (b *storeImpl) AckKeysIndexed(keys ...string) error {
	return b.crud.AckKeysIndexed(keys...)
}

// GetKeysToIndex returns the keys that need to be indexed
func (b *storeImpl) GetKeysToIndex() ([]string, error) {
	return b.crud.GetKeysToIndex()
}
