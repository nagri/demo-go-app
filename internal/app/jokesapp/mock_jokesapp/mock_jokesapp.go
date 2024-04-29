// Code generated by MockGen. DO NOT EDIT.
// Source: jokesapp.go
//
// Generated by this command:
//
//	mockgen --source jokesapp.go --destination=mock_jokesapp/mock_jokesapp.go
//

// Package mock_jokesapp is a generated GoMock package.
package mock_jokesapp

import (
	reflect "reflect"

	jokesapi "github.com/nagri/demo-go-app/internal/pkg/jokesapi"
	gomock "go.uber.org/mock/gomock"
)

// MockJokeinterface is a mock of Jokeinterface interface.
type MockJokeinterface struct {
	ctrl     *gomock.Controller
	recorder *MockJokeinterfaceMockRecorder
}

// MockJokeinterfaceMockRecorder is the mock recorder for MockJokeinterface.
type MockJokeinterfaceMockRecorder struct {
	mock *MockJokeinterface
}

// NewMockJokeinterface creates a new mock instance.
func NewMockJokeinterface(ctrl *gomock.Controller) *MockJokeinterface {
	mock := &MockJokeinterface{ctrl: ctrl}
	mock.recorder = &MockJokeinterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJokeinterface) EXPECT() *MockJokeinterfaceMockRecorder {
	return m.recorder
}

// GetJoke mocks base method.
func (m *MockJokeinterface) GetJoke() (*jokesapi.Joke, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJoke")
	ret0, _ := ret[0].(*jokesapi.Joke)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJoke indicates an expected call of GetJoke.
func (mr *MockJokeinterfaceMockRecorder) GetJoke() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJoke", reflect.TypeOf((*MockJokeinterface)(nil).GetJoke))
}