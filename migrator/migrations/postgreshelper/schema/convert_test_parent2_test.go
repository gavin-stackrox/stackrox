// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTestParent2Serialization(t *testing.T) {
	obj := &storage.TestParent2{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertTestParent2FromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertTestParent2ToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
