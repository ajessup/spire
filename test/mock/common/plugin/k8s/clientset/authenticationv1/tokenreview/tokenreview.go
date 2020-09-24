// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/authentication/v1 (interfaces: TokenReviewInterface)

// Package mock_tokenreview is a generated GoMock package.
package mock_tokenreview

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/authentication/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
)

// MockTokenReviewInterface is a mock of TokenReviewInterface interface
type MockTokenReviewInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTokenReviewInterfaceMockRecorder
}

// MockTokenReviewInterfaceMockRecorder is the mock recorder for MockTokenReviewInterface
type MockTokenReviewInterfaceMockRecorder struct {
	mock *MockTokenReviewInterface
}

// NewMockTokenReviewInterface creates a new mock instance
func NewMockTokenReviewInterface(ctrl *gomock.Controller) *MockTokenReviewInterface {
	mock := &MockTokenReviewInterface{ctrl: ctrl}
	mock.recorder = &MockTokenReviewInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenReviewInterface) EXPECT() *MockTokenReviewInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTokenReviewInterface) Create(arg0 context.Context, arg1 *v1.TokenReview, arg2 v10.CreateOptions) (*v1.TokenReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.TokenReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTokenReviewInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTokenReviewInterface)(nil).Create), arg0, arg1, arg2)
}