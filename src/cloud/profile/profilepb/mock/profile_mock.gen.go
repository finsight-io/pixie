// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_profilepb is a generated GoMock package.
package mock_profilepb

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	uuidpb "px.dev/pixie/src/api/proto/uuidpb"
	profilepb "px.dev/pixie/src/cloud/profile/profilepb"
)

// MockProfileServiceClient is a mock of ProfileServiceClient interface.
type MockProfileServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceClientMockRecorder
}

// MockProfileServiceClientMockRecorder is the mock recorder for MockProfileServiceClient.
type MockProfileServiceClientMockRecorder struct {
	mock *MockProfileServiceClient
}

// NewMockProfileServiceClient creates a new mock instance.
func NewMockProfileServiceClient(ctrl *gomock.Controller) *MockProfileServiceClient {
	mock := &MockProfileServiceClient{ctrl: ctrl}
	mock.recorder = &MockProfileServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileServiceClient) EXPECT() *MockProfileServiceClientMockRecorder {
	return m.recorder
}

// CreateOrgAndUser mocks base method.
func (m *MockProfileServiceClient) CreateOrgAndUser(ctx context.Context, in *profilepb.CreateOrgAndUserRequest, opts ...grpc.CallOption) (*profilepb.CreateOrgAndUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrgAndUser", varargs...)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser.
func (mr *MockProfileServiceClientMockRecorder) CreateOrgAndUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateOrgAndUser), varargs...)
}

// CreateUser mocks base method.
func (m *MockProfileServiceClient) CreateUser(ctx context.Context, in *profilepb.CreateUserRequest, opts ...grpc.CallOption) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockProfileServiceClientMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateUser), varargs...)
}

// GetUser mocks base method.
func (m *MockProfileServiceClient) GetUser(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockProfileServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUser), varargs...)
}

// GetUserAttributes mocks base method.
func (m *MockProfileServiceClient) GetUserAttributes(ctx context.Context, in *profilepb.GetUserAttributesRequest, opts ...grpc.CallOption) (*profilepb.GetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserAttributes", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttributes indicates an expected call of GetUserAttributes.
func (mr *MockProfileServiceClientMockRecorder) GetUserAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttributes", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserAttributes), varargs...)
}

// GetUserByAuthProviderID mocks base method.
func (m *MockProfileServiceClient) GetUserByAuthProviderID(ctx context.Context, in *profilepb.GetUserByAuthProviderIDRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByAuthProviderID", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuthProviderID indicates an expected call of GetUserByAuthProviderID.
func (mr *MockProfileServiceClientMockRecorder) GetUserByAuthProviderID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuthProviderID", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserByAuthProviderID), varargs...)
}

// GetUserByEmail mocks base method.
func (m *MockProfileServiceClient) GetUserByEmail(ctx context.Context, in *profilepb.GetUserByEmailRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByEmail", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockProfileServiceClientMockRecorder) GetUserByEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserByEmail), varargs...)
}

// GetUserSettings mocks base method.
func (m *MockProfileServiceClient) GetUserSettings(ctx context.Context, in *profilepb.GetUserSettingsRequest, opts ...grpc.CallOption) (*profilepb.GetUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserSettings", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSettings indicates an expected call of GetUserSettings.
func (mr *MockProfileServiceClientMockRecorder) GetUserSettings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSettings", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserSettings), varargs...)
}

// SetUserAttributes mocks base method.
func (m *MockProfileServiceClient) SetUserAttributes(ctx context.Context, in *profilepb.SetUserAttributesRequest, opts ...grpc.CallOption) (*profilepb.SetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetUserAttributes", varargs...)
	ret0, _ := ret[0].(*profilepb.SetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserAttributes indicates an expected call of SetUserAttributes.
func (mr *MockProfileServiceClientMockRecorder) SetUserAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAttributes", reflect.TypeOf((*MockProfileServiceClient)(nil).SetUserAttributes), varargs...)
}

// UpdateUser mocks base method.
func (m *MockProfileServiceClient) UpdateUser(ctx context.Context, in *profilepb.UpdateUserRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockProfileServiceClientMockRecorder) UpdateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateUser), varargs...)
}

// UpdateUserSettings mocks base method.
func (m *MockProfileServiceClient) UpdateUserSettings(ctx context.Context, in *profilepb.UpdateUserSettingsRequest, opts ...grpc.CallOption) (*profilepb.UpdateUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserSettings", varargs...)
	ret0, _ := ret[0].(*profilepb.UpdateUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSettings indicates an expected call of UpdateUserSettings.
func (mr *MockProfileServiceClientMockRecorder) UpdateUserSettings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSettings", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateUserSettings), varargs...)
}

// MockProfileServiceServer is a mock of ProfileServiceServer interface.
type MockProfileServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceServerMockRecorder
}

// MockProfileServiceServerMockRecorder is the mock recorder for MockProfileServiceServer.
type MockProfileServiceServerMockRecorder struct {
	mock *MockProfileServiceServer
}

// NewMockProfileServiceServer creates a new mock instance.
func NewMockProfileServiceServer(ctrl *gomock.Controller) *MockProfileServiceServer {
	mock := &MockProfileServiceServer{ctrl: ctrl}
	mock.recorder = &MockProfileServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileServiceServer) EXPECT() *MockProfileServiceServerMockRecorder {
	return m.recorder
}

// CreateOrgAndUser mocks base method.
func (m *MockProfileServiceServer) CreateOrgAndUser(arg0 context.Context, arg1 *profilepb.CreateOrgAndUserRequest) (*profilepb.CreateOrgAndUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgAndUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser.
func (mr *MockProfileServiceServerMockRecorder) CreateOrgAndUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateOrgAndUser), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockProfileServiceServer) CreateUser(arg0 context.Context, arg1 *profilepb.CreateUserRequest) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockProfileServiceServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateUser), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockProfileServiceServer) GetUser(arg0 context.Context, arg1 *uuidpb.UUID) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockProfileServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUser), arg0, arg1)
}

// GetUserAttributes mocks base method.
func (m *MockProfileServiceServer) GetUserAttributes(arg0 context.Context, arg1 *profilepb.GetUserAttributesRequest) (*profilepb.GetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAttributes", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttributes indicates an expected call of GetUserAttributes.
func (mr *MockProfileServiceServerMockRecorder) GetUserAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttributes", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserAttributes), arg0, arg1)
}

// GetUserByAuthProviderID mocks base method.
func (m *MockProfileServiceServer) GetUserByAuthProviderID(arg0 context.Context, arg1 *profilepb.GetUserByAuthProviderIDRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuthProviderID", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuthProviderID indicates an expected call of GetUserByAuthProviderID.
func (mr *MockProfileServiceServerMockRecorder) GetUserByAuthProviderID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuthProviderID", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserByAuthProviderID), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockProfileServiceServer) GetUserByEmail(arg0 context.Context, arg1 *profilepb.GetUserByEmailRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockProfileServiceServerMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserSettings mocks base method.
func (m *MockProfileServiceServer) GetUserSettings(arg0 context.Context, arg1 *profilepb.GetUserSettingsRequest) (*profilepb.GetUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSettings", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSettings indicates an expected call of GetUserSettings.
func (mr *MockProfileServiceServerMockRecorder) GetUserSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSettings", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserSettings), arg0, arg1)
}

// SetUserAttributes mocks base method.
func (m *MockProfileServiceServer) SetUserAttributes(arg0 context.Context, arg1 *profilepb.SetUserAttributesRequest) (*profilepb.SetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAttributes", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.SetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserAttributes indicates an expected call of SetUserAttributes.
func (mr *MockProfileServiceServerMockRecorder) SetUserAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAttributes", reflect.TypeOf((*MockProfileServiceServer)(nil).SetUserAttributes), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockProfileServiceServer) UpdateUser(arg0 context.Context, arg1 *profilepb.UpdateUserRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockProfileServiceServerMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserSettings mocks base method.
func (m *MockProfileServiceServer) UpdateUserSettings(arg0 context.Context, arg1 *profilepb.UpdateUserSettingsRequest) (*profilepb.UpdateUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSettings", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UpdateUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSettings indicates an expected call of UpdateUserSettings.
func (mr *MockProfileServiceServerMockRecorder) UpdateUserSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSettings", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateUserSettings), arg0, arg1)
}

// MockOrgServiceClient is a mock of OrgServiceClient interface.
type MockOrgServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrgServiceClientMockRecorder
}

// MockOrgServiceClientMockRecorder is the mock recorder for MockOrgServiceClient.
type MockOrgServiceClientMockRecorder struct {
	mock *MockOrgServiceClient
}

// NewMockOrgServiceClient creates a new mock instance.
func NewMockOrgServiceClient(ctrl *gomock.Controller) *MockOrgServiceClient {
	mock := &MockOrgServiceClient{ctrl: ctrl}
	mock.recorder = &MockOrgServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrgServiceClient) EXPECT() *MockOrgServiceClientMockRecorder {
	return m.recorder
}

// AddOrgIDEConfig mocks base method.
func (m *MockOrgServiceClient) AddOrgIDEConfig(ctx context.Context, in *profilepb.AddOrgIDEConfigRequest, opts ...grpc.CallOption) (*profilepb.AddOrgIDEConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOrgIDEConfig", varargs...)
	ret0, _ := ret[0].(*profilepb.AddOrgIDEConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrgIDEConfig indicates an expected call of AddOrgIDEConfig.
func (mr *MockOrgServiceClientMockRecorder) AddOrgIDEConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrgIDEConfig", reflect.TypeOf((*MockOrgServiceClient)(nil).AddOrgIDEConfig), varargs...)
}

// CreateInviteToken mocks base method.
func (m *MockOrgServiceClient) CreateInviteToken(ctx context.Context, in *profilepb.CreateInviteTokenRequest, opts ...grpc.CallOption) (*profilepb.InviteToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInviteToken", varargs...)
	ret0, _ := ret[0].(*profilepb.InviteToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInviteToken indicates an expected call of CreateInviteToken.
func (mr *MockOrgServiceClientMockRecorder) CreateInviteToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInviteToken", reflect.TypeOf((*MockOrgServiceClient)(nil).CreateInviteToken), varargs...)
}

// CreateOrg mocks base method.
func (m *MockOrgServiceClient) CreateOrg(ctx context.Context, in *profilepb.CreateOrgRequest, opts ...grpc.CallOption) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrg", varargs...)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrg indicates an expected call of CreateOrg.
func (mr *MockOrgServiceClientMockRecorder) CreateOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockOrgServiceClient)(nil).CreateOrg), varargs...)
}

