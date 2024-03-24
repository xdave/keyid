package interfaces

type NotificationChannel interface {
	Send(notification Notification)
	Receive() Notification
}
