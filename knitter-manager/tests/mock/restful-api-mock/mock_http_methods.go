// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ZTE/Knitter/pkg/noauth_openstack (interfaces: HttpMethods)

/*
Copyright 2018 ZTE Corporation. All rights reserved.
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

package mocknoauthopenstack

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of HttpMethods interface
type MockHTTPMethods struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPMethodsRecorder
}

// Recorder for MockHTTPMethods (not exported)
type MockHTTPMethodsRecorder struct {
	mock *MockHTTPMethods
}

func NewMockHTTPMethods(ctrl *gomock.Controller) *MockHTTPMethods {
	mock := &MockHTTPMethods{ctrl: ctrl}
	mock.recorder = &MockHTTPMethodsRecorder{mock}
	return mock
}

func (_m *MockHTTPMethods) EXPECT() *MockHTTPMethodsRecorder {
	return _m.recorder
}

func (_m *MockHTTPMethods) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *MockHTTPMethodsRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockHTTPMethods) Get(_param0 string) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *MockHTTPMethodsRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockHTTPMethods) Post(_param0 string, _param1 map[string]interface{}) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Post", _param0, _param1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *MockHTTPMethodsRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Post", arg0, arg1)
}