// DeleteOrgIDEConfig mocks base method.
func (m *MockOrgServiceClient) DeleteOrgIDEConfig(ctx context.Context, in *profilepb.DeleteOrgIDEConfigRequest, opts ...grpc.CallOption) (*profilepb.DeleteOrgIDEConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOrgIDEConfig", varargs...)
	ret0, _ := ret[0].(*profilepb.DeleteOrgIDEConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgIDEConfig indicates an expected call of DeleteOrgIDEConfig.
func (mr *MockOrgServiceClientMockRecorder) DeleteOrgIDEConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgIDEConfig", reflect.TypeOf((*MockOrgServiceClient)(nil).DeleteOrgIDEConfig), varargs...)
}

// GetOrg mocks base method.
func (m *MockOrgServiceClient) GetOrg(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg.
func (mr *MockOrgServiceClientMockRecorder) GetOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockOrgServiceClient)(nil).GetOrg), varargs...)
}

// GetOrgByDomain mocks base method.
func (m *MockOrgServiceClient) GetOrgByDomain(ctx context.Context, in *profilepb.GetOrgByDomainRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgByDomain", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain.
func (mr *MockOrgServiceClientMockRecorder) GetOrgByDomain(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockOrgServiceClient)(nil).GetOrgByDomain), varargs...)
}

// GetOrgByName mocks base method.
func (m *MockOrgServiceClient) GetOrgByName(ctx context.Context, in *profilepb.GetOrgByNameRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgByName", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByName indicates an expected call of GetOrgByName.
func (mr *MockOrgServiceClientMockRecorder) GetOrgByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByName", reflect.TypeOf((*MockOrgServiceClient)(nil).GetOrgByName), varargs...)
}

