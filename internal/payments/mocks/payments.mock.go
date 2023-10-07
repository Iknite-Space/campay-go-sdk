// Code generated by MockGen. DO NOT EDIT.
// Source: ./payment.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	payment "github.com/Iknite-Space/campay-go-sdk/internal/models/payment"
	gomock "github.com/golang/mock/gomock"
)

// MockPaymentService is a mock of PaymentService interface.
type MockPaymentService struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceMockRecorder
}

// MockPaymentServiceMockRecorder is the mock recorder for MockPaymentService.
type MockPaymentServiceMockRecorder struct {
	mock *MockPaymentService
}

// NewMockPaymentService creates a new mock instance.
func NewMockPaymentService(ctrl *gomock.Controller) *MockPaymentService {
	mock := &MockPaymentService{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentService) EXPECT() *MockPaymentServiceMockRecorder {
	return m.recorder
}

// InitiatePayments mocks base method.
func (m *MockPaymentService) InitiatePayments(ctx context.Context, req payment.RequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiatePayments", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitiatePayments indicates an expected call of InitiatePayments.
func (mr *MockPaymentServiceMockRecorder) InitiatePayments(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiatePayments", reflect.TypeOf((*MockPaymentService)(nil).InitiatePayments), ctx, req)
}
