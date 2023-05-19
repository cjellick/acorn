// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/acorn-io/acorn/pkg/client (interfaces: Client,ProjectClientFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/acorn-io/acorn/pkg/apis/api.acorn.io/v1"
	v10 "github.com/acorn-io/acorn/pkg/apis/internal.acorn.io/v1"
	client "github.com/acorn-io/acorn/pkg/client"
	term "github.com/acorn-io/acorn/pkg/client/term"
	gomock "github.com/golang/mock/gomock"
	client0 "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AcornImageBuild mocks base method.
func (m *MockClient) AcornImageBuild(arg0 context.Context, arg1 string, arg2 *client.AcornImageBuildOptions) (*v10.AppImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcornImageBuild", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.AppImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcornImageBuild indicates an expected call of AcornImageBuild.
func (mr *MockClientMockRecorder) AcornImageBuild(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcornImageBuild", reflect.TypeOf((*MockClient)(nil).AcornImageBuild), arg0, arg1, arg2)
}

// AcornImageBuildDelete mocks base method.
func (m *MockClient) AcornImageBuildDelete(arg0 context.Context, arg1 string) (*v1.AcornImageBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcornImageBuildDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.AcornImageBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcornImageBuildDelete indicates an expected call of AcornImageBuildDelete.
func (mr *MockClientMockRecorder) AcornImageBuildDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcornImageBuildDelete", reflect.TypeOf((*MockClient)(nil).AcornImageBuildDelete), arg0, arg1)
}

// AcornImageBuildGet mocks base method.
func (m *MockClient) AcornImageBuildGet(arg0 context.Context, arg1 string) (*v1.AcornImageBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcornImageBuildGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.AcornImageBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcornImageBuildGet indicates an expected call of AcornImageBuildGet.
func (mr *MockClientMockRecorder) AcornImageBuildGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcornImageBuildGet", reflect.TypeOf((*MockClient)(nil).AcornImageBuildGet), arg0, arg1)
}

// AcornImageBuildList mocks base method.
func (m *MockClient) AcornImageBuildList(arg0 context.Context) ([]v1.AcornImageBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcornImageBuildList", arg0)
	ret0, _ := ret[0].([]v1.AcornImageBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcornImageBuildList indicates an expected call of AcornImageBuildList.
func (mr *MockClientMockRecorder) AcornImageBuildList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcornImageBuildList", reflect.TypeOf((*MockClient)(nil).AcornImageBuildList), arg0)
}

// AppConfirmUpgrade mocks base method.
func (m *MockClient) AppConfirmUpgrade(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppConfirmUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppConfirmUpgrade indicates an expected call of AppConfirmUpgrade.
func (mr *MockClientMockRecorder) AppConfirmUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppConfirmUpgrade", reflect.TypeOf((*MockClient)(nil).AppConfirmUpgrade), arg0, arg1)
}

// AppDelete mocks base method.
func (m *MockClient) AppDelete(arg0 context.Context, arg1 string) (*v1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppDelete indicates an expected call of AppDelete.
func (mr *MockClientMockRecorder) AppDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppDelete", reflect.TypeOf((*MockClient)(nil).AppDelete), arg0, arg1)
}

// AppGet mocks base method.
func (m *MockClient) AppGet(arg0 context.Context, arg1 string) (*v1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppGet indicates an expected call of AppGet.
func (mr *MockClientMockRecorder) AppGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGet", reflect.TypeOf((*MockClient)(nil).AppGet), arg0, arg1)
}

// AppList mocks base method.
func (m *MockClient) AppList(arg0 context.Context) ([]v1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppList", arg0)
	ret0, _ := ret[0].([]v1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppList indicates an expected call of AppList.
func (mr *MockClientMockRecorder) AppList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppList", reflect.TypeOf((*MockClient)(nil).AppList), arg0)
}

// AppLog mocks base method.
func (m *MockClient) AppLog(arg0 context.Context, arg1 string, arg2 *client.LogOptions) (<-chan v1.LogMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppLog", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan v1.LogMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppLog indicates an expected call of AppLog.
func (mr *MockClientMockRecorder) AppLog(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppLog", reflect.TypeOf((*MockClient)(nil).AppLog), arg0, arg1, arg2)
}

// AppPullImage mocks base method.
func (m *MockClient) AppPullImage(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppPullImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppPullImage indicates an expected call of AppPullImage.
func (mr *MockClientMockRecorder) AppPullImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppPullImage", reflect.TypeOf((*MockClient)(nil).AppPullImage), arg0, arg1)
}

// AppRun mocks base method.
func (m *MockClient) AppRun(arg0 context.Context, arg1 string, arg2 *client.AppRunOptions) (*v1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppRun indicates an expected call of AppRun.
func (mr *MockClientMockRecorder) AppRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRun", reflect.TypeOf((*MockClient)(nil).AppRun), arg0, arg1, arg2)
}

// AppStart mocks base method.
func (m *MockClient) AppStart(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppStart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppStart indicates an expected call of AppStart.
func (mr *MockClientMockRecorder) AppStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppStart", reflect.TypeOf((*MockClient)(nil).AppStart), arg0, arg1)
}

// AppStop mocks base method.
func (m *MockClient) AppStop(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppStop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppStop indicates an expected call of AppStop.
func (mr *MockClientMockRecorder) AppStop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppStop", reflect.TypeOf((*MockClient)(nil).AppStop), arg0, arg1)
}

// AppUpdate mocks base method.
func (m *MockClient) AppUpdate(arg0 context.Context, arg1 string, arg2 *client.AppUpdateOptions) (*v1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppUpdate indicates an expected call of AppUpdate.
func (mr *MockClientMockRecorder) AppUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppUpdate", reflect.TypeOf((*MockClient)(nil).AppUpdate), arg0, arg1, arg2)
}

// ComputeClassGet mocks base method.
func (m *MockClient) ComputeClassGet(arg0 context.Context, arg1 string) (*v1.ComputeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeClassGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.ComputeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeClassGet indicates an expected call of ComputeClassGet.
func (mr *MockClientMockRecorder) ComputeClassGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeClassGet", reflect.TypeOf((*MockClient)(nil).ComputeClassGet), arg0, arg1)
}

// ComputeClassList mocks base method.
func (m *MockClient) ComputeClassList(arg0 context.Context) ([]v1.ComputeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeClassList", arg0)
	ret0, _ := ret[0].([]v1.ComputeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeClassList indicates an expected call of ComputeClassList.
func (mr *MockClientMockRecorder) ComputeClassList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeClassList", reflect.TypeOf((*MockClient)(nil).ComputeClassList), arg0)
}

// ContainerReplicaDelete mocks base method.
func (m *MockClient) ContainerReplicaDelete(arg0 context.Context, arg1 string) (*v1.ContainerReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerReplicaDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.ContainerReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerReplicaDelete indicates an expected call of ContainerReplicaDelete.
func (mr *MockClientMockRecorder) ContainerReplicaDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerReplicaDelete", reflect.TypeOf((*MockClient)(nil).ContainerReplicaDelete), arg0, arg1)
}

// ContainerReplicaExec mocks base method.
func (m *MockClient) ContainerReplicaExec(arg0 context.Context, arg1 string, arg2 []string, arg3 bool, arg4 *client.ContainerReplicaExecOptions) (*term.ExecIO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerReplicaExec", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*term.ExecIO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerReplicaExec indicates an expected call of ContainerReplicaExec.
func (mr *MockClientMockRecorder) ContainerReplicaExec(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerReplicaExec", reflect.TypeOf((*MockClient)(nil).ContainerReplicaExec), arg0, arg1, arg2, arg3, arg4)
}

// ContainerReplicaGet mocks base method.
func (m *MockClient) ContainerReplicaGet(arg0 context.Context, arg1 string) (*v1.ContainerReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerReplicaGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.ContainerReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerReplicaGet indicates an expected call of ContainerReplicaGet.
func (mr *MockClientMockRecorder) ContainerReplicaGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerReplicaGet", reflect.TypeOf((*MockClient)(nil).ContainerReplicaGet), arg0, arg1)
}

// ContainerReplicaList mocks base method.
func (m *MockClient) ContainerReplicaList(arg0 context.Context, arg1 *client.ContainerReplicaListOptions) ([]v1.ContainerReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerReplicaList", arg0, arg1)
	ret0, _ := ret[0].([]v1.ContainerReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerReplicaList indicates an expected call of ContainerReplicaList.
func (mr *MockClientMockRecorder) ContainerReplicaList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerReplicaList", reflect.TypeOf((*MockClient)(nil).ContainerReplicaList), arg0, arg1)
}

// ContainerReplicaPortForward mocks base method.
func (m *MockClient) ContainerReplicaPortForward(arg0 context.Context, arg1 string, arg2 int) (client.PortForwardDialer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerReplicaPortForward", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.PortForwardDialer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerReplicaPortForward indicates an expected call of ContainerReplicaPortForward.
func (mr *MockClientMockRecorder) ContainerReplicaPortForward(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerReplicaPortForward", reflect.TypeOf((*MockClient)(nil).ContainerReplicaPortForward), arg0, arg1, arg2)
}

// CredentialCreate mocks base method.
func (m *MockClient) CredentialCreate(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) (*v1.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialCreate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*v1.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialCreate indicates an expected call of CredentialCreate.
func (mr *MockClientMockRecorder) CredentialCreate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialCreate", reflect.TypeOf((*MockClient)(nil).CredentialCreate), arg0, arg1, arg2, arg3, arg4)
}

// CredentialDelete mocks base method.
func (m *MockClient) CredentialDelete(arg0 context.Context, arg1 string) (*v1.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialDelete indicates an expected call of CredentialDelete.
func (mr *MockClientMockRecorder) CredentialDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialDelete", reflect.TypeOf((*MockClient)(nil).CredentialDelete), arg0, arg1)
}

// CredentialGet mocks base method.
func (m *MockClient) CredentialGet(arg0 context.Context, arg1 string) (*v1.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialGet indicates an expected call of CredentialGet.
func (mr *MockClientMockRecorder) CredentialGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialGet", reflect.TypeOf((*MockClient)(nil).CredentialGet), arg0, arg1)
}

// CredentialList mocks base method.
func (m *MockClient) CredentialList(arg0 context.Context) ([]v1.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialList", arg0)
	ret0, _ := ret[0].([]v1.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialList indicates an expected call of CredentialList.
func (mr *MockClientMockRecorder) CredentialList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialList", reflect.TypeOf((*MockClient)(nil).CredentialList), arg0)
}

// CredentialUpdate mocks base method.
func (m *MockClient) CredentialUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) (*v1.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialUpdate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*v1.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialUpdate indicates an expected call of CredentialUpdate.
func (mr *MockClientMockRecorder) CredentialUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialUpdate", reflect.TypeOf((*MockClient)(nil).CredentialUpdate), arg0, arg1, arg2, arg3, arg4)
}

// EventStream mocks base method.
func (m *MockClient) EventStream(arg0 context.Context, arg1 *client.EventStreamOptions) (<-chan v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventStream", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventStream indicates an expected call of EventStream.
func (mr *MockClientMockRecorder) EventStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventStream", reflect.TypeOf((*MockClient)(nil).EventStream), arg0, arg1)
}

// GetClient mocks base method.
func (m *MockClient) GetClient() (client0.WithWatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(client0.WithWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient.
func (mr *MockClientMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockClient)(nil).GetClient))
}

// GetNamespace mocks base method.
func (m *MockClient) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockClientMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockClient)(nil).GetNamespace))
}

// GetProject mocks base method.
func (m *MockClient) GetProject() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProject indicates an expected call of GetProject.
func (mr *MockClientMockRecorder) GetProject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockClient)(nil).GetProject))
}

