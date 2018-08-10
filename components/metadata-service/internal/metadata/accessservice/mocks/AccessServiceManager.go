// Code generated by mockery v1.0.0
package mocks

import apperrors "github.com/kyma-project/kyma/components/metadata-service/internal/apperrors"
import mock "github.com/stretchr/testify/mock"

// AccessServiceManager is an autogenerated mock type for the AccessServiceManager type
type AccessServiceManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: remoteEnvironment, serviceId, serviceName
func (_m *AccessServiceManager) Create(remoteEnvironment string, serviceId string, serviceName string) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, serviceId, serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, serviceId, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: serviceName
func (_m *AccessServiceManager) Delete(serviceName string) apperrors.AppError {
	ret := _m.Called(serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Upsert provides a mock function with given fields: remoteEnvironment, serviceId, serviceName
func (_m *AccessServiceManager) Upsert(remoteEnvironment string, serviceId string, serviceName string) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, serviceId, serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, serviceId, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}