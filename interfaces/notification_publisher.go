package interfaces

type NotificationPublisher interface {
	Publish(notification Notification)
}
