/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azurefile

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	volume "github.com/kata-containers/kata-containers/src/runtime/pkg/direct-volume"
)

// MockDirectVolume is a mock of DirectVolume interface.
type MockDirectVolume struct {
	ctrl     *gomock.Controller
	recorder *MockDirectVolumeMockRecorder
}

// MockDirectVolumeMockRecorder is the mock recorder for MockDirectVolume.
type MockDirectVolumeMockRecorder struct {
	mock *MockDirectVolume
}

// NewMockDirectVolume creates a new mock instance.
func NewMockDirectVolume(ctrl *gomock.Controller) *MockDirectVolume {
	mock := &MockDirectVolume{ctrl: ctrl}
	mock.recorder = &MockDirectVolumeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectVolume) EXPECT() *MockDirectVolumeMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockDirectVolume) Add(volumePath, mountInfo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", volumePath, mountInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockDirectVolumeMockRecorder) Add(volumePath, mountInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDirectVolume)(nil).Add), volumePath, mountInfo)
}

// GetSandboxIdForVolume mocks base method.
func (m *MockDirectVolume) GetSandboxIdForVolume(volumePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxIdForVolume", volumePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxIdForVolume indicates an expected call of GetSandboxIdForVolume.
func (mr *MockDirectVolumeMockRecorder) GetSandboxIdForVolume(volumePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxIdForVolume", reflect.TypeOf((*MockDirectVolume)(nil).GetSandboxIdForVolume), volumePath)
}

// RecordSandboxId mocks base method.
func (m *MockDirectVolume) RecordSandboxId(sandboxId, volumePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordSandboxId", sandboxId, volumePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordSandboxId indicates an expected call of RecordSandboxId.
func (mr *MockDirectVolumeMockRecorder) RecordSandboxId(sandboxId, volumePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordSandboxId", reflect.TypeOf((*MockDirectVolume)(nil).RecordSandboxId), sandboxId, volumePath)
}

// Remove mocks base method.
func (m *MockDirectVolume) Remove(volumePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", volumePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockDirectVolumeMockRecorder) Remove(volumePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDirectVolume)(nil).Remove), volumePath)
}

// VolumeMountInfo mocks base method.
func (m *MockDirectVolume) VolumeMountInfo(volumePath string) (*volume.MountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeMountInfo", volumePath)
	ret0, _ := ret[0].(*volume.MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeMountInfo indicates an expected call of VolumeMountInfo.
func (mr *MockDirectVolumeMockRecorder) VolumeMountInfo(volumePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeMountInfo", reflect.TypeOf((*MockDirectVolume)(nil).VolumeMountInfo), volumePath)
}
