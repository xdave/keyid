package interfaces

type NotificationHandler interface {
	GetType() string
	Handle(notification Notification)
}
