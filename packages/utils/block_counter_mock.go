// Code generated by mockery v1.0.0
package utils

import mock "github.com/stretchr/testify/mock"

// mockIntervalBlocksCounter is an autogenerated mock type for the intervalBlocksCounter type
type mockIntervalBlocksCounter struct {
	mock.Mock
}

// count provides a mock function with given fields: state
func (_m *mockIntervalBlocksCounter) count(state blockGenerationState) (int, error) {
	ret := _m.Called(state)

	var r0 int
	if rf, ok := ret.Get(0).(func(blockGenerationState) int); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(blockGenerationState) error); ok {
		r1 = rf(state)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
