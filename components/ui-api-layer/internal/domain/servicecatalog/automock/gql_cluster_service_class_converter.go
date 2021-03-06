// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// gqlClusterServiceClassConverter is an autogenerated mock type for the gqlClusterServiceClassConverter type
type gqlClusterServiceClassConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlClusterServiceClassConverter) ToGQL(in *v1beta1.ClusterServiceClass) (*gqlschema.ClusterServiceClass, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ClusterServiceClass
	if rf, ok := ret.Get(0).(func(*v1beta1.ClusterServiceClass) *gqlschema.ClusterServiceClass); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ClusterServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ClusterServiceClass) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlClusterServiceClassConverter) ToGQLs(in []*v1beta1.ClusterServiceClass) ([]gqlschema.ClusterServiceClass, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ClusterServiceClass
	if rf, ok := ret.Get(0).(func([]*v1beta1.ClusterServiceClass) []gqlschema.ClusterServiceClass); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ClusterServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1beta1.ClusterServiceClass) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
