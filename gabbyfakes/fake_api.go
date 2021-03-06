// Code generated by counterfeiter. DO NOT EDIT.
package gabbyfakes

import (
	"context"
	"sync"

	"github.com/joefitzgerald/gabby"
)

type FakeAPI struct {
	GetMeStub        func(context.Context) (string, error)
	getMeMutex       sync.RWMutex
	getMeArgsForCall []struct {
		arg1 context.Context
	}
	getMeReturns struct {
		result1 string
		result2 error
	}
	getMeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRecurringEventsStub        func(context.Context) (gabby.Events, error)
	getRecurringEventsMutex       sync.RWMutex
	getRecurringEventsArgsForCall []struct {
		arg1 context.Context
	}
	getRecurringEventsReturns struct {
		result1 gabby.Events
		result2 error
	}
	getRecurringEventsReturnsOnCall map[int]struct {
		result1 gabby.Events
		result2 error
	}
	GetRecurringEventsWithInstancesForWeeksStub        func(context.Context, int) (gabby.Events, error)
	getRecurringEventsWithInstancesForWeeksMutex       sync.RWMutex
	getRecurringEventsWithInstancesForWeeksArgsForCall []struct {
		arg1 context.Context
		arg2 int
	}
	getRecurringEventsWithInstancesForWeeksReturns struct {
		result1 gabby.Events
		result2 error
	}
	getRecurringEventsWithInstancesForWeeksReturnsOnCall map[int]struct {
		result1 gabby.Events
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) GetMe(arg1 context.Context) (string, error) {
	fake.getMeMutex.Lock()
	ret, specificReturn := fake.getMeReturnsOnCall[len(fake.getMeArgsForCall)]
	fake.getMeArgsForCall = append(fake.getMeArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetMeStub
	fakeReturns := fake.getMeReturns
	fake.recordInvocation("GetMe", []interface{}{arg1})
	fake.getMeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetMeCallCount() int {
	fake.getMeMutex.RLock()
	defer fake.getMeMutex.RUnlock()
	return len(fake.getMeArgsForCall)
}

func (fake *FakeAPI) GetMeCalls(stub func(context.Context) (string, error)) {
	fake.getMeMutex.Lock()
	defer fake.getMeMutex.Unlock()
	fake.GetMeStub = stub
}

func (fake *FakeAPI) GetMeArgsForCall(i int) context.Context {
	fake.getMeMutex.RLock()
	defer fake.getMeMutex.RUnlock()
	argsForCall := fake.getMeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPI) GetMeReturns(result1 string, result2 error) {
	fake.getMeMutex.Lock()
	defer fake.getMeMutex.Unlock()
	fake.GetMeStub = nil
	fake.getMeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetMeReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMeMutex.Lock()
	defer fake.getMeMutex.Unlock()
	fake.GetMeStub = nil
	if fake.getMeReturnsOnCall == nil {
		fake.getMeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getMeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetRecurringEvents(arg1 context.Context) (gabby.Events, error) {
	fake.getRecurringEventsMutex.Lock()
	ret, specificReturn := fake.getRecurringEventsReturnsOnCall[len(fake.getRecurringEventsArgsForCall)]
	fake.getRecurringEventsArgsForCall = append(fake.getRecurringEventsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetRecurringEventsStub
	fakeReturns := fake.getRecurringEventsReturns
	fake.recordInvocation("GetRecurringEvents", []interface{}{arg1})
	fake.getRecurringEventsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetRecurringEventsCallCount() int {
	fake.getRecurringEventsMutex.RLock()
	defer fake.getRecurringEventsMutex.RUnlock()
	return len(fake.getRecurringEventsArgsForCall)
}

func (fake *FakeAPI) GetRecurringEventsCalls(stub func(context.Context) (gabby.Events, error)) {
	fake.getRecurringEventsMutex.Lock()
	defer fake.getRecurringEventsMutex.Unlock()
	fake.GetRecurringEventsStub = stub
}

func (fake *FakeAPI) GetRecurringEventsArgsForCall(i int) context.Context {
	fake.getRecurringEventsMutex.RLock()
	defer fake.getRecurringEventsMutex.RUnlock()
	argsForCall := fake.getRecurringEventsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPI) GetRecurringEventsReturns(result1 gabby.Events, result2 error) {
	fake.getRecurringEventsMutex.Lock()
	defer fake.getRecurringEventsMutex.Unlock()
	fake.GetRecurringEventsStub = nil
	fake.getRecurringEventsReturns = struct {
		result1 gabby.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetRecurringEventsReturnsOnCall(i int, result1 gabby.Events, result2 error) {
	fake.getRecurringEventsMutex.Lock()
	defer fake.getRecurringEventsMutex.Unlock()
	fake.GetRecurringEventsStub = nil
	if fake.getRecurringEventsReturnsOnCall == nil {
		fake.getRecurringEventsReturnsOnCall = make(map[int]struct {
			result1 gabby.Events
			result2 error
		})
	}
	fake.getRecurringEventsReturnsOnCall[i] = struct {
		result1 gabby.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeks(arg1 context.Context, arg2 int) (gabby.Events, error) {
	fake.getRecurringEventsWithInstancesForWeeksMutex.Lock()
	ret, specificReturn := fake.getRecurringEventsWithInstancesForWeeksReturnsOnCall[len(fake.getRecurringEventsWithInstancesForWeeksArgsForCall)]
	fake.getRecurringEventsWithInstancesForWeeksArgsForCall = append(fake.getRecurringEventsWithInstancesForWeeksArgsForCall, struct {
		arg1 context.Context
		arg2 int
	}{arg1, arg2})
	stub := fake.GetRecurringEventsWithInstancesForWeeksStub
	fakeReturns := fake.getRecurringEventsWithInstancesForWeeksReturns
	fake.recordInvocation("GetRecurringEventsWithInstancesForWeeks", []interface{}{arg1, arg2})
	fake.getRecurringEventsWithInstancesForWeeksMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeksCallCount() int {
	fake.getRecurringEventsWithInstancesForWeeksMutex.RLock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.RUnlock()
	return len(fake.getRecurringEventsWithInstancesForWeeksArgsForCall)
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeksCalls(stub func(context.Context, int) (gabby.Events, error)) {
	fake.getRecurringEventsWithInstancesForWeeksMutex.Lock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.Unlock()
	fake.GetRecurringEventsWithInstancesForWeeksStub = stub
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeksArgsForCall(i int) (context.Context, int) {
	fake.getRecurringEventsWithInstancesForWeeksMutex.RLock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.RUnlock()
	argsForCall := fake.getRecurringEventsWithInstancesForWeeksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeksReturns(result1 gabby.Events, result2 error) {
	fake.getRecurringEventsWithInstancesForWeeksMutex.Lock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.Unlock()
	fake.GetRecurringEventsWithInstancesForWeeksStub = nil
	fake.getRecurringEventsWithInstancesForWeeksReturns = struct {
		result1 gabby.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetRecurringEventsWithInstancesForWeeksReturnsOnCall(i int, result1 gabby.Events, result2 error) {
	fake.getRecurringEventsWithInstancesForWeeksMutex.Lock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.Unlock()
	fake.GetRecurringEventsWithInstancesForWeeksStub = nil
	if fake.getRecurringEventsWithInstancesForWeeksReturnsOnCall == nil {
		fake.getRecurringEventsWithInstancesForWeeksReturnsOnCall = make(map[int]struct {
			result1 gabby.Events
			result2 error
		})
	}
	fake.getRecurringEventsWithInstancesForWeeksReturnsOnCall[i] = struct {
		result1 gabby.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMeMutex.RLock()
	defer fake.getMeMutex.RUnlock()
	fake.getRecurringEventsMutex.RLock()
	defer fake.getRecurringEventsMutex.RUnlock()
	fake.getRecurringEventsWithInstancesForWeeksMutex.RLock()
	defer fake.getRecurringEventsWithInstancesForWeeksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPI) recordInvocation(key string, args []interface{}) {
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

var _ gabby.API = new(FakeAPI)
