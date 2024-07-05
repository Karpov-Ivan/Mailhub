// Code generated by MockGen. DO NOT EDIT.
// Source: ./session_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	proto "mail/internal/microservice/session/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSessionServiceClient is a mock of SessionServiceClient interface.
type MockSessionServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSessionServiceClientMockRecorder
}

// MockSessionServiceClientMockRecorder is the mock recorder for MockSessionServiceClient.
type MockSessionServiceClientMockRecorder struct {
	mock *MockSessionServiceClient
}

// NewMockSessionServiceClient creates a new mock instance.
func NewMockSessionServiceClient(ctrl *gomock.Controller) *MockSessionServiceClient {
	mock := &MockSessionServiceClient{ctrl: ctrl}
	mock.recorder = &MockSessionServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionServiceClient) EXPECT() *MockSessionServiceClientMockRecorder {
	return m.recorder
}

// CleanupExpiredSessions mock base method.
func (m *MockSessionServiceClient) CleanupExpiredSessions(ctx context.Context, in *proto.CleanupExpiredSessionsRequest, opts ...grpc.CallOption) (*proto.CleanupExpiredSessionsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CleanupExpiredSessions", varargs...)
	ret0, _ := ret[0].(*proto.CleanupExpiredSessionsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanupExpiredSessions indicates an expected call of CleanupExpiredSessions.
func (mr *MockSessionServiceClientMockRecorder) CleanupExpiredSessions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupExpiredSessions", reflect.TypeOf((*MockSessionServiceClient)(nil).CleanupExpiredSessions), varargs...)
}

// CreateSession mock base method.
func (m *MockSessionServiceClient) CreateSession(ctx context.Context, in *proto.CreateSessionRequest, opts ...grpc.CallOption) (*proto.CreateSessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSession", varargs...)
	ret0, _ := ret[0].(*proto.CreateSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionServiceClientMockRecorder) CreateSession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionServiceClient)(nil).CreateSession), varargs...)
}

// DeleteSession mock base method.
func (m *MockSessionServiceClient) DeleteSession(ctx context.Context, in *proto.DeleteSessionRequest, opts ...grpc.CallOption) (*proto.DeleteSessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSession", varargs...)
	ret0, _ := ret[0].(*proto.DeleteSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionServiceClientMockRecorder) DeleteSession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionServiceClient)(nil).DeleteSession), varargs...)
}

// GetLoginBySession mock base method.
func (m *MockSessionServiceClient) GetLoginBySession(ctx context.Context, in *proto.GetLoginBySessionRequest, opts ...grpc.CallOption) (*proto.GetLoginBySessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoginBySession", varargs...)
	ret0, _ := ret[0].(*proto.GetLoginBySessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginBySession indicates an expected call of GetLoginBySession.
func (mr *MockSessionServiceClientMockRecorder) GetLoginBySession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginBySession", reflect.TypeOf((*MockSessionServiceClient)(nil).GetLoginBySession), varargs...)
}

// GetProfileIDBySession mock base method.
func (m *MockSessionServiceClient) GetProfileIDBySession(ctx context.Context, in *proto.GetLoginBySessionRequest, opts ...grpc.CallOption) (*proto.GetProfileIDBySessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProfileIDBySession", varargs...)
	ret0, _ := ret[0].(*proto.GetProfileIDBySessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIDBySession indicates an expected call of GetProfileIDBySession.
func (mr *MockSessionServiceClientMockRecorder) GetProfileIDBySession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIDBySession", reflect.TypeOf((*MockSessionServiceClient)(nil).GetProfileIDBySession), varargs...)
}

// GetSession mock base method.
func (m *MockSessionServiceClient) GetSession(ctx context.Context, in *proto.GetSessionRequest, opts ...grpc.CallOption) (*proto.GetSessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSession", varargs...)
	ret0, _ := ret[0].(*proto.GetSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockSessionServiceClientMockRecorder) GetSession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionServiceClient)(nil).GetSession), varargs...)
}

// MockSessionServiceServer is a mock of SessionServiceServer interface.
type MockSessionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSessionServiceServerMockRecorder
}

// MockSessionServiceServerMockRecorder is the mock recorder for MockSessionServiceServer.
type MockSessionServiceServerMockRecorder struct {
	mock *MockSessionServiceServer
}

