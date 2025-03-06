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
)

type OrderItem struct {
	Base
	OrderIdRefer int64  `gorm:"size:256;index"`
	Quantity     *int32 //用struct update更新的时候，用*区分是否是0值,如果要更新成0值，这样用update也行，不会忽略，只会忽略没改初始值的位置
	Cost         *float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}

func CreateItemLists(db *gorm.DB, ctx context.Context, itemList []*OrderItem) (err error) {
	err = db.WithContext(ctx).Create(&itemList).Error
	return
}

func UpdateItemLists(db *gorm.DB, ctx context.Context, itemList []*OrderItem) (err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range itemList {
			// 直接使用结构体更新（非空指针字段会被更新）
			if err := tx.Model(&item).
				Where("id = ? AND order_id_refer = ?", item.ID, item.OrderIdRefer).
				Updates(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return
}
