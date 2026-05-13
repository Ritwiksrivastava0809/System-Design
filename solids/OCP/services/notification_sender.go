package services

import "system-design/solids/OCP/models"

type NotificationSender struct {
	notificationService models.NotificationService
}

func (ns NotificationSender) Notify(message string) error {
	return ns.notificationService.SendNotification(message)
}

func NewNotificationSender(notificationService models.NotificationService) NotificationSender {
	return NotificationSender{notificationService: notificationService}
}