// NewMockSessionServiceServer creates a new mock instance.
func NewMockSessionServiceServer(ctrl *gomock.Controller) *MockSessionServiceServer {
	mock := &MockSessionServiceServer{ctrl: ctrl}
	mock.recorder = &MockSessionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionServiceServer) EXPECT() *MockSessionServiceServerMockRecorder {
	return m.recorder
}

// CleanupExpiredSessions mock base method.
func (m *MockSessionServiceServer) CleanupExpiredSessions(arg0 context.Context, arg1 *proto.CleanupExpiredSessionsRequest) (*proto.CleanupExpiredSessionsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupExpiredSessions", arg0, arg1)
	ret0, _ := ret[0].(*proto.CleanupExpiredSessionsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanupExpiredSessions indicates an expected call of CleanupExpiredSessions.
func (mr *MockSessionServiceServerMockRecorder) CleanupExpiredSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupExpiredSessions", reflect.TypeOf((*MockSessionServiceServer)(nil).CleanupExpiredSessions), arg0, arg1)
}

// CreateSession mock base method.
func (m *MockSessionServiceServer) CreateSession(arg0 context.Context, arg1 *proto.CreateSessionRequest) (*proto.CreateSessionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreateSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionServiceServerMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionServiceServer)(nil).CreateSession), arg0, arg1)
}

// DeleteSession mock base method.
func (m *MockSessionServiceServer) DeleteSession(arg0 context.Context, arg1 *proto.DeleteSessionRequest) (*proto.DeleteSessionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeleteSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionServiceServerMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionServiceServer)(nil).DeleteSession), arg0, arg1)
}

// GetLoginBySession mock base method.
func (m *MockSessionServiceServer) GetLoginBySession(arg0 context.Context, arg1 *proto.GetLoginBySessionRequest) (*proto.GetLoginBySessionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginBySession", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetLoginBySessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginBySession indicates an expected call of GetLoginBySession.
func (mr *MockSessionServiceServerMockRecorder) GetLoginBySession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginBySession", reflect.TypeOf((*MockSessionServiceServer)(nil).GetLoginBySession), arg0, arg1)
}

// GetProfileIDBySession mock base method.
func (m *MockSessionServiceServer) GetProfileIDBySession(arg0 context.Context, arg1 *proto.GetLoginBySessionRequest) (*proto.GetProfileIDBySessionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileIDBySession", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetProfileIDBySessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIDBySession indicates an expected call of GetProfileIDBySession.
func (mr *MockSessionServiceServerMockRecorder) GetProfileIDBySession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIDBySession", reflect.TypeOf((*MockSessionServiceServer)(nil).GetProfileIDBySession), arg0, arg1)
}

// GetSession mock base method.
func (m *MockSessionServiceServer) GetSession(arg0 context.Context, arg1 *proto.GetSessionRequest) (*proto.GetSessionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockSessionServiceServerMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionServiceServer)(nil).GetSession), arg0, arg1)
}

// mustEmbedUnimplementedSessionServiceServer mock base method.
func (m *MockSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSessionServiceServer")
}

// mustEmbedUnimplementedSessionServiceServer indicates an expected call of mustEmbedUnimplementedSessionServiceServer.
func (mr *MockSessionServiceServerMockRecorder) mustEmbedUnimplementedSessionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSessionServiceServer", reflect.TypeOf((*MockSessionServiceServer)(nil).mustEmbedUnimplementedSessionServiceServer))
}

// MockUnsafeSessionServiceServer is a mock of UnsafeSessionServiceServer interface.
type MockUnsafeSessionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSessionServiceServerMockRecorder
}

// MockUnsafeSessionServiceServerMockRecorder is the mock recorder for MockUnsafeSessionServiceServer.
type MockUnsafeSessionServiceServerMockRecorder struct {
	mock *MockUnsafeSessionServiceServer
}

// NewMockUnsafeSessionServiceServer creates a new mock instance.
func NewMockUnsafeSessionServiceServer(ctrl *gomock.Controller) *MockUnsafeSessionServiceServer {
	mock := &MockUnsafeSessionServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSessionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSessionServiceServer) EXPECT() *MockUnsafeSessionServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSessionServiceServer mock base method.
func (m *MockUnsafeSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSessionServiceServer")
}

// mustEmbedUnimplementedSessionServiceServer indicates an expected call of mustEmbedUnimplementedSessionServiceServer.
func (mr *MockUnsafeSessionServiceServerMockRecorder) mustEmbedUnimplementedSessionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSessionServiceServer", reflect.TypeOf((*MockUnsafeSessionServiceServer)(nil).mustEmbedUnimplementedSessionServiceServer))
}
