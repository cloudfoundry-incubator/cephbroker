// This file was generated by counterfeiter
package cephfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/cephbroker/cephbrokerlocal"
)

type FakeInvoker struct {
	InvokeStub        func(logger lager.Logger, executable string, args []string) error
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		logger     lager.Logger
		executable string
		args       []string
	}
	invokeReturns struct {
		result1 error
	}
}

func (fake *FakeInvoker) Invoke(logger lager.Logger, executable string, args []string) error {
	fake.invokeMutex.Lock()
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		logger     lager.Logger
		executable string
		args       []string
	}{logger, executable, args})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(logger, executable, args)
	} else {
		return fake.invokeReturns.result1
	}
}

func (fake *FakeInvoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeInvoker) InvokeArgsForCall(i int) (lager.Logger, string, []string) {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invokeArgsForCall[i].logger, fake.invokeArgsForCall[i].executable, fake.invokeArgsForCall[i].args
}

func (fake *FakeInvoker) InvokeReturns(result1 error) {
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 error
	}{result1}
}

var _ cephbrokerlocal.Invoker = new(FakeInvoker)
