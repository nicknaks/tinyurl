// Code generated by MockGen. DO NOT EDIT.
// Source: url.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	models "tinyUrl/internal/pkg/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUrlRepositoryInterface is a mock of UrlRepositoryInterface interface.
type MockUrlRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUrlRepositoryInterfaceMockRecorder
}

// MockUrlRepositoryInterfaceMockRecorder is the mock recorder for MockUrlRepositoryInterface.
type MockUrlRepositoryInterfaceMockRecorder struct {
	mock *MockUrlRepositoryInterface
}

// NewMockUrlRepositoryInterface creates a new mock instance.
func NewMockUrlRepositoryInterface(ctrl *gomock.Controller) *MockUrlRepositoryInterface {
	mock := &MockUrlRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUrlRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlRepositoryInterface) EXPECT() *MockUrlRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddTinyURLBYURL mocks base method.
func (m *MockUrlRepositoryInterface) AddTinyURLBYURL(url, tinyUrl models.Url) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTinyURLBYURL", url, tinyUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTinyURLBYURL indicates an expected call of AddTinyURLBYURL.
func (mr *MockUrlRepositoryInterfaceMockRecorder) AddTinyURLBYURL(url, tinyUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTinyURLBYURL", reflect.TypeOf((*MockUrlRepositoryInterface)(nil).AddTinyURLBYURL), url, tinyUrl)
}

// GetTinyUrlByUrl mocks base method.
func (m *MockUrlRepositoryInterface) GetTinyUrlByUrl(url models.Url) (models.Url, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTinyUrlByUrl", url)
	ret0, _ := ret[0].(models.Url)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTinyUrlByUrl indicates an expected call of GetTinyUrlByUrl.
func (mr *MockUrlRepositoryInterfaceMockRecorder) GetTinyUrlByUrl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTinyUrlByUrl", reflect.TypeOf((*MockUrlRepositoryInterface)(nil).GetTinyUrlByUrl), url)
}

// GetURLByTinyURL mocks base method.
func (m *MockUrlRepositoryInterface) GetURLByTinyURL(tinyUrl models.Url) (models.Url, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURLByTinyURL", tinyUrl)
	ret0, _ := ret[0].(models.Url)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURLByTinyURL indicates an expected call of GetURLByTinyURL.
func (mr *MockUrlRepositoryInterfaceMockRecorder) GetURLByTinyURL(tinyUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURLByTinyURL", reflect.TypeOf((*MockUrlRepositoryInterface)(nil).GetURLByTinyURL), tinyUrl)
}
