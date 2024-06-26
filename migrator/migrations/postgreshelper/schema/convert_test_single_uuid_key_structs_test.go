// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTestSingleUUIDKeyStructSerialization(t *testing.T) {
	obj := &storage.TestSingleUUIDKeyStruct{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertTestSingleUUIDKeyStructFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertTestSingleUUIDKeyStructToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
