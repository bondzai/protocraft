package repositories

import "github.com/bondzai/protocraft/core/entities"

type CoreMessaging interface {
	PushNotification(m *entities.CorePushNotificationDto) error
}
