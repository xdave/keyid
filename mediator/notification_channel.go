package mediator

import (
	"github.com/xdave/keyid/interfaces"

	"go.uber.org/fx"
)

type NotificationChannel struct {
	channel chan interfaces.Notification
}

type NotificationChannelParams struct {
	fx.In
}

type NotificationChannelResult struct {
	fx.Out
	Channel interfaces.NotificationChannel
}

func NewNotificationChannel(params NotificationChannelParams) NotificationChannelResult {
	return NotificationChannelResult{
		Channel: &NotificationChannel{
			channel: make(chan interfaces.Notification, 10),
		},
	}

}

func (c *NotificationChannel) Send(notification interfaces.Notification) {
	go func() {
		c.channel <- notification
	}()
}

func (c *NotificationChannel) Receive() interfaces.Notification {
	return <-c.channel
}
