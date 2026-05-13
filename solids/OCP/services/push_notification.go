package services

// PushNotificationService implements the NotificationService interface for sending push notifications.
type PushNotificationService struct{}

// SendNotification sends a push notification with the given message.
func (p *PushNotificationService) SendNotification(message string) error {
	// Simulate sending a push notification
	println("Sending push notification:", message)
	return nil
}
