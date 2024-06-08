// Code generated by MockGen. DO NOT EDIT.
// Source: internal/genproto/auditpb/hub_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=internal/genproto/auditpb/hub_grpc.pb.go -destination=internal/mock/grpc/audit_client.go -package=mockgrpc
//

// Package mockgrpc is a generated GoMock package.
package mockgrpc

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuditServiceClient is a mock of AuditServiceClient interface.
type MockAuditServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditServiceClientMockRecorder
}

// MockAuditServiceClientMockRecorder is the mock recorder for MockAuditServiceClient.
type MockAuditServiceClientMockRecorder struct {
	mock *MockAuditServiceClient
}

// NewMockAuditServiceClient creates a new mock instance.
func NewMockAuditServiceClient(ctrl *gomock.Controller) *MockAuditServiceClient {
	mock := &MockAuditServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuditServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditServiceClient) EXPECT() *MockAuditServiceClientMockRecorder {
	return m.recorder
}

// MockAuditServiceServer is a mock of AuditServiceServer interface.
type MockAuditServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuditServiceServerMockRecorder
}

// MockAuditServiceServerMockRecorder is the mock recorder for MockAuditServiceServer.
type MockAuditServiceServerMockRecorder struct {
	mock *MockAuditServiceServer
}

// NewMockAuditServiceServer creates a new mock instance.
func NewMockAuditServiceServer(ctrl *gomock.Controller) *MockAuditServiceServer {
	mock := &MockAuditServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuditServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditServiceServer) EXPECT() *MockAuditServiceServerMockRecorder {
	return m.recorder
}

// MockUnsafeAuditServiceServer is a mock of UnsafeAuditServiceServer interface.
type MockUnsafeAuditServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuditServiceServerMockRecorder
}

// MockUnsafeAuditServiceServerMockRecorder is the mock recorder for MockUnsafeAuditServiceServer.
type MockUnsafeAuditServiceServerMockRecorder struct {
	mock *MockUnsafeAuditServiceServer
}

// NewMockUnsafeAuditServiceServer creates a new mock instance.
func NewMockUnsafeAuditServiceServer(ctrl *gomock.Controller) *MockUnsafeAuditServiceServer {
	mock := &MockUnsafeAuditServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuditServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuditServiceServer) EXPECT() *MockUnsafeAuditServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuditServiceServer mocks base method.
func (m *MockUnsafeAuditServiceServer) mustEmbedUnimplementedAuditServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuditServiceServer")
}

// mustEmbedUnimplementedAuditServiceServer indicates an expected call of mustEmbedUnimplementedAuditServiceServer.
func (mr *MockUnsafeAuditServiceServerMockRecorder) mustEmbedUnimplementedAuditServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuditServiceServer", reflect.TypeOf((*MockUnsafeAuditServiceServer)(nil).mustEmbedUnimplementedAuditServiceServer))
}
