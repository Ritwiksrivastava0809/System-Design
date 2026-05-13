package models

//Notification service defines the contract for sending notifications.

type NotificationService interface {
	SendNotification(message string) error //send notification with the given message
}
