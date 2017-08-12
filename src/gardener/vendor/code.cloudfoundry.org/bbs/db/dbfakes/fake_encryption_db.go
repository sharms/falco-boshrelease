// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/lager"
)

type FakeEncryptionDB struct {
	EncryptionKeyLabelStub        func(logger lager.Logger) (string, error)
	encryptionKeyLabelMutex       sync.RWMutex
	encryptionKeyLabelArgsForCall []struct {
		logger lager.Logger
	}
	encryptionKeyLabelReturns struct {
		result1 string
		result2 error
	}
	encryptionKeyLabelReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SetEncryptionKeyLabelStub        func(logger lager.Logger, encryptionKeyLabel string) error
	setEncryptionKeyLabelMutex       sync.RWMutex
	setEncryptionKeyLabelArgsForCall []struct {
		logger             lager.Logger
		encryptionKeyLabel string
	}
	setEncryptionKeyLabelReturns struct {
		result1 error
	}
	setEncryptionKeyLabelReturnsOnCall map[int]struct {
		result1 error
	}
	PerformEncryptionStub        func(logger lager.Logger) error
	performEncryptionMutex       sync.RWMutex
	performEncryptionArgsForCall []struct {
		logger lager.Logger
	}
	performEncryptionReturns struct {
		result1 error
	}
	performEncryptionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEncryptionDB) EncryptionKeyLabel(logger lager.Logger) (string, error) {
	fake.encryptionKeyLabelMutex.Lock()
	ret, specificReturn := fake.encryptionKeyLabelReturnsOnCall[len(fake.encryptionKeyLabelArgsForCall)]
	fake.encryptionKeyLabelArgsForCall = append(fake.encryptionKeyLabelArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("EncryptionKeyLabel", []interface{}{logger})
	fake.encryptionKeyLabelMutex.Unlock()
	if fake.EncryptionKeyLabelStub != nil {
		return fake.EncryptionKeyLabelStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.encryptionKeyLabelReturns.result1, fake.encryptionKeyLabelReturns.result2
}

func (fake *FakeEncryptionDB) EncryptionKeyLabelCallCount() int {
	fake.encryptionKeyLabelMutex.RLock()
	defer fake.encryptionKeyLabelMutex.RUnlock()
	return len(fake.encryptionKeyLabelArgsForCall)
}

func (fake *FakeEncryptionDB) EncryptionKeyLabelArgsForCall(i int) lager.Logger {
	fake.encryptionKeyLabelMutex.RLock()
	defer fake.encryptionKeyLabelMutex.RUnlock()
	return fake.encryptionKeyLabelArgsForCall[i].logger
}

func (fake *FakeEncryptionDB) EncryptionKeyLabelReturns(result1 string, result2 error) {
	fake.EncryptionKeyLabelStub = nil
	fake.encryptionKeyLabelReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptionDB) EncryptionKeyLabelReturnsOnCall(i int, result1 string, result2 error) {
	fake.EncryptionKeyLabelStub = nil
	if fake.encryptionKeyLabelReturnsOnCall == nil {
		fake.encryptionKeyLabelReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.encryptionKeyLabelReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptionDB) SetEncryptionKeyLabel(logger lager.Logger, encryptionKeyLabel string) error {
	fake.setEncryptionKeyLabelMutex.Lock()
	ret, specificReturn := fake.setEncryptionKeyLabelReturnsOnCall[len(fake.setEncryptionKeyLabelArgsForCall)]
	fake.setEncryptionKeyLabelArgsForCall = append(fake.setEncryptionKeyLabelArgsForCall, struct {
		logger             lager.Logger
		encryptionKeyLabel string
	}{logger, encryptionKeyLabel})
	fake.recordInvocation("SetEncryptionKeyLabel", []interface{}{logger, encryptionKeyLabel})
	fake.setEncryptionKeyLabelMutex.Unlock()
	if fake.SetEncryptionKeyLabelStub != nil {
		return fake.SetEncryptionKeyLabelStub(logger, encryptionKeyLabel)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setEncryptionKeyLabelReturns.result1
}

func (fake *FakeEncryptionDB) SetEncryptionKeyLabelCallCount() int {
	fake.setEncryptionKeyLabelMutex.RLock()
	defer fake.setEncryptionKeyLabelMutex.RUnlock()
	return len(fake.setEncryptionKeyLabelArgsForCall)
}

func (fake *FakeEncryptionDB) SetEncryptionKeyLabelArgsForCall(i int) (lager.Logger, string) {
	fake.setEncryptionKeyLabelMutex.RLock()
	defer fake.setEncryptionKeyLabelMutex.RUnlock()
	return fake.setEncryptionKeyLabelArgsForCall[i].logger, fake.setEncryptionKeyLabelArgsForCall[i].encryptionKeyLabel
}

func (fake *FakeEncryptionDB) SetEncryptionKeyLabelReturns(result1 error) {
	fake.SetEncryptionKeyLabelStub = nil
	fake.setEncryptionKeyLabelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEncryptionDB) SetEncryptionKeyLabelReturnsOnCall(i int, result1 error) {
	fake.SetEncryptionKeyLabelStub = nil
	if fake.setEncryptionKeyLabelReturnsOnCall == nil {
		fake.setEncryptionKeyLabelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setEncryptionKeyLabelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEncryptionDB) PerformEncryption(logger lager.Logger) error {
	fake.performEncryptionMutex.Lock()
	ret, specificReturn := fake.performEncryptionReturnsOnCall[len(fake.performEncryptionArgsForCall)]
	fake.performEncryptionArgsForCall = append(fake.performEncryptionArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("PerformEncryption", []interface{}{logger})
	fake.performEncryptionMutex.Unlock()
	if fake.PerformEncryptionStub != nil {
		return fake.PerformEncryptionStub(logger)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.performEncryptionReturns.result1
}

func (fake *FakeEncryptionDB) PerformEncryptionCallCount() int {
	fake.performEncryptionMutex.RLock()
	defer fake.performEncryptionMutex.RUnlock()
	return len(fake.performEncryptionArgsForCall)
}

func (fake *FakeEncryptionDB) PerformEncryptionArgsForCall(i int) lager.Logger {
	fake.performEncryptionMutex.RLock()
	defer fake.performEncryptionMutex.RUnlock()
	return fake.performEncryptionArgsForCall[i].logger
}

func (fake *FakeEncryptionDB) PerformEncryptionReturns(result1 error) {
	fake.PerformEncryptionStub = nil
	fake.performEncryptionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEncryptionDB) PerformEncryptionReturnsOnCall(i int, result1 error) {
	fake.PerformEncryptionStub = nil
	if fake.performEncryptionReturnsOnCall == nil {
		fake.performEncryptionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.performEncryptionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEncryptionDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.encryptionKeyLabelMutex.RLock()
	defer fake.encryptionKeyLabelMutex.RUnlock()
	fake.setEncryptionKeyLabelMutex.RLock()
	defer fake.setEncryptionKeyLabelMutex.RUnlock()
	fake.performEncryptionMutex.RLock()
	defer fake.performEncryptionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEncryptionDB) recordInvocation(key string, args []interface{}) {
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

var _ db.EncryptionDB = new(FakeEncryptionDB)