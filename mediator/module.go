package mediator

import (
	"github.com/xdave/keyid/interfaces"

	"go.uber.org/fx"
)

var Module = fx.Module("mediator",
	fx.Provide(
		NewNotificationChannel,
		NewNotificationPublisher,
		NewMediator,
	),
	fx.Invoke(func(mediator interfaces.Mediator) {}),
)
