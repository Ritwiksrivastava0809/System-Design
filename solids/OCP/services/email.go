package services

type EmailNotificationService struct{}

// EmailNotification service implements the NotificationService interface to send email notifications.
func (ens *EmailNotificationService) SendNotification(message string) error {
	// Logic to send email notification
	println("Sending email notification with message:", message)
	return nil
}
