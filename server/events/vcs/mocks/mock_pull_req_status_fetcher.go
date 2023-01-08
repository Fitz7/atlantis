// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: PullReqStatusFetcher)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockPullReqStatusFetcher struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPullReqStatusFetcher(options ...pegomock.Option) *MockPullReqStatusFetcher {
	mock := &MockPullReqStatusFetcher{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPullReqStatusFetcher) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPullReqStatusFetcher) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPullReqStatusFetcher) FetchPullStatus(_param0 models.PullRequest) (models.PullReqStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPullReqStatusFetcher().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchPullStatus", params, []reflect.Type{reflect.TypeOf((*models.PullReqStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullReqStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullReqStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockPullReqStatusFetcher) VerifyWasCalledOnce() *VerifierMockPullReqStatusFetcher {
	return &VerifierMockPullReqStatusFetcher{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPullReqStatusFetcher) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPullReqStatusFetcher {
	return &VerifierMockPullReqStatusFetcher{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPullReqStatusFetcher) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPullReqStatusFetcher {
	return &VerifierMockPullReqStatusFetcher{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPullReqStatusFetcher) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPullReqStatusFetcher {
	return &VerifierMockPullReqStatusFetcher{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPullReqStatusFetcher struct {
	mock                   *MockPullReqStatusFetcher
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPullReqStatusFetcher) FetchPullStatus(_param0 models.PullRequest) *MockPullReqStatusFetcher_FetchPullStatus_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchPullStatus", params, verifier.timeout)
	return &MockPullReqStatusFetcher_FetchPullStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPullReqStatusFetcher_FetchPullStatus_OngoingVerification struct {
	mock              *MockPullReqStatusFetcher
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPullReqStatusFetcher_FetchPullStatus_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockPullReqStatusFetcher_FetchPullStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}