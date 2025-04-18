// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/registry (interfaces: VMGetter)
//
// Generated by this command:
//
//	mockgen -package=registrymock -destination=registrymock/vm_getter.go -mock_names=VMGetter=VMGetter . VMGetter
//

// Package registrymock is a generated GoMock package.
package registrymock

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	vms "github.com/ava-labs/avalanchego/vms"
	gomock "go.uber.org/mock/gomock"
)

// VMGetter is a mock of VMGetter interface.
type VMGetter struct {
	ctrl     *gomock.Controller
	recorder *VMGetterMockRecorder
	isgomock struct{}
}

// VMGetterMockRecorder is the mock recorder for VMGetter.
type VMGetterMockRecorder struct {
	mock *VMGetter
}

// NewVMGetter creates a new mock instance.
func NewVMGetter(ctrl *gomock.Controller) *VMGetter {
	mock := &VMGetter{ctrl: ctrl}
	mock.recorder = &VMGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *VMGetter) EXPECT() *VMGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *VMGetter) Get() (map[ids.ID]vms.Factory, map[ids.ID]vms.Factory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(map[ids.ID]vms.Factory)
	ret1, _ := ret[1].(map[ids.ID]vms.Factory)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *VMGetterMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*VMGetter)(nil).Get))
}
