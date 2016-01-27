// This file was generated by counterfeiter
package fakes

import (
	"net/url"
	"sync"

	"github.com/cloudfoundry/go-ccapi/v3/client"
)

type FakeClient struct {
	GetApplicationsStub        func(queryParams url.Values) ([]byte, error)
	getApplicationsMutex       sync.RWMutex
	getApplicationsArgsForCall []struct {
		queryParams url.Values
	}
	getApplicationsReturns struct {
		result1 []byte
		result2 error
	}
	GetResourceStub        func(path string) ([]byte, error)
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		path string
	}
	getResourceReturns struct {
		result1 []byte
		result2 error
	}
	GetResourcesStub        func(path string, limit int) ([]byte, error)
	getResourcesMutex       sync.RWMutex
	getResourcesArgsForCall []struct {
		path  string
		limit int
	}
	getResourcesReturns struct {
		result1 []byte
		result2 error
	}
	RefreshAuthTokenStub        func() error
	refreshAuthTokenMutex       sync.RWMutex
	refreshAuthTokenArgsForCall []struct{}
	refreshAuthTokenReturns     struct {
		result1 error
	}
	GetAuthTokenStub        func() string
	getAuthTokenMutex       sync.RWMutex
	getAuthTokenArgsForCall []struct{}
	getAuthTokenReturns     struct {
		result1 string
	}
}

func (fake *FakeClient) GetApplications(queryParams url.Values) ([]byte, error) {
	fake.getApplicationsMutex.Lock()
	fake.getApplicationsArgsForCall = append(fake.getApplicationsArgsForCall, struct {
		queryParams url.Values
	}{queryParams})
	fake.getApplicationsMutex.Unlock()
	if fake.GetApplicationsStub != nil {
		return fake.GetApplicationsStub(queryParams)
	} else {
		return fake.getApplicationsReturns.result1, fake.getApplicationsReturns.result2
	}
}

func (fake *FakeClient) GetApplicationsCallCount() int {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return len(fake.getApplicationsArgsForCall)
}

func (fake *FakeClient) GetApplicationsArgsForCall(i int) url.Values {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return fake.getApplicationsArgsForCall[i].queryParams
}

func (fake *FakeClient) GetApplicationsReturns(result1 []byte, result2 error) {
	fake.GetApplicationsStub = nil
	fake.getApplicationsReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetResource(path string) ([]byte, error) {
	fake.getResourceMutex.Lock()
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		path string
	}{path})
	fake.getResourceMutex.Unlock()
	if fake.GetResourceStub != nil {
		return fake.GetResourceStub(path)
	} else {
		return fake.getResourceReturns.result1, fake.getResourceReturns.result2
	}
}

func (fake *FakeClient) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeClient) GetResourceArgsForCall(i int) string {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return fake.getResourceArgsForCall[i].path
}

func (fake *FakeClient) GetResourceReturns(result1 []byte, result2 error) {
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetResources(path string, limit int) ([]byte, error) {
	fake.getResourcesMutex.Lock()
	fake.getResourcesArgsForCall = append(fake.getResourcesArgsForCall, struct {
		path  string
		limit int
	}{path, limit})
	fake.getResourcesMutex.Unlock()
	if fake.GetResourcesStub != nil {
		return fake.GetResourcesStub(path, limit)
	} else {
		return fake.getResourcesReturns.result1, fake.getResourcesReturns.result2
	}
}

func (fake *FakeClient) GetResourcesCallCount() int {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return len(fake.getResourcesArgsForCall)
}

func (fake *FakeClient) GetResourcesArgsForCall(i int) (string, int) {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return fake.getResourcesArgsForCall[i].path, fake.getResourcesArgsForCall[i].limit
}

func (fake *FakeClient) GetResourcesReturns(result1 []byte, result2 error) {
	fake.GetResourcesStub = nil
	fake.getResourcesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RefreshAuthToken() error {
	fake.refreshAuthTokenMutex.Lock()
	fake.refreshAuthTokenArgsForCall = append(fake.refreshAuthTokenArgsForCall, struct{}{})
	fake.refreshAuthTokenMutex.Unlock()
	if fake.RefreshAuthTokenStub != nil {
		return fake.RefreshAuthTokenStub()
	} else {
		return fake.refreshAuthTokenReturns.result1
	}
}

func (fake *FakeClient) RefreshAuthTokenCallCount() int {
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	return len(fake.refreshAuthTokenArgsForCall)
}

func (fake *FakeClient) RefreshAuthTokenReturns(result1 error) {
	fake.RefreshAuthTokenStub = nil
	fake.refreshAuthTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetAuthToken() string {
	fake.getAuthTokenMutex.Lock()
	fake.getAuthTokenArgsForCall = append(fake.getAuthTokenArgsForCall, struct{}{})
	fake.getAuthTokenMutex.Unlock()
	if fake.GetAuthTokenStub != nil {
		return fake.GetAuthTokenStub()
	} else {
		return fake.getAuthTokenReturns.result1
	}
}

func (fake *FakeClient) GetAuthTokenCallCount() int {
	fake.getAuthTokenMutex.RLock()
	defer fake.getAuthTokenMutex.RUnlock()
	return len(fake.getAuthTokenArgsForCall)
}

func (fake *FakeClient) GetAuthTokenReturns(result1 string) {
	fake.GetAuthTokenStub = nil
	fake.getAuthTokenReturns = struct {
		result1 string
	}{result1}
}

var _ client.Client = new(FakeClient)
