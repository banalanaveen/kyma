// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/apps/v1"

// statefulSetLister is an autogenerated mock type for the statefulSetLister type
type statefulSetLister struct {
	mock.Mock
}

// ListStatefulSets provides a mock function with given fields: environment
func (_m *statefulSetLister) ListStatefulSets(environment string) ([]*v1.StatefulSet, error) {
	ret := _m.Called(environment)

	var r0 []*v1.StatefulSet
	if rf, ok := ret.Get(0).(func(string) []*v1.StatefulSet); ok {
		r0 = rf(environment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.StatefulSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(environment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}