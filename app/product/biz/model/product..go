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
	"fmt"

	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	Store       int64   `json:"store"`
	Status      bool    `json:"status"`
}

func (p Product) TableName() string {
	return "product"
}

func GetProductByID(db *gorm.DB, ctx context.Context, id int64) (product *Product, err error) {
	p := Product{}
	p.Base.ID = id
	err = db.WithContext(ctx).Model(&Product{}).Where(&p).First(&product).Error
	product.ID = id
	return
}

func CreateProduct(db *gorm.DB, ctx context.Context, product *Product) error {

	return db.WithContext(ctx).Create(product).Error
}

func SearchProduct(db *gorm.DB, ctx context.Context, query string) (product []Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Where("name REGEXP ?", fmt.Sprintf("^.*%s.*$", query)).Find(&product).Error
	return
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Base.ID, _ = utils.GenerateID() // 使用 Snowflake 算法生成 ID

	return
}
