// Code generated by mockery v2.53.3. DO NOT EDIT.

package function

import (
	schemapb "github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	mock "github.com/stretchr/testify/mock"
)

// MockFunctionRunner is an autogenerated mock type for the FunctionRunner type
type MockFunctionRunner struct {
	mock.Mock
}

type MockFunctionRunner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFunctionRunner) EXPECT() *MockFunctionRunner_Expecter {
	return &MockFunctionRunner_Expecter{mock: &_m.Mock}
}

// BatchRun provides a mock function with given fields: inputs
func (_m *MockFunctionRunner) BatchRun(inputs ...interface{}) ([]interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, inputs...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchRun")
	}

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(...interface{}) ([]interface{}, error)); ok {
		return rf(inputs...)
	}
	if rf, ok := ret.Get(0).(func(...interface{}) []interface{}); ok {
		r0 = rf(inputs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(inputs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionRunner_BatchRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchRun'
type MockFunctionRunner_BatchRun_Call struct {
	*mock.Call
}

// BatchRun is a helper method to define mock.On call
//   - inputs ...interface{}
func (_e *MockFunctionRunner_Expecter) BatchRun(inputs ...interface{}) *MockFunctionRunner_BatchRun_Call {
	return &MockFunctionRunner_BatchRun_Call{Call: _e.mock.On("BatchRun",
		append([]interface{}{}, inputs...)...)}
}

func (_c *MockFunctionRunner_BatchRun_Call) Run(run func(inputs ...interface{})) *MockFunctionRunner_BatchRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockFunctionRunner_BatchRun_Call) Return(_a0 []interface{}, _a1 error) *MockFunctionRunner_BatchRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionRunner_BatchRun_Call) RunAndReturn(run func(...interface{}) ([]interface{}, error)) *MockFunctionRunner_BatchRun_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *MockFunctionRunner) Close() {
	_m.Called()
}

// MockFunctionRunner_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockFunctionRunner_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockFunctionRunner_Expecter) Close() *MockFunctionRunner_Close_Call {
	return &MockFunctionRunner_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockFunctionRunner_Close_Call) Run(run func()) *MockFunctionRunner_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFunctionRunner_Close_Call) Return() *MockFunctionRunner_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFunctionRunner_Close_Call) RunAndReturn(run func()) *MockFunctionRunner_Close_Call {
	_c.Run(run)
	return _c
}

// GetInputFields provides a mock function with no fields
func (_m *MockFunctionRunner) GetInputFields() []*schemapb.FieldSchema {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetInputFields")
	}

	var r0 []*schemapb.FieldSchema
	if rf, ok := ret.Get(0).(func() []*schemapb.FieldSchema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schemapb.FieldSchema)
		}
	}

	return r0
}

// MockFunctionRunner_GetInputFields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInputFields'
type MockFunctionRunner_GetInputFields_Call struct {
	*mock.Call
}

// GetInputFields is a helper method to define mock.On call
func (_e *MockFunctionRunner_Expecter) GetInputFields() *MockFunctionRunner_GetInputFields_Call {
	return &MockFunctionRunner_GetInputFields_Call{Call: _e.mock.On("GetInputFields")}
}

func (_c *MockFunctionRunner_GetInputFields_Call) Run(run func()) *MockFunctionRunner_GetInputFields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFunctionRunner_GetInputFields_Call) Return(_a0 []*schemapb.FieldSchema) *MockFunctionRunner_GetInputFields_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionRunner_GetInputFields_Call) RunAndReturn(run func() []*schemapb.FieldSchema) *MockFunctionRunner_GetInputFields_Call {
	_c.Call.Return(run)
	return _c
}

// GetOutputFields provides a mock function with no fields
func (_m *MockFunctionRunner) GetOutputFields() []*schemapb.FieldSchema {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOutputFields")
	}

	var r0 []*schemapb.FieldSchema
	if rf, ok := ret.Get(0).(func() []*schemapb.FieldSchema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schemapb.FieldSchema)
		}
	}

	return r0
}

// MockFunctionRunner_GetOutputFields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOutputFields'
type MockFunctionRunner_GetOutputFields_Call struct {
	*mock.Call
}

// GetOutputFields is a helper method to define mock.On call
func (_e *MockFunctionRunner_Expecter) GetOutputFields() *MockFunctionRunner_GetOutputFields_Call {
	return &MockFunctionRunner_GetOutputFields_Call{Call: _e.mock.On("GetOutputFields")}
}

func (_c *MockFunctionRunner_GetOutputFields_Call) Run(run func()) *MockFunctionRunner_GetOutputFields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFunctionRunner_GetOutputFields_Call) Return(_a0 []*schemapb.FieldSchema) *MockFunctionRunner_GetOutputFields_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionRunner_GetOutputFields_Call) RunAndReturn(run func() []*schemapb.FieldSchema) *MockFunctionRunner_GetOutputFields_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchema provides a mock function with no fields
func (_m *MockFunctionRunner) GetSchema() *schemapb.FunctionSchema {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSchema")
	}

	var r0 *schemapb.FunctionSchema
	if rf, ok := ret.Get(0).(func() *schemapb.FunctionSchema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schemapb.FunctionSchema)
		}
	}

	return r0
}

// MockFunctionRunner_GetSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchema'
type MockFunctionRunner_GetSchema_Call struct {
	*mock.Call
}

// GetSchema is a helper method to define mock.On call
func (_e *MockFunctionRunner_Expecter) GetSchema() *MockFunctionRunner_GetSchema_Call {
	return &MockFunctionRunner_GetSchema_Call{Call: _e.mock.On("GetSchema")}
}

func (_c *MockFunctionRunner_GetSchema_Call) Run(run func()) *MockFunctionRunner_GetSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFunctionRunner_GetSchema_Call) Return(_a0 *schemapb.FunctionSchema) *MockFunctionRunner_GetSchema_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionRunner_GetSchema_Call) RunAndReturn(run func() *schemapb.FunctionSchema) *MockFunctionRunner_GetSchema_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFunctionRunner creates a new instance of MockFunctionRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFunctionRunner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFunctionRunner {
	mock := &MockFunctionRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
