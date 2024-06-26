package mediator

import (
	"context"
	"time"

	"github.com/xdave/keyid/interfaces"

	"go.uber.org/fx"
)

type Mediator struct {
	notifications        interfaces.NotificationChannel
	notificationHandlers []interfaces.NotificationHandler
	running              bool
}

type MediatorParams struct {
	fx.In
	fx.Lifecycle
	NotificationChannel  interfaces.NotificationChannel
	NotificationHandlers []interfaces.NotificationHandler `group:"notification_handlers"`
}

type MediatorResult struct {
	fx.Out
	Mediator interfaces.Mediator
}

func NewMediator(params MediatorParams) MediatorResult {
	mediator := &Mediator{
		notifications:        params.NotificationChannel,
		notificationHandlers: params.NotificationHandlers,
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			mediator.running = true
			go mediator.HandleNotifications()
			return nil
		},
	})

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			mediator.running = false
			return nil
		},
	})

	return MediatorResult{Mediator: mediator}
}

func (m *Mediator) HandleNotifications() {
	for m.running {
		notification := m.notifications.Receive()
		for _, handler := range m.notificationHandlers {
			if notification != nil && notification.GetType() == handler.GetType() {
				go handler.Handle(notification)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
