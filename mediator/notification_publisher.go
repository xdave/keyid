package mediator

import (
	"github.com/xdave/keyid/interfaces"

	"go.uber.org/fx"
)

type NotificationPublisher struct {
	channel interfaces.NotificationChannel
}

type NotificationPublisherParams struct {
	fx.In
	Channel interfaces.NotificationChannel
}

type NotificationPublisherResult struct {
	fx.Out
	Publisher interfaces.NotificationPublisher
}

func NewNotificationPublisher(params NotificationPublisherParams) NotificationPublisherResult {
	return NotificationPublisherResult{
		Publisher: &NotificationPublisher{channel: params.Channel},
	}
}

func (p *NotificationPublisher) Publish(notification interfaces.Notification) {
	p.channel.Send(notification)
}
