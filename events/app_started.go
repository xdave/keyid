package events

import (
	"reflect"

	"github.com/xdave/keyid/interfaces"
)

type AppStarted struct{}

func NewAppStarted() interfaces.Notification {
	return &AppStarted{}
}

func (e *AppStarted) GetType() string {
	return reflect.TypeOf(e).String()
}
