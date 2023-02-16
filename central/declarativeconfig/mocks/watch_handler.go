// Code generated by MockGen. DO NOT EDIT.
// Source: watch_handler.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockdeclarativeConfigReconciler is a mock of declarativeConfigReconciler interface.
type MockdeclarativeConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockdeclarativeConfigReconcilerMockRecorder
}

// MockdeclarativeConfigReconcilerMockRecorder is the mock recorder for MockdeclarativeConfigReconciler.
type MockdeclarativeConfigReconcilerMockRecorder struct {
	mock *MockdeclarativeConfigReconciler
}

// NewMockdeclarativeConfigReconciler creates a new mock instance.
func NewMockdeclarativeConfigReconciler(ctrl *gomock.Controller) *MockdeclarativeConfigReconciler {
	mock := &MockdeclarativeConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockdeclarativeConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdeclarativeConfigReconciler) EXPECT() *MockdeclarativeConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDeclarativeConfigs mocks base method.
func (m *MockdeclarativeConfigReconciler) ReconcileDeclarativeConfigs(fileContents [][]byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileDeclarativeConfigs", fileContents)
}

// ReconcileDeclarativeConfigs indicates an expected call of ReconcileDeclarativeConfigs.
func (mr *MockdeclarativeConfigReconcilerMockRecorder) ReconcileDeclarativeConfigs(fileContents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeclarativeConfigs", reflect.TypeOf((*MockdeclarativeConfigReconciler)(nil).ReconcileDeclarativeConfigs), fileContents)
}