// GetOrgIDEConfigs mocks base method.
func (m *MockOrgServiceClient) GetOrgIDEConfigs(ctx context.Context, in *profilepb.GetOrgIDEConfigsRequest, opts ...grpc.CallOption) (*profilepb.GetOrgIDEConfigsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgIDEConfigs", varargs...)
	ret0, _ := ret[0].(*profilepb.GetOrgIDEConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgIDEConfigs indicates an expected call of GetOrgIDEConfigs.
func (mr *MockOrgServiceClientMockRecorder) GetOrgIDEConfigs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgIDEConfigs", reflect.TypeOf((*MockOrgServiceClient)(nil).GetOrgIDEConfigs), varargs...)
}

// GetOrgs mocks base method.
func (m *MockOrgServiceClient) GetOrgs(ctx context.Context, in *profilepb.GetOrgsRequest, opts ...grpc.CallOption) (*profilepb.GetOrgsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgs", varargs...)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockOrgServiceClientMockRecorder) GetOrgs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockOrgServiceClient)(nil).GetOrgs), varargs...)
}

// GetUsersInOrg mocks base method.
func (m *MockOrgServiceClient) GetUsersInOrg(ctx context.Context, in *profilepb.GetUsersInOrgRequest, opts ...grpc.CallOption) (*profilepb.GetUsersInOrgResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersInOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUsersInOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInOrg indicates an expected call of GetUsersInOrg.
func (mr *MockOrgServiceClientMockRecorder) GetUsersInOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInOrg", reflect.TypeOf((*MockOrgServiceClient)(nil).GetUsersInOrg), varargs...)
}

// RevokeAllInviteTokens mocks base method.
func (m *MockOrgServiceClient) RevokeAllInviteTokens(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeAllInviteTokens", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAllInviteTokens indicates an expected call of RevokeAllInviteTokens.
func (mr *MockOrgServiceClientMockRecorder) RevokeAllInviteTokens(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllInviteTokens", reflect.TypeOf((*MockOrgServiceClient)(nil).RevokeAllInviteTokens), varargs...)
}

// UpdateOrg mocks base method.
func (m *MockOrgServiceClient) UpdateOrg(ctx context.Context, in *profilepb.UpdateOrgRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrg indicates an expected call of UpdateOrg.
func (mr *MockOrgServiceClientMockRecorder) UpdateOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrg", reflect.TypeOf((*MockOrgServiceClient)(nil).UpdateOrg), varargs...)
}

// VerifyInviteToken mocks base method.
func (m *MockOrgServiceClient) VerifyInviteToken(ctx context.Context, in *profilepb.InviteToken, opts ...grpc.CallOption) (*profilepb.VerifyInviteTokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyInviteToken", varargs...)
	ret0, _ := ret[0].(*profilepb.VerifyInviteTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyInviteToken indicates an expected call of VerifyInviteToken.
func (mr *MockOrgServiceClientMockRecorder) VerifyInviteToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyInviteToken", reflect.TypeOf((*MockOrgServiceClient)(nil).VerifyInviteToken), varargs...)
}

// MockOrgServiceServer is a mock of OrgServiceServer interface.
type MockOrgServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockOrgServiceServerMockRecorder
}

// MockOrgServiceServerMockRecorder is the mock recorder for MockOrgServiceServer.
type MockOrgServiceServerMockRecorder struct {
	mock *MockOrgServiceServer
}

// NewMockOrgServiceServer creates a new mock instance.
func NewMockOrgServiceServer(ctrl *gomock.Controller) *MockOrgServiceServer {
	mock := &MockOrgServiceServer{ctrl: ctrl}
	mock.recorder = &MockOrgServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrgServiceServer) EXPECT() *MockOrgServiceServerMockRecorder {
	return m.recorder
}

// AddOrgIDEConfig mocks base method.
func (m *MockOrgServiceServer) AddOrgIDEConfig(arg0 context.Context, arg1 *profilepb.AddOrgIDEConfigRequest) (*profilepb.AddOrgIDEConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrgIDEConfig", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.AddOrgIDEConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrgIDEConfig indicates an expected call of AddOrgIDEConfig.
func (mr *MockOrgServiceServerMockRecorder) AddOrgIDEConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrgIDEConfig", reflect.TypeOf((*MockOrgServiceServer)(nil).AddOrgIDEConfig), arg0, arg1)
}

// CreateInviteToken mocks base method.
func (m *MockOrgServiceServer) CreateInviteToken(arg0 context.Context, arg1 *profilepb.CreateInviteTokenRequest) (*profilepb.InviteToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInviteToken", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.InviteToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInviteToken indicates an expected call of CreateInviteToken.
func (mr *MockOrgServiceServerMockRecorder) CreateInviteToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInviteToken", reflect.TypeOf((*MockOrgServiceServer)(nil).CreateInviteToken), arg0, arg1)
}

// CreateOrg mocks base method.
func (m *MockOrgServiceServer) CreateOrg(arg0 context.Context, arg1 *profilepb.CreateOrgRequest) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrg", arg0, arg1)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrg indicates an expected call of CreateOrg.
func (mr *MockOrgServiceServerMockRecorder) CreateOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockOrgServiceServer)(nil).CreateOrg), arg0, arg1)
}

// DeleteOrgIDEConfig mocks base method.
func (m *MockOrgServiceServer) DeleteOrgIDEConfig(arg0 context.Context, arg1 *profilepb.DeleteOrgIDEConfigRequest) (*profilepb.DeleteOrgIDEConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgIDEConfig", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.DeleteOrgIDEConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgIDEConfig indicates an expected call of DeleteOrgIDEConfig.
func (mr *MockOrgServiceServerMockRecorder) DeleteOrgIDEConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgIDEConfig", reflect.TypeOf((*MockOrgServiceServer)(nil).DeleteOrgIDEConfig), arg0, arg1)
}

// GetOrg mocks base method.
func (m *MockOrgServiceServer) GetOrg(arg0 context.Context, arg1 *uuidpb.UUID) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg.
func (mr *MockOrgServiceServerMockRecorder) GetOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockOrgServiceServer)(nil).GetOrg), arg0, arg1)
}

// GetOrgByDomain mocks base method.
func (m *MockOrgServiceServer) GetOrgByDomain(arg0 context.Context, arg1 *profilepb.GetOrgByDomainRequest) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgByDomain", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain.
func (mr *MockOrgServiceServerMockRecorder) GetOrgByDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockOrgServiceServer)(nil).GetOrgByDomain), arg0, arg1)
}

// GetOrgByName mocks base method.
func (m *MockOrgServiceServer) GetOrgByName(arg0 context.Context, arg1 *profilepb.GetOrgByNameRequest) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgByName", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByName indicates an expected call of GetOrgByName.
func (mr *MockOrgServiceServerMockRecorder) GetOrgByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByName", reflect.TypeOf((*MockOrgServiceServer)(nil).GetOrgByName), arg0, arg1)
}

// GetOrgIDEConfigs mocks base method.
func (m *MockOrgServiceServer) GetOrgIDEConfigs(arg0 context.Context, arg1 *profilepb.GetOrgIDEConfigsRequest) (*profilepb.GetOrgIDEConfigsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgIDEConfigs", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetOrgIDEConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgIDEConfigs indicates an expected call of GetOrgIDEConfigs.
func (mr *MockOrgServiceServerMockRecorder) GetOrgIDEConfigs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgIDEConfigs", reflect.TypeOf((*MockOrgServiceServer)(nil).GetOrgIDEConfigs), arg0, arg1)
}

// GetOrgs mocks base method.
func (m *MockOrgServiceServer) GetOrgs(arg0 context.Context, arg1 *profilepb.GetOrgsRequest) (*profilepb.GetOrgsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockOrgServiceServerMockRecorder) GetOrgs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockOrgServiceServer)(nil).GetOrgs), arg0, arg1)
}

// GetUsersInOrg mocks base method.
func (m *MockOrgServiceServer) GetUsersInOrg(arg0 context.Context, arg1 *profilepb.GetUsersInOrgRequest) (*profilepb.GetUsersInOrgResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersInOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUsersInOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInOrg indicates an expected call of GetUsersInOrg.
func (mr *MockOrgServiceServerMockRecorder) GetUsersInOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInOrg", reflect.TypeOf((*MockOrgServiceServer)(nil).GetUsersInOrg), arg0, arg1)
}

// RevokeAllInviteTokens mocks base method.
func (m *MockOrgServiceServer) RevokeAllInviteTokens(arg0 context.Context, arg1 *uuidpb.UUID) (*types.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAllInviteTokens", arg0, arg1)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAllInviteTokens indicates an expected call of RevokeAllInviteTokens.
func (mr *MockOrgServiceServerMockRecorder) RevokeAllInviteTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllInviteTokens", reflect.TypeOf((*MockOrgServiceServer)(nil).RevokeAllInviteTokens), arg0, arg1)
}

// UpdateOrg mocks base method.
func (m *MockOrgServiceServer) UpdateOrg(arg0 context.Context, arg1 *profilepb.UpdateOrgRequest) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrg indicates an expected call of UpdateOrg.
func (mr *MockOrgServiceServerMockRecorder) UpdateOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrg", reflect.TypeOf((*MockOrgServiceServer)(nil).UpdateOrg), arg0, arg1)
}

// VerifyInviteToken mocks base method.
func (m *MockOrgServiceServer) VerifyInviteToken(arg0 context.Context, arg1 *profilepb.InviteToken) (*profilepb.VerifyInviteTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyInviteToken", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.VerifyInviteTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyInviteToken indicates an expected call of VerifyInviteToken.
func (mr *MockOrgServiceServerMockRecorder) VerifyInviteToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyInviteToken", reflect.TypeOf((*MockOrgServiceServer)(nil).VerifyInviteToken), arg0, arg1)
}
