// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by MockGen. DO NOT EDIT.
// Source: virtualmachineclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_virtualmachineclient -source virtualmachineclient/interface.go
//

// Package mock_virtualmachineclient is a generated GoMock package.
package mock_virtualmachineclient

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// BeginAttachDetachDataDisks mocks base method.
func (m *MockInterface) BeginAttachDetachDataDisks(ctx context.Context, resourceGroupName, vmName string, parameters armcompute.AttachDetachDataDisksRequest, options *armcompute.VirtualMachinesClientBeginAttachDetachDataDisksOptions) (*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginAttachDetachDataDisks", ctx, resourceGroupName, vmName, parameters, options)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginAttachDetachDataDisks indicates an expected call of BeginAttachDetachDataDisks.
func (mr *MockInterfaceMockRecorder) BeginAttachDetachDataDisks(ctx, resourceGroupName, vmName, parameters, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginAttachDetachDataDisks", reflect.TypeOf((*MockInterface)(nil).BeginAttachDetachDataDisks), ctx, resourceGroupName, vmName, parameters, options)
}

// BeginCreateOrUpdate mocks base method.
func (m *MockInterface) BeginCreateOrUpdate(ctx context.Context, resourceGroupName, vmName string, parameters armcompute.VirtualMachine, options *armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginCreateOrUpdate", ctx, resourceGroupName, vmName, parameters, options)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginCreateOrUpdate indicates an expected call of BeginCreateOrUpdate.
func (mr *MockInterfaceMockRecorder) BeginCreateOrUpdate(ctx, resourceGroupName, vmName, parameters, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginCreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).BeginCreateOrUpdate), ctx, resourceGroupName, vmName, parameters, options)
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, resourceParam armcompute.VirtualMachine) (*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, resourceParam)
	ret0, _ := ret[0].(*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, resourceParam any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, resourceParam)
}

// Delete mocks base method.
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, resourceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, resourceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, resourceName)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string, expand *string) (*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, expand)
	ret0, _ := ret[0].(*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName, expand any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName, expand)
}

// InstanceView mocks base method.
func (m *MockInterface) InstanceView(ctx context.Context, resourceGroupName, vmName string) (*armcompute.VirtualMachineInstanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceView", ctx, resourceGroupName, vmName)
	ret0, _ := ret[0].(*armcompute.VirtualMachineInstanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceView indicates an expected call of InstanceView.
func (mr *MockInterfaceMockRecorder) InstanceView(ctx, resourceGroupName, vmName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceView", reflect.TypeOf((*MockInterface)(nil).InstanceView), ctx, resourceGroupName, vmName)
}

// List mocks base method.
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
}

// ListVMInstanceView mocks base method.
func (m *MockInterface) ListVMInstanceView(ctx context.Context, resourceGroupName string) ([]*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVMInstanceView", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVMInstanceView indicates an expected call of ListVMInstanceView.
func (mr *MockInterfaceMockRecorder) ListVMInstanceView(ctx, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVMInstanceView", reflect.TypeOf((*MockInterface)(nil).ListVMInstanceView), ctx, resourceGroupName)
}
