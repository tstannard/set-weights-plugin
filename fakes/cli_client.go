// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
)

type CliClient struct {
	GetAppGUIDStub        func(plugin.CliConnection, string) (string, error)
	getAppGUIDMutex       sync.RWMutex
	getAppGUIDArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
	}
	getAppGUIDReturns struct {
		result1 string
		result2 error
	}
	getAppGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetDomainGUIDStub        func(plugin.CliConnection, string) (string, error)
	getDomainGUIDMutex       sync.RWMutex
	getDomainGUIDArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
	}
	getDomainGUIDReturns struct {
		result1 string
		result2 error
	}
	getDomainGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRouteGUIDStub        func(plugin.CliConnection, string, string) (string, error)
	getRouteGUIDMutex       sync.RWMutex
	getRouteGUIDArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 string
	}
	getRouteGUIDReturns struct {
		result1 string
		result2 error
	}
	getRouteGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRouteMappingGUIDStub        func(plugin.CliConnection, string, string) (string, error)
	getRouteMappingGUIDMutex       sync.RWMutex
	getRouteMappingGUIDArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 string
	}
	getRouteMappingGUIDReturns struct {
		result1 string
		result2 error
	}
	getRouteMappingGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SetRouteMappingWeightStub        func(plugin.CliConnection, string, int) error
	setRouteMappingWeightMutex       sync.RWMutex
	setRouteMappingWeightArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 int
	}
	setRouteMappingWeightReturns struct {
		result1 error
	}
	setRouteMappingWeightReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CliClient) GetAppGUID(arg1 plugin.CliConnection, arg2 string) (string, error) {
	fake.getAppGUIDMutex.Lock()
	ret, specificReturn := fake.getAppGUIDReturnsOnCall[len(fake.getAppGUIDArgsForCall)]
	fake.getAppGUIDArgsForCall = append(fake.getAppGUIDArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetAppGUID", []interface{}{arg1, arg2})
	fake.getAppGUIDMutex.Unlock()
	if fake.GetAppGUIDStub != nil {
		return fake.GetAppGUIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAppGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CliClient) GetAppGUIDCallCount() int {
	fake.getAppGUIDMutex.RLock()
	defer fake.getAppGUIDMutex.RUnlock()
	return len(fake.getAppGUIDArgsForCall)
}

func (fake *CliClient) GetAppGUIDCalls(stub func(plugin.CliConnection, string) (string, error)) {
	fake.getAppGUIDMutex.Lock()
	defer fake.getAppGUIDMutex.Unlock()
	fake.GetAppGUIDStub = stub
}

func (fake *CliClient) GetAppGUIDArgsForCall(i int) (plugin.CliConnection, string) {
	fake.getAppGUIDMutex.RLock()
	defer fake.getAppGUIDMutex.RUnlock()
	argsForCall := fake.getAppGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CliClient) GetAppGUIDReturns(result1 string, result2 error) {
	fake.getAppGUIDMutex.Lock()
	defer fake.getAppGUIDMutex.Unlock()
	fake.GetAppGUIDStub = nil
	fake.getAppGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetAppGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getAppGUIDMutex.Lock()
	defer fake.getAppGUIDMutex.Unlock()
	fake.GetAppGUIDStub = nil
	if fake.getAppGUIDReturnsOnCall == nil {
		fake.getAppGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getAppGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetDomainGUID(arg1 plugin.CliConnection, arg2 string) (string, error) {
	fake.getDomainGUIDMutex.Lock()
	ret, specificReturn := fake.getDomainGUIDReturnsOnCall[len(fake.getDomainGUIDArgsForCall)]
	fake.getDomainGUIDArgsForCall = append(fake.getDomainGUIDArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetDomainGUID", []interface{}{arg1, arg2})
	fake.getDomainGUIDMutex.Unlock()
	if fake.GetDomainGUIDStub != nil {
		return fake.GetDomainGUIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDomainGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CliClient) GetDomainGUIDCallCount() int {
	fake.getDomainGUIDMutex.RLock()
	defer fake.getDomainGUIDMutex.RUnlock()
	return len(fake.getDomainGUIDArgsForCall)
}

func (fake *CliClient) GetDomainGUIDCalls(stub func(plugin.CliConnection, string) (string, error)) {
	fake.getDomainGUIDMutex.Lock()
	defer fake.getDomainGUIDMutex.Unlock()
	fake.GetDomainGUIDStub = stub
}

func (fake *CliClient) GetDomainGUIDArgsForCall(i int) (plugin.CliConnection, string) {
	fake.getDomainGUIDMutex.RLock()
	defer fake.getDomainGUIDMutex.RUnlock()
	argsForCall := fake.getDomainGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CliClient) GetDomainGUIDReturns(result1 string, result2 error) {
	fake.getDomainGUIDMutex.Lock()
	defer fake.getDomainGUIDMutex.Unlock()
	fake.GetDomainGUIDStub = nil
	fake.getDomainGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetDomainGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getDomainGUIDMutex.Lock()
	defer fake.getDomainGUIDMutex.Unlock()
	fake.GetDomainGUIDStub = nil
	if fake.getDomainGUIDReturnsOnCall == nil {
		fake.getDomainGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getDomainGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetRouteGUID(arg1 plugin.CliConnection, arg2 string, arg3 string) (string, error) {
	fake.getRouteGUIDMutex.Lock()
	ret, specificReturn := fake.getRouteGUIDReturnsOnCall[len(fake.getRouteGUIDArgsForCall)]
	fake.getRouteGUIDArgsForCall = append(fake.getRouteGUIDArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetRouteGUID", []interface{}{arg1, arg2, arg3})
	fake.getRouteGUIDMutex.Unlock()
	if fake.GetRouteGUIDStub != nil {
		return fake.GetRouteGUIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getRouteGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CliClient) GetRouteGUIDCallCount() int {
	fake.getRouteGUIDMutex.RLock()
	defer fake.getRouteGUIDMutex.RUnlock()
	return len(fake.getRouteGUIDArgsForCall)
}

func (fake *CliClient) GetRouteGUIDCalls(stub func(plugin.CliConnection, string, string) (string, error)) {
	fake.getRouteGUIDMutex.Lock()
	defer fake.getRouteGUIDMutex.Unlock()
	fake.GetRouteGUIDStub = stub
}

func (fake *CliClient) GetRouteGUIDArgsForCall(i int) (plugin.CliConnection, string, string) {
	fake.getRouteGUIDMutex.RLock()
	defer fake.getRouteGUIDMutex.RUnlock()
	argsForCall := fake.getRouteGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CliClient) GetRouteGUIDReturns(result1 string, result2 error) {
	fake.getRouteGUIDMutex.Lock()
	defer fake.getRouteGUIDMutex.Unlock()
	fake.GetRouteGUIDStub = nil
	fake.getRouteGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetRouteGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getRouteGUIDMutex.Lock()
	defer fake.getRouteGUIDMutex.Unlock()
	fake.GetRouteGUIDStub = nil
	if fake.getRouteGUIDReturnsOnCall == nil {
		fake.getRouteGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getRouteGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetRouteMappingGUID(arg1 plugin.CliConnection, arg2 string, arg3 string) (string, error) {
	fake.getRouteMappingGUIDMutex.Lock()
	ret, specificReturn := fake.getRouteMappingGUIDReturnsOnCall[len(fake.getRouteMappingGUIDArgsForCall)]
	fake.getRouteMappingGUIDArgsForCall = append(fake.getRouteMappingGUIDArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetRouteMappingGUID", []interface{}{arg1, arg2, arg3})
	fake.getRouteMappingGUIDMutex.Unlock()
	if fake.GetRouteMappingGUIDStub != nil {
		return fake.GetRouteMappingGUIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getRouteMappingGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CliClient) GetRouteMappingGUIDCallCount() int {
	fake.getRouteMappingGUIDMutex.RLock()
	defer fake.getRouteMappingGUIDMutex.RUnlock()
	return len(fake.getRouteMappingGUIDArgsForCall)
}

func (fake *CliClient) GetRouteMappingGUIDCalls(stub func(plugin.CliConnection, string, string) (string, error)) {
	fake.getRouteMappingGUIDMutex.Lock()
	defer fake.getRouteMappingGUIDMutex.Unlock()
	fake.GetRouteMappingGUIDStub = stub
}

func (fake *CliClient) GetRouteMappingGUIDArgsForCall(i int) (plugin.CliConnection, string, string) {
	fake.getRouteMappingGUIDMutex.RLock()
	defer fake.getRouteMappingGUIDMutex.RUnlock()
	argsForCall := fake.getRouteMappingGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CliClient) GetRouteMappingGUIDReturns(result1 string, result2 error) {
	fake.getRouteMappingGUIDMutex.Lock()
	defer fake.getRouteMappingGUIDMutex.Unlock()
	fake.GetRouteMappingGUIDStub = nil
	fake.getRouteMappingGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) GetRouteMappingGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getRouteMappingGUIDMutex.Lock()
	defer fake.getRouteMappingGUIDMutex.Unlock()
	fake.GetRouteMappingGUIDStub = nil
	if fake.getRouteMappingGUIDReturnsOnCall == nil {
		fake.getRouteMappingGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getRouteMappingGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CliClient) SetRouteMappingWeight(arg1 plugin.CliConnection, arg2 string, arg3 int) error {
	fake.setRouteMappingWeightMutex.Lock()
	ret, specificReturn := fake.setRouteMappingWeightReturnsOnCall[len(fake.setRouteMappingWeightArgsForCall)]
	fake.setRouteMappingWeightArgsForCall = append(fake.setRouteMappingWeightArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetRouteMappingWeight", []interface{}{arg1, arg2, arg3})
	fake.setRouteMappingWeightMutex.Unlock()
	if fake.SetRouteMappingWeightStub != nil {
		return fake.SetRouteMappingWeightStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setRouteMappingWeightReturns
	return fakeReturns.result1
}

func (fake *CliClient) SetRouteMappingWeightCallCount() int {
	fake.setRouteMappingWeightMutex.RLock()
	defer fake.setRouteMappingWeightMutex.RUnlock()
	return len(fake.setRouteMappingWeightArgsForCall)
}

func (fake *CliClient) SetRouteMappingWeightCalls(stub func(plugin.CliConnection, string, int) error) {
	fake.setRouteMappingWeightMutex.Lock()
	defer fake.setRouteMappingWeightMutex.Unlock()
	fake.SetRouteMappingWeightStub = stub
}

func (fake *CliClient) SetRouteMappingWeightArgsForCall(i int) (plugin.CliConnection, string, int) {
	fake.setRouteMappingWeightMutex.RLock()
	defer fake.setRouteMappingWeightMutex.RUnlock()
	argsForCall := fake.setRouteMappingWeightArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CliClient) SetRouteMappingWeightReturns(result1 error) {
	fake.setRouteMappingWeightMutex.Lock()
	defer fake.setRouteMappingWeightMutex.Unlock()
	fake.SetRouteMappingWeightStub = nil
	fake.setRouteMappingWeightReturns = struct {
		result1 error
	}{result1}
}

func (fake *CliClient) SetRouteMappingWeightReturnsOnCall(i int, result1 error) {
	fake.setRouteMappingWeightMutex.Lock()
	defer fake.setRouteMappingWeightMutex.Unlock()
	fake.SetRouteMappingWeightStub = nil
	if fake.setRouteMappingWeightReturnsOnCall == nil {
		fake.setRouteMappingWeightReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRouteMappingWeightReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CliClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppGUIDMutex.RLock()
	defer fake.getAppGUIDMutex.RUnlock()
	fake.getDomainGUIDMutex.RLock()
	defer fake.getDomainGUIDMutex.RUnlock()
	fake.getRouteGUIDMutex.RLock()
	defer fake.getRouteGUIDMutex.RUnlock()
	fake.getRouteMappingGUIDMutex.RLock()
	defer fake.getRouteMappingGUIDMutex.RUnlock()
	fake.setRouteMappingWeightMutex.RLock()
	defer fake.setRouteMappingWeightMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CliClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
