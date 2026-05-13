package main

import (
	"system-design/solids/OCP/services"
)

func main() {
	//Existing code remains unchanged. We can add new notification services without modifying existing code, adhering to the Open/Closed Principle.
	emailService := &services.EmailNotificationService{}
	smsService := &services.SMSNotificationService{}
	// Adding a new notification service (PushNotificationService) without modifying existing code.
	pushService := &services.PushNotificationService{}

	emailSender := services.NewNotificationSender(emailService)
	smsSender := services.NewNotificationSender(smsService)
	// Using the new PushNotificationService without changing existing code.
	pushSender := services.NewNotificationSender(pushService)

	emailSender.Notify("Hello via Email!")
	smsSender.Notify("Hello via SMS!")
	// Notify using the new PushNotificationService without modifying existing code.
	pushSender.Notify("Hello via Push Notification!")
}
