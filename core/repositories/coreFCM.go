package repositories

import (
	"github.com/bondzai/protocraft/core/entities"
	"github.com/labstack/gommon/log"
)

type coreFCMMessaging struct{}

func NewCoreFCMMessaging() CoreMessaging {
	return &coreFCMMessaging{}
}

func (c *coreFCMMessaging) PushNotification(m *entities.CorePushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Debugf("Pushed FCM notification with data: %v", m)

	return nil
}
