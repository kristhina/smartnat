/* Copyright 2019 DevFactory FZ LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import net "net"
import v1 "k8s.io/api/core/v1"
import v1alpha1 "github.com/DevFactory/smartnat/pkg/apis/smartnat/v1alpha1"

// Scrubber is an autogenerated mock type for the Scrubber type
type Scrubber struct {
	mock.Mock
}

// ScrubMapping provides a mock function with given fields: sn
func (_m *Scrubber) ScrubMapping(sn *v1alpha1.Mapping) (bool, bool, *net.IP) {
	ret := _m.Called(sn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*v1alpha1.Mapping) bool); ok {
		r0 = rf(sn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*v1alpha1.Mapping) bool); ok {
		r1 = rf(sn)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 *net.IP
	if rf, ok := ret.Get(2).(func(*v1alpha1.Mapping) *net.IP); ok {
		r2 = rf(sn)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*net.IP)
		}
	}

	return r0, r1, r2
}

// ValidateEndpoints provides a mock function with given fields: _a0, endpoints
func (_m *Scrubber) ValidateEndpoints(_a0 *v1alpha1.Mapping, endpoints *v1.Endpoints) error {
	ret := _m.Called(_a0, endpoints)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Mapping, *v1.Endpoints) error); ok {
		r0 = rf(_a0, endpoints)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
