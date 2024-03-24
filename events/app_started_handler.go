package events

import (
	"reflect"

	"github.com/xdave/keyid/interfaces"

	"go.uber.org/fx"
)

type AppStartedHandler struct {
	publisher interfaces.NotificationPublisher
}

type AppStartedHandlerParams struct {
	fx.In
	Publisher interfaces.NotificationPublisher
}

type AppStartedHandlerResult struct {
	fx.Out
	Handler interfaces.NotificationHandler `group:"notification_handlers"`
}

func NewAppStartedHandler(params AppStartedHandlerParams) AppStartedHandlerResult {
	return AppStartedHandlerResult{
		Handler: &AppStartedHandler{
			publisher: params.Publisher,
		},
	}
}

func (handler *AppStartedHandler) GetType() string {
	return reflect.TypeOf(&AppStarted{}).String()
}

func (handler *AppStartedHandler) Handle(notification interfaces.Notification) {
}
