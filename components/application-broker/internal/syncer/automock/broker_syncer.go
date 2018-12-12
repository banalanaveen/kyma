// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"

// BrokerSyncer is an autogenerated mock type for the BrokerSyncer type
type BrokerSyncer struct {
	mock.Mock
}

// Sync provides a mock function with given fields: labelSelector, retries
func (_m *BrokerSyncer) Sync(labelSelector string, retries int) error {
	ret := _m.Called(labelSelector, retries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(labelSelector, retries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}