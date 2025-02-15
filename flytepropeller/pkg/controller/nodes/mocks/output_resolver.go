// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	executors "github.com/flyteorg/flyte/flytepropeller/pkg/controller/executors"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// OutputResolver is an autogenerated mock type for the OutputResolver type
type OutputResolver struct {
	mock.Mock
}

type OutputResolver_ExtractOutput struct {
	*mock.Call
}

func (_m OutputResolver_ExtractOutput) Return(values *core.Literal, err error) *OutputResolver_ExtractOutput {
	return &OutputResolver_ExtractOutput{Call: _m.Call.Return(values, err)}
}

func (_m *OutputResolver) OnExtractOutput(ctx context.Context, nl executors.NodeLookup, n v1alpha1.ExecutableNode, bindToVar string) *OutputResolver_ExtractOutput {
	c_call := _m.On("ExtractOutput", ctx, nl, n, bindToVar)
	return &OutputResolver_ExtractOutput{Call: c_call}
}

func (_m *OutputResolver) OnExtractOutputMatch(matchers ...interface{}) *OutputResolver_ExtractOutput {
	c_call := _m.On("ExtractOutput", matchers...)
	return &OutputResolver_ExtractOutput{Call: c_call}
}

// ExtractOutput provides a mock function with given fields: ctx, nl, n, bindToVar
func (_m *OutputResolver) ExtractOutput(ctx context.Context, nl executors.NodeLookup, n v1alpha1.ExecutableNode, bindToVar string) (*core.Literal, error) {
	ret := _m.Called(ctx, nl, n, bindToVar)

	var r0 *core.Literal
	if rf, ok := ret.Get(0).(func(context.Context, executors.NodeLookup, v1alpha1.ExecutableNode, string) *core.Literal); ok {
		r0 = rf(ctx, nl, n, bindToVar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Literal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, executors.NodeLookup, v1alpha1.ExecutableNode, string) error); ok {
		r1 = rf(ctx, nl, n, bindToVar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
