package app

import (
	"github.com/xdave/keyid/client"
	"github.com/xdave/keyid/events"
	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/mediator"

	"go.uber.org/fx"
)

var Module = fx.Module("app",
	client.Module,
	mediator.Module,
	events.Module,
	fx.Invoke(func(publisher interfaces.NotificationPublisher) {
		publisher.Publish(&events.AppStarted{})
	}),
)
