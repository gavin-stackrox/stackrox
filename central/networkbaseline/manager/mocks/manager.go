// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	networkgraph "github.com/stackrox/rox/pkg/networkgraph"
	timestamp "github.com/stackrox/rox/pkg/timestamp"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// ProcessDeploymentCreate mocks base method
func (m *MockManager) ProcessDeploymentCreate(deploymentID, deploymentName, clusterID, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDeploymentCreate", deploymentID, deploymentName, clusterID, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessDeploymentCreate indicates an expected call of ProcessDeploymentCreate
func (mr *MockManagerMockRecorder) ProcessDeploymentCreate(deploymentID, deploymentName, clusterID, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDeploymentCreate", reflect.TypeOf((*MockManager)(nil).ProcessDeploymentCreate), deploymentID, deploymentName, clusterID, namespace)
}

// ProcessDeploymentDelete mocks base method
func (m *MockManager) ProcessDeploymentDelete(deploymentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDeploymentDelete", deploymentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessDeploymentDelete indicates an expected call of ProcessDeploymentDelete
func (mr *MockManagerMockRecorder) ProcessDeploymentDelete(deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDeploymentDelete", reflect.TypeOf((*MockManager)(nil).ProcessDeploymentDelete), deploymentID)
}

// ProcessFlowUpdate mocks base method
func (m *MockManager) ProcessFlowUpdate(flows map[networkgraph.NetworkConnIndicator]timestamp.MicroTS) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessFlowUpdate", flows)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessFlowUpdate indicates an expected call of ProcessFlowUpdate
func (mr *MockManagerMockRecorder) ProcessFlowUpdate(flows interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessFlowUpdate", reflect.TypeOf((*MockManager)(nil).ProcessFlowUpdate), flows)
}

// ProcessBaselineStatusUpdate mocks base method
func (m *MockManager) ProcessBaselineStatusUpdate(ctx context.Context, modifyRequest *v1.ModifyBaselineStatusForPeersRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBaselineStatusUpdate", ctx, modifyRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBaselineStatusUpdate indicates an expected call of ProcessBaselineStatusUpdate
func (mr *MockManagerMockRecorder) ProcessBaselineStatusUpdate(ctx, modifyRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBaselineStatusUpdate", reflect.TypeOf((*MockManager)(nil).ProcessBaselineStatusUpdate), ctx, modifyRequest)
}

// ProcessNetworkPolicyUpdate mocks base method
func (m *MockManager) ProcessNetworkPolicyUpdate(ctx context.Context, action central.ResourceAction, policy *storage.NetworkPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNetworkPolicyUpdate", ctx, action, policy)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessNetworkPolicyUpdate indicates an expected call of ProcessNetworkPolicyUpdate
func (mr *MockManagerMockRecorder) ProcessNetworkPolicyUpdate(ctx, action, policy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNetworkPolicyUpdate", reflect.TypeOf((*MockManager)(nil).ProcessNetworkPolicyUpdate), ctx, action, policy)
}
