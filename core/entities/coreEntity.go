package entities

import "time"

type (
	InsertCoreDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Core struct {
		Id        uint32    `json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	CorePushNotificationDto struct {
		Title        string `json:"title"`
		Amount       uint32 `json:"amount"`
		ReportedTime string `json:"createdAt"`
	}
)
