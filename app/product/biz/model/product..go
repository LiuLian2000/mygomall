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

import time

type Product struct {
	Base
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Categories  []string `json:"categories"`
}
type Category struct {
	Base
	Name        string    `gorm:"unique" json:"name"`
	Description string    `json:"description"`
	
}

func (p Product) TableName() string {
	return "product"
}

func GetProductByID(db *gorm.DB, ctx context.Context, id int64) (product *Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Where(&Product{Base:{ID:id}}).First(&product).Error
	return
}

func GetCategoryByName(db *gorm.DB, ctx context.Context, name string) (category *Category, err error) {
	err = db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).First(&category).Error
	return
}

func CreateProduct(db *gorm.DB, ctx context.Context, product *Product) error {
	return db.WithContext(ctx).Create(product).Error
}

func CreateCategory(db *gorm.DB, ctx context.Context, category *Category) error {
	return db.WithContext(ctx).Create(category).Error
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Base.ID, _ = utils.GenerateID() // 使用 Snowflake 算法生成 ID
	return
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.Base.ID, _ = utils.GenerateID() // 使用 Snowflake 算法生成 ID
	return
}
