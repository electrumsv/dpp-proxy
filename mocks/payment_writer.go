// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/libsv/pptcl"
	"sync"
)

// Ensure, that PaymentWriterMock does implement pptcl.PaymentWriter.
// If this is not the case, regenerate this file with moq.
var _ pptcl.PaymentWriter = &PaymentWriterMock{}

// PaymentWriterMock is a mock implementation of pptcl.PaymentWriter.
//
// 	func TestSomethingThatUsesPaymentWriter(t *testing.T) {
//
// 		// make and configure a mocked pptcl.PaymentWriter
// 		mockedPaymentWriter := &PaymentWriterMock{
// 			PaymentCreateFunc: func(ctx context.Context, args pptcl.PaymentCreateArgs, req pptcl.PaymentCreate) error {
// 				panic("mock out the PaymentCreate method")
// 			},
// 		}
//
// 		// use mockedPaymentWriter in code that requires pptcl.PaymentWriter
// 		// and then make assertions.
//
// 	}
type PaymentWriterMock struct {
	// PaymentCreateFunc mocks the PaymentCreate method.
	PaymentCreateFunc func(ctx context.Context, args pptcl.PaymentCreateArgs, req pptcl.PaymentCreate) error

	// calls tracks calls to the methods.
	calls struct {
		// PaymentCreate holds details about calls to the PaymentCreate method.
		PaymentCreate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Args is the args argument value.
			Args pptcl.PaymentCreateArgs
			// Req is the req argument value.
			Req pptcl.PaymentCreate
		}
	}
	lockPaymentCreate sync.RWMutex
}

// PaymentCreate calls PaymentCreateFunc.
func (mock *PaymentWriterMock) PaymentCreate(ctx context.Context, args pptcl.PaymentCreateArgs, req pptcl.PaymentCreate) error {
	if mock.PaymentCreateFunc == nil {
		panic("PaymentWriterMock.PaymentCreateFunc: method is nil but PaymentWriter.PaymentCreate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Args pptcl.PaymentCreateArgs
		Req  pptcl.PaymentCreate
	}{
		Ctx:  ctx,
		Args: args,
		Req:  req,
	}
	mock.lockPaymentCreate.Lock()
	mock.calls.PaymentCreate = append(mock.calls.PaymentCreate, callInfo)
	mock.lockPaymentCreate.Unlock()
	return mock.PaymentCreateFunc(ctx, args, req)
}

// PaymentCreateCalls gets all the calls that were made to PaymentCreate.
// Check the length with:
//     len(mockedPaymentWriter.PaymentCreateCalls())
func (mock *PaymentWriterMock) PaymentCreateCalls() []struct {
	Ctx  context.Context
	Args pptcl.PaymentCreateArgs
	Req  pptcl.PaymentCreate
} {
	var calls []struct {
		Ctx  context.Context
		Args pptcl.PaymentCreateArgs
		Req  pptcl.PaymentCreate
	}
	mock.lockPaymentCreate.RLock()
	calls = mock.calls.PaymentCreate
	mock.lockPaymentCreate.RUnlock()
	return calls
}
