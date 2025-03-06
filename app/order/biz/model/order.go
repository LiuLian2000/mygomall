// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Address struct {
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

const (
	OrderStatePlaced   = 0
	OrderStatePaid     = 1
	OrderStateCanceled = 2
)

type Order struct {
	Base
	UserId     int64       `gorm:"index:idx_orders_user_deleted"`
	Address    Address     `gorm:"embedded"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer"`
	OrderState int32
	DeletedAt  soft_delete.DeletedAt `gorm:"index:idx_orders_user_deleted"`
}

func (o Order) TableName() string {
	return "order"
}

// TODO db这里查一下有没有加withcontext
func CreateOrder(db *gorm.DB, ctx context.Context, order *Order) (err error) {
	err = db.WithContext(ctx).Create(order).Error
	return
}

func UpdateOrder(db *gorm.DB, ctx context.Context, order *Order) (err error) {
	err = db.WithContext(ctx).Model(&order).
		Where("id = ?", order.ID).
		Updates(order).Error
	return
}

func ListOrder(db *gorm.DB, ctx context.Context, userId int64) (orders []Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

func ListDeletedOrder(db *gorm.DB, ctx context.Context, userId int64) (orders []Order, err error) {
	err = db.WithContext(ctx).Unscoped().Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

// func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
// 	err = db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
// 	return
// }

// func GetOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (order Order, err error) {
// 	err = db.Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
// 	return
// }

// func UpdateOrderState(db *gorm.DB, ctx context.Context, userId uint32, orderId string, state OrderState) error {
// 	return db.Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
// }