// ImageDelete mocks base method.
func (m *MockClient) ImageDelete(arg0 context.Context, arg1 string, arg2 *client.ImageDeleteOptions) (*v1.Image, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Image)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImageDelete indicates an expected call of ImageDelete.
func (mr *MockClientMockRecorder) ImageDelete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageDelete", reflect.TypeOf((*MockClient)(nil).ImageDelete), arg0, arg1, arg2)
}

// ImageDetails mocks base method.
func (m *MockClient) ImageDetails(arg0 context.Context, arg1 string, arg2 *client.ImageDetailsOptions) (*client.ImageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageDetails", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.ImageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageDetails indicates an expected call of ImageDetails.
func (mr *MockClientMockRecorder) ImageDetails(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageDetails", reflect.TypeOf((*MockClient)(nil).ImageDetails), arg0, arg1, arg2)
}

// ImageGet mocks base method.
func (m *MockClient) ImageGet(arg0 context.Context, arg1 string) (*v1.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageGet indicates an expected call of ImageGet.
func (mr *MockClientMockRecorder) ImageGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageGet", reflect.TypeOf((*MockClient)(nil).ImageGet), arg0, arg1)
}

// ImageList mocks base method.
func (m *MockClient) ImageList(arg0 context.Context) ([]v1.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageList", arg0)
	ret0, _ := ret[0].([]v1.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageList indicates an expected call of ImageList.
func (mr *MockClientMockRecorder) ImageList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageList", reflect.TypeOf((*MockClient)(nil).ImageList), arg0)
}

// ImagePull mocks base method.
func (m *MockClient) ImagePull(arg0 context.Context, arg1 string, arg2 *client.ImagePullOptions) (<-chan client.ImageProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePull", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.ImageProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagePull indicates an expected call of ImagePull.
func (mr *MockClientMockRecorder) ImagePull(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePull", reflect.TypeOf((*MockClient)(nil).ImagePull), arg0, arg1, arg2)
}

// ImagePush mocks base method.
func (m *MockClient) ImagePush(arg0 context.Context, arg1 string, arg2 *client.ImagePushOptions) (<-chan client.ImageProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePush", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.ImageProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagePush indicates an expected call of ImagePush.
func (mr *MockClientMockRecorder) ImagePush(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePush", reflect.TypeOf((*MockClient)(nil).ImagePush), arg0, arg1, arg2)
}

// ImageTag mocks base method.
func (m *MockClient) ImageTag(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageTag", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageTag indicates an expected call of ImageTag.
func (mr *MockClientMockRecorder) ImageTag(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageTag", reflect.TypeOf((*MockClient)(nil).ImageTag), arg0, arg1, arg2)
}

// Info mocks base method.
func (m *MockClient) Info(arg0 context.Context) ([]v1.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].([]v1.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockClientMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockClient)(nil).Info), arg0)
}

// ProjectCreate mocks base method.
func (m *MockClient) ProjectCreate(arg0 context.Context, arg1, arg2 string, arg3 []string) (*v1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectCreate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectCreate indicates an expected call of ProjectCreate.
func (mr *MockClientMockRecorder) ProjectCreate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectCreate", reflect.TypeOf((*MockClient)(nil).ProjectCreate), arg0, arg1, arg2, arg3)
}

// ProjectDelete mocks base method.
func (m *MockClient) ProjectDelete(arg0 context.Context, arg1 string) (*v1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectDelete indicates an expected call of ProjectDelete.
func (mr *MockClientMockRecorder) ProjectDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectDelete", reflect.TypeOf((*MockClient)(nil).ProjectDelete), arg0, arg1)
}

// ProjectGet mocks base method.
func (m *MockClient) ProjectGet(arg0 context.Context, arg1 string) (*v1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectGet indicates an expected call of ProjectGet.
func (mr *MockClientMockRecorder) ProjectGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectGet", reflect.TypeOf((*MockClient)(nil).ProjectGet), arg0, arg1)
}

// ProjectList mocks base method.
func (m *MockClient) ProjectList(arg0 context.Context) ([]v1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectList", arg0)
	ret0, _ := ret[0].([]v1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectList indicates an expected call of ProjectList.
func (mr *MockClientMockRecorder) ProjectList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectList", reflect.TypeOf((*MockClient)(nil).ProjectList), arg0)
}

// ProjectUpdate mocks base method.
func (m *MockClient) ProjectUpdate(arg0 context.Context, arg1 *v1.Project, arg2 string, arg3 []string) (*v1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdate indicates an expected call of ProjectUpdate.
func (mr *MockClientMockRecorder) ProjectUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdate", reflect.TypeOf((*MockClient)(nil).ProjectUpdate), arg0, arg1, arg2, arg3)
}

// RegionGet mocks base method.
func (m *MockClient) RegionGet(arg0 context.Context, arg1 string) (*v1.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegionGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegionGet indicates an expected call of RegionGet.
func (mr *MockClientMockRecorder) RegionGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegionGet", reflect.TypeOf((*MockClient)(nil).RegionGet), arg0, arg1)
}

// RegionList mocks base method.
func (m *MockClient) RegionList(arg0 context.Context) ([]v1.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegionList", arg0)
	ret0, _ := ret[0].([]v1.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegionList indicates an expected call of RegionList.
func (mr *MockClientMockRecorder) RegionList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegionList", reflect.TypeOf((*MockClient)(nil).RegionList), arg0)
}

// SecretCreate mocks base method.
func (m *MockClient) SecretCreate(arg0 context.Context, arg1, arg2 string, arg3 map[string][]byte) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretCreate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretCreate indicates an expected call of SecretCreate.
func (mr *MockClientMockRecorder) SecretCreate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretCreate", reflect.TypeOf((*MockClient)(nil).SecretCreate), arg0, arg1, arg2, arg3)
}

// SecretDelete mocks base method.
func (m *MockClient) SecretDelete(arg0 context.Context, arg1 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretDelete indicates an expected call of SecretDelete.
func (mr *MockClientMockRecorder) SecretDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretDelete", reflect.TypeOf((*MockClient)(nil).SecretDelete), arg0, arg1)
}

// SecretGet mocks base method.
func (m *MockClient) SecretGet(arg0 context.Context, arg1 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGet indicates an expected call of SecretGet.
func (mr *MockClientMockRecorder) SecretGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGet", reflect.TypeOf((*MockClient)(nil).SecretGet), arg0, arg1)
}

// SecretList mocks base method.
func (m *MockClient) SecretList(arg0 context.Context) ([]v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretList", arg0)
	ret0, _ := ret[0].([]v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretList indicates an expected call of SecretList.
func (mr *MockClientMockRecorder) SecretList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretList", reflect.TypeOf((*MockClient)(nil).SecretList), arg0)
}

// SecretReveal mocks base method.
func (m *MockClient) SecretReveal(arg0 context.Context, arg1 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretReveal", arg0, arg1)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretReveal indicates an expected call of SecretReveal.
func (mr *MockClientMockRecorder) SecretReveal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretReveal", reflect.TypeOf((*MockClient)(nil).SecretReveal), arg0, arg1)
}

// SecretUpdate mocks base method.
func (m *MockClient) SecretUpdate(arg0 context.Context, arg1 string, arg2 map[string][]byte) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretUpdate indicates an expected call of SecretUpdate.
func (mr *MockClientMockRecorder) SecretUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretUpdate", reflect.TypeOf((*MockClient)(nil).SecretUpdate), arg0, arg1, arg2)
}

// VolumeClassGet mocks base method.
func (m *MockClient) VolumeClassGet(arg0 context.Context, arg1 string) (*v1.VolumeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeClassGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.VolumeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeClassGet indicates an expected call of VolumeClassGet.
func (mr *MockClientMockRecorder) VolumeClassGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeClassGet", reflect.TypeOf((*MockClient)(nil).VolumeClassGet), arg0, arg1)
}

// VolumeClassList mocks base method.
func (m *MockClient) VolumeClassList(arg0 context.Context) ([]v1.VolumeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeClassList", arg0)
	ret0, _ := ret[0].([]v1.VolumeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeClassList indicates an expected call of VolumeClassList.
func (mr *MockClientMockRecorder) VolumeClassList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeClassList", reflect.TypeOf((*MockClient)(nil).VolumeClassList), arg0)
}

// VolumeDelete mocks base method.
func (m *MockClient) VolumeDelete(arg0 context.Context, arg1 string) (*v1.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeDelete", arg0, arg1)
	ret0, _ := ret[0].(*v1.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeDelete indicates an expected call of VolumeDelete.
func (mr *MockClientMockRecorder) VolumeDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeDelete", reflect.TypeOf((*MockClient)(nil).VolumeDelete), arg0, arg1)
}

// VolumeGet mocks base method.
func (m *MockClient) VolumeGet(arg0 context.Context, arg1 string) (*v1.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeGet", arg0, arg1)
	ret0, _ := ret[0].(*v1.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeGet indicates an expected call of VolumeGet.
func (mr *MockClientMockRecorder) VolumeGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeGet", reflect.TypeOf((*MockClient)(nil).VolumeGet), arg0, arg1)
}

// VolumeList mocks base method.
func (m *MockClient) VolumeList(arg0 context.Context) ([]v1.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeList", arg0)
	ret0, _ := ret[0].([]v1.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeList indicates an expected call of VolumeList.
func (mr *MockClientMockRecorder) VolumeList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeList", reflect.TypeOf((*MockClient)(nil).VolumeList), arg0)
}

// MockProjectClientFactory is a mock of ProjectClientFactory interface.
type MockProjectClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockProjectClientFactoryMockRecorder
}

// MockProjectClientFactoryMockRecorder is the mock recorder for MockProjectClientFactory.
type MockProjectClientFactoryMockRecorder struct {
	mock *MockProjectClientFactory
}

// NewMockProjectClientFactory creates a new mock instance.
func NewMockProjectClientFactory(ctrl *gomock.Controller) *MockProjectClientFactory {
	mock := &MockProjectClientFactory{ctrl: ctrl}
	mock.recorder = &MockProjectClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectClientFactory) EXPECT() *MockProjectClientFactoryMockRecorder {
	return m.recorder
}

// DefaultProject mocks base method.
func (m *MockProjectClientFactory) DefaultProject() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultProject")
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultProject indicates an expected call of DefaultProject.
func (mr *MockProjectClientFactoryMockRecorder) DefaultProject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultProject", reflect.TypeOf((*MockProjectClientFactory)(nil).DefaultProject))
}

// ForProject mocks base method.
func (m *MockProjectClientFactory) ForProject(arg0 context.Context, arg1 string) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForProject", arg0, arg1)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForProject indicates an expected call of ForProject.
func (mr *MockProjectClientFactoryMockRecorder) ForProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForProject", reflect.TypeOf((*MockProjectClientFactory)(nil).ForProject), arg0, arg1)
}

// List mocks base method.
func (m *MockProjectClientFactory) List(arg0 context.Context) ([]client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectClientFactoryMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectClientFactory)(nil).List), arg0)
}
