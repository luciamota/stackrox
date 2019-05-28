// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/namespace/index (interfaces: Indexer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
)

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddNamespaceMetadata mocks base method
func (m *MockIndexer) AddNamespaceMetadata(arg0 *storage.NamespaceMetadata) error {
	ret := m.ctrl.Call(m, "AddNamespaceMetadata", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNamespaceMetadata indicates an expected call of AddNamespaceMetadata
func (mr *MockIndexerMockRecorder) AddNamespaceMetadata(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNamespaceMetadata", reflect.TypeOf((*MockIndexer)(nil).AddNamespaceMetadata), arg0)
}

// AddNamespaceMetadatas mocks base method
func (m *MockIndexer) AddNamespaceMetadatas(arg0 []*storage.NamespaceMetadata) error {
	ret := m.ctrl.Call(m, "AddNamespaceMetadatas", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNamespaceMetadatas indicates an expected call of AddNamespaceMetadatas
func (mr *MockIndexerMockRecorder) AddNamespaceMetadatas(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNamespaceMetadatas", reflect.TypeOf((*MockIndexer)(nil).AddNamespaceMetadatas), arg0)
}

// DeleteNamespaceMetadata mocks base method
func (m *MockIndexer) DeleteNamespaceMetadata(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteNamespaceMetadata", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespaceMetadata indicates an expected call of DeleteNamespaceMetadata
func (mr *MockIndexerMockRecorder) DeleteNamespaceMetadata(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespaceMetadata", reflect.TypeOf((*MockIndexer)(nil).DeleteNamespaceMetadata), arg0)
}

// Search mocks base method
func (m *MockIndexer) Search(arg0 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), arg0)
}
