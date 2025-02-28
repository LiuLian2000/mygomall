package model

import (
	"context"

	"gorm.io/gorm"
)

type LocalMessage struct {
	Base
	Topic       string
	MessageBody string
	Status      int
	RetryCount  int
}

type Product struct {
	ProductId int64
	Quantity  int32
}

type OrderMessageBody struct {
	UserId int64
	Items  []Product
}

func (oi LocalMessage) TableName() string {
	return "local_message"
}

func CreateLocalMessage(db *gorm.DB, ctx context.Context, msg *LocalMessage) (err error) {
	err = db.WithContext(ctx).Create(&msg).Error
	return
}

func QueryPendingMessages(db *gorm.DB, ctx context.Context) (pendingmessages []*LocalMessage, err error) {
	err = db.WithContext(ctx).Where("status = ?", 0).Find(&pendingmessages).Error
	return
}

func MarkMessageSended(db *gorm.DB, ctx context.Context, msg *LocalMessage) (err error) {
	err = db.WithContext(ctx).Model(&msg).Update("status", 1).Error
	return
}
