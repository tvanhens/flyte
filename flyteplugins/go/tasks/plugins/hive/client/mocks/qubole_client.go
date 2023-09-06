// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/flyteorg/flyteplugins/go/tasks/plugins/hive/client"

	mock "github.com/stretchr/testify/mock"
)

// QuboleClient is an autogenerated mock type for the QuboleClient type
type QuboleClient struct {
	mock.Mock
}

type QuboleClient_ExecuteHiveCommand struct {
	*mock.Call
}

func (_m QuboleClient_ExecuteHiveCommand) Return(_a0 *client.QuboleCommandDetails, _a1 error) *QuboleClient_ExecuteHiveCommand {
	return &QuboleClient_ExecuteHiveCommand{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *QuboleClient) OnExecuteHiveCommand(ctx context.Context, commandStr string, timeoutVal uint32, clusterPrimaryLabel string, accountKey string, tags []string, commandMetadata client.CommandMetadata) *QuboleClient_ExecuteHiveCommand {
	c_call := _m.On("ExecuteHiveCommand", ctx, commandStr, timeoutVal, clusterPrimaryLabel, accountKey, tags, commandMetadata)
	return &QuboleClient_ExecuteHiveCommand{Call: c_call}
}

func (_m *QuboleClient) OnExecuteHiveCommandMatch(matchers ...interface{}) *QuboleClient_ExecuteHiveCommand {
	c_call := _m.On("ExecuteHiveCommand", matchers...)
	return &QuboleClient_ExecuteHiveCommand{Call: c_call}
}

// ExecuteHiveCommand provides a mock function with given fields: ctx, commandStr, timeoutVal, clusterPrimaryLabel, accountKey, tags, commandMetadata
func (_m *QuboleClient) ExecuteHiveCommand(ctx context.Context, commandStr string, timeoutVal uint32, clusterPrimaryLabel string, accountKey string, tags []string, commandMetadata client.CommandMetadata) (*client.QuboleCommandDetails, error) {
	ret := _m.Called(ctx, commandStr, timeoutVal, clusterPrimaryLabel, accountKey, tags, commandMetadata)

	var r0 *client.QuboleCommandDetails
	if rf, ok := ret.Get(0).(func(context.Context, string, uint32, string, string, []string, client.CommandMetadata) *client.QuboleCommandDetails); ok {
		r0 = rf(ctx, commandStr, timeoutVal, clusterPrimaryLabel, accountKey, tags, commandMetadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.QuboleCommandDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint32, string, string, []string, client.CommandMetadata) error); ok {
		r1 = rf(ctx, commandStr, timeoutVal, clusterPrimaryLabel, accountKey, tags, commandMetadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type QuboleClient_GetCommandStatus struct {
	*mock.Call
}

func (_m QuboleClient_GetCommandStatus) Return(_a0 client.QuboleStatus, _a1 error) *QuboleClient_GetCommandStatus {
	return &QuboleClient_GetCommandStatus{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *QuboleClient) OnGetCommandStatus(ctx context.Context, commandID string, accountKey string) *QuboleClient_GetCommandStatus {
	c_call := _m.On("GetCommandStatus", ctx, commandID, accountKey)
	return &QuboleClient_GetCommandStatus{Call: c_call}
}

func (_m *QuboleClient) OnGetCommandStatusMatch(matchers ...interface{}) *QuboleClient_GetCommandStatus {
	c_call := _m.On("GetCommandStatus", matchers...)
	return &QuboleClient_GetCommandStatus{Call: c_call}
}

// GetCommandStatus provides a mock function with given fields: ctx, commandID, accountKey
func (_m *QuboleClient) GetCommandStatus(ctx context.Context, commandID string, accountKey string) (client.QuboleStatus, error) {
	ret := _m.Called(ctx, commandID, accountKey)

	var r0 client.QuboleStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string) client.QuboleStatus); ok {
		r0 = rf(ctx, commandID, accountKey)
	} else {
		r0 = ret.Get(0).(client.QuboleStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, commandID, accountKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type QuboleClient_KillCommand struct {
	*mock.Call
}

func (_m QuboleClient_KillCommand) Return(_a0 error) *QuboleClient_KillCommand {
	return &QuboleClient_KillCommand{Call: _m.Call.Return(_a0)}
}

func (_m *QuboleClient) OnKillCommand(ctx context.Context, commandID string, accountKey string) *QuboleClient_KillCommand {
	c_call := _m.On("KillCommand", ctx, commandID, accountKey)
	return &QuboleClient_KillCommand{Call: c_call}
}

func (_m *QuboleClient) OnKillCommandMatch(matchers ...interface{}) *QuboleClient_KillCommand {
	c_call := _m.On("KillCommand", matchers...)
	return &QuboleClient_KillCommand{Call: c_call}
}

// KillCommand provides a mock function with given fields: ctx, commandID, accountKey
func (_m *QuboleClient) KillCommand(ctx context.Context, commandID string, accountKey string) error {
	ret := _m.Called(ctx, commandID, accountKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, commandID, accountKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
