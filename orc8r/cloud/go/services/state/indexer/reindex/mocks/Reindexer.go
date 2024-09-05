// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"magma/orc8r/cloud/go/services/state/indexer"
)

// Reindexer is an autogenerated mock type for the Reindexer type
type Reindexer struct {
	mock.Mock
}

// GetIndexerVersions provides a mock function with given fields:
func (_m *Reindexer) GetIndexerVersions() ([]*indexer.Versions, error) {
	ret := _m.Called()

	var r0 []*indexer.Versions
	if rf, ok := ret.Get(0).(func() []*indexer.Versions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*indexer.Versions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Run provides a mock function with given fields: ctx
func (_m *Reindexer) Run(ctx context.Context) {
	_m.Called(ctx)
}

// RunUnsafe provides a mock function with given fields: ctx, indexerID, sendUpdate
func (_m *Reindexer) RunUnsafe(ctx context.Context, indexerID string, sendUpdate func(string)) error {
	ret := _m.Called(ctx, indexerID, sendUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(string)) error); ok {
		r0 = rf(ctx, indexerID, sendUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
