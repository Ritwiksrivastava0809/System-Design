package services

type SMSNotificationService struct{}

// SMSNotification service implements the NotificationService interface to send SMS notifications.
func (sns *SMSNotificationService) SendNotification(message string) error {
	// Logic to send SMS notification
	println("Sending SMS notification with message:", message)
	return nil
}
