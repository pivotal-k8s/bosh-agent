// Code generated by counterfeiter. DO NOT EDIT.
package certfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform/cert"
)

type FakeManager struct {
	UpdateCertificatesStub        func(certs string) error
	updateCertificatesMutex       sync.RWMutex
	updateCertificatesArgsForCall []struct {
		certs string
	}
	updateCertificatesReturns struct {
		result1 error
	}
	updateCertificatesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) UpdateCertificates(certs string) error {
	fake.updateCertificatesMutex.Lock()
	ret, specificReturn := fake.updateCertificatesReturnsOnCall[len(fake.updateCertificatesArgsForCall)]
	fake.updateCertificatesArgsForCall = append(fake.updateCertificatesArgsForCall, struct {
		certs string
	}{certs})
	fake.recordInvocation("UpdateCertificates", []interface{}{certs})
	fake.updateCertificatesMutex.Unlock()
	if fake.UpdateCertificatesStub != nil {
		return fake.UpdateCertificatesStub(certs)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateCertificatesReturns.result1
}

func (fake *FakeManager) UpdateCertificatesCallCount() int {
	fake.updateCertificatesMutex.RLock()
	defer fake.updateCertificatesMutex.RUnlock()
	return len(fake.updateCertificatesArgsForCall)
}

func (fake *FakeManager) UpdateCertificatesArgsForCall(i int) string {
	fake.updateCertificatesMutex.RLock()
	defer fake.updateCertificatesMutex.RUnlock()
	return fake.updateCertificatesArgsForCall[i].certs
}

func (fake *FakeManager) UpdateCertificatesReturns(result1 error) {
	fake.UpdateCertificatesStub = nil
	fake.updateCertificatesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateCertificatesReturnsOnCall(i int, result1 error) {
	fake.UpdateCertificatesStub = nil
	if fake.updateCertificatesReturnsOnCall == nil {
		fake.updateCertificatesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateCertificatesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateCertificatesMutex.RLock()
	defer fake.updateCertificatesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ cert.Manager = new(FakeManager)
