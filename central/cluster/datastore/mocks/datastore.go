// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/cluster/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
	time "time"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddCluster mocks base method
func (m *MockDataStore) AddCluster(arg0 *storage.Cluster) (string, error) {
	ret := m.ctrl.Call(m, "AddCluster", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCluster indicates an expected call of AddCluster
func (mr *MockDataStoreMockRecorder) AddCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockDataStore)(nil).AddCluster), arg0)
}

// CountClusters mocks base method
func (m *MockDataStore) CountClusters() (int, error) {
	ret := m.ctrl.Call(m, "CountClusters")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountClusters indicates an expected call of CountClusters
func (mr *MockDataStoreMockRecorder) CountClusters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountClusters", reflect.TypeOf((*MockDataStore)(nil).CountClusters))
}

// GetCluster mocks base method
func (m *MockDataStore) GetCluster(arg0 string) (*storage.Cluster, bool, error) {
	ret := m.ctrl.Call(m, "GetCluster", arg0)
	ret0, _ := ret[0].(*storage.Cluster)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockDataStoreMockRecorder) GetCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockDataStore)(nil).GetCluster), arg0)
}

// GetClusters mocks base method
func (m *MockDataStore) GetClusters() ([]*storage.Cluster, error) {
	ret := m.ctrl.Call(m, "GetClusters")
	ret0, _ := ret[0].([]*storage.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters
func (mr *MockDataStoreMockRecorder) GetClusters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockDataStore)(nil).GetClusters))
}

// RemoveCluster mocks base method
func (m *MockDataStore) RemoveCluster(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveCluster", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCluster indicates an expected call of RemoveCluster
func (mr *MockDataStoreMockRecorder) RemoveCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCluster", reflect.TypeOf((*MockDataStore)(nil).RemoveCluster), arg0)
}

// Search mocks base method
func (m *MockDataStore) Search(arg0 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), arg0)
}

// UpdateCluster mocks base method
func (m *MockDataStore) UpdateCluster(arg0 *storage.Cluster) error {
	ret := m.ctrl.Call(m, "UpdateCluster", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockDataStoreMockRecorder) UpdateCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockDataStore)(nil).UpdateCluster), arg0)
}

// UpdateClusterContactTime mocks base method
func (m *MockDataStore) UpdateClusterContactTime(arg0 string, arg1 time.Time) error {
	ret := m.ctrl.Call(m, "UpdateClusterContactTime", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterContactTime indicates an expected call of UpdateClusterContactTime
func (mr *MockDataStoreMockRecorder) UpdateClusterContactTime(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterContactTime", reflect.TypeOf((*MockDataStore)(nil).UpdateClusterContactTime), arg0, arg1)
}

// UpdateClusterStatus mocks base method
func (m *MockDataStore) UpdateClusterStatus(arg0 string, arg1 *storage.ClusterStatus) error {
	ret := m.ctrl.Call(m, "UpdateClusterStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterStatus indicates an expected call of UpdateClusterStatus
func (mr *MockDataStoreMockRecorder) UpdateClusterStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterStatus", reflect.TypeOf((*MockDataStore)(nil).UpdateClusterStatus), arg0, arg1)
}
