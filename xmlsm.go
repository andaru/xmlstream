package xmlstream

import (
	"context"
	"io"
)

// The state machine is used to trigger validation and processing of
// XML input for parsing and data extraction purposes.
type StateMachine struct {
	Begin StateFn
	End   StateFn
	Error error

	state StateFn
}

type StateFn func(context.Context, *StateMachine) StateFn

func New(begin StateFn) *StateMachine { return &StateMachine{Begin: begin} }

func (sm *StateMachine) Run(ctx context.Context) error {
	for state := sm.Begin; state != nil; state = state(ctx, sm) {
	}
	for state := sm.End; state != nil; state = state(ctx, sm) {
	}
	return sm.Error
}

func BailWithError(err error) StateFn {
	return func(ctx context.Context, sm *StateMachine) StateFn {
		sm.Error = err
		return nil
	}
}

func IgnoreEOF(ctx context.Context, sm *StateMachine) StateFn {
	if sm.Error == io.EOF {
		sm.Error = nil
	}
	return nil
}
