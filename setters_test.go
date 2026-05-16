package usecases

import (
	"testing"

	"github.com/algo2go/kite-mcp-domain"

	"github.com/stretchr/testify/assert"
)

// TestSetEventStoreSetters covers the 16 SetEventStore one-line setters
// that were previously at 0% coverage. Each setter is structurally
// identical: `func (uc *X) SetEventStore(s EventAppender) { uc.eventStore = s }`.
// The table walks every concrete UseCase type, calls SetEventStore with a
// non-nil mockEventAppender, and asserts the field is wired through.
//
// Background: these setters are wired by the composition root via
// `kc/manager_init.go` after construction so the use case can append
// audit events. They had zero coverage because all existing tests either
// construct the use case directly (skipping the setter path) or wire
// events at construction time via the New*UseCase variadic options.
// Asserting the setter behaviour locks in the wire-through contract.
func TestSetEventStoreSetters(t *testing.T) {
	tests := []struct {
		name string
		// build returns the use case with its eventStore field accessible
		// via the getter closure so we can assert post-set.
		build func() (target interface{ SetEventStore(EventAppender) }, getEventStore func() EventAppender)
	}{
		{
			name: "ConvertPositionUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &ConvertPositionUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "CreateCompositeAlertUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &CreateCompositeAlertUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PlaceGTTUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PlaceGTTUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "ModifyGTTUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &ModifyGTTUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "DeleteGTTUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &DeleteGTTUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PlaceMFOrderUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PlaceMFOrderUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "CancelMFOrderUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &CancelMFOrderUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PlaceMFSIPUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PlaceMFSIPUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "CancelMFSIPUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &CancelMFSIPUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PlaceNativeAlertUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PlaceNativeAlertUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "ModifyNativeAlertUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &ModifyNativeAlertUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "DeleteNativeAlertUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &DeleteNativeAlertUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PaperTradingToggleUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PaperTradingToggleUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "PaperTradingResetUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &PaperTradingResetUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "SetTrailingStopUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &SetTrailingStopUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
		{
			name: "CancelTrailingStopUseCase",
			build: func() (interface{ SetEventStore(EventAppender) }, func() EventAppender) {
				uc := &CancelTrailingStopUseCase{}
				return uc, func() EventAppender { return uc.eventStore }
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, getEventStore := tt.build()
			require := assert.New(t)

			// pre-condition: zero-value field is nil
			require.Nil(getEventStore(), "eventStore should be nil before SetEventStore is called")

			// act: wire a non-nil EventAppender
			appender := &mockEventAppender{}
			uc.SetEventStore(appender)

			// assert: field wired through to the same instance
			require.Same(appender, getEventStore(), "SetEventStore must wire the provided EventAppender into uc.eventStore")
		})
	}
}

// TestSetEventDispatcherSetters covers the 7 SetEventDispatcher one-line
// setters that were previously at 0% coverage. Each setter is structurally
// identical: `func (uc *X) SetEventDispatcher(d *domain.EventDispatcher) { uc.events = d }`.
// The table walks every concrete UseCase type, calls SetEventDispatcher
// with a domain.NewEventDispatcher() instance, and asserts the field is
// wired through.
//
// Note: the production setter accepts *domain.EventDispatcher (a concrete
// type from the kite-mcp-domain module), unlike SetEventStore which
// accepts an interface. The setter contract is identical: assign and
// return.
func TestSetEventDispatcherSetters(t *testing.T) {
	tests := []struct {
		name  string
		build func() (target interface {
			SetEventDispatcher(*domain.EventDispatcher)
		}, getEvents func() *domain.EventDispatcher)
	}{
		{
			name: "DeleteAlertUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &DeleteAlertUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "CancelOrderUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &CancelOrderUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "CloseAllPositionsUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &CloseAllPositionsUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "ClosePositionUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &ClosePositionUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "CreateAlertUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &CreateAlertUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "ModifyOrderUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &ModifyOrderUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
		{
			name: "PlaceOrderUseCase",
			build: func() (interface {
				SetEventDispatcher(*domain.EventDispatcher)
			}, func() *domain.EventDispatcher) {
				uc := &PlaceOrderUseCase{}
				return uc, func() *domain.EventDispatcher { return uc.events }
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, getEvents := tt.build()
			require := assert.New(t)

			// pre-condition: zero-value field is nil
			require.Nil(getEvents(), "events should be nil before SetEventDispatcher is called")

			// act: wire a non-nil dispatcher (canonical constructor used
			// throughout the test suite — see family_usecases_test.go and
			// mf_usecases_test.go)
			dispatcher := domain.NewEventDispatcher()
			uc.SetEventDispatcher(dispatcher)

			// assert: field wired through to the same instance
			require.Same(dispatcher, getEvents(), "SetEventDispatcher must wire the provided dispatcher into uc.events")
		})
	}
}
