// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"log"

	"github.com/Group-lifelong-youth-training/mygomall/app/product/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/product/biz/model"
)

// 插入几条数据测试

func main() {
	mysql.Init()

	// //自动迁移
	// if err := mysql.DB.AutoMigrate(&model.Product{}); err != nil {
	// 	log.Fatalf("AutoMigrate failed: %v", err)
	// }

	sql, err := mysql.DB.DB()
	if err != nil {
		log.Fatal("DB() fault")
	}
	defer sql.Close()
	ctx := context.Background()

	ps := []model.Product{
		{
			Base:        model.Base{ID: 123}, // 如果不需要手动设置 ID，可以去掉
			Name:        "test1",
			Description: "456",

			Price:   33.0,
			Picture: "666",
		},
		{
			Base:        model.Base{ID: 456}, // 同上
			Name:        "test2",
			Description: "457",

			Price:   34.0,
			Picture: "666",
		},
	}
	log.Printf("ps:%v\n", ps)

	// 插入产品
	for _, p := range ps {

		if err := model.CreateProduct(mysql.DB, ctx, &p); err != nil {
			log.Fatalf("Error creating product %s: %v", p.Name, err)
		}
	}
	// mysql.Init()
	// ctx := context.Background()
	// s := service.NewGetProductService(ctx)
	// req := &product.GetProductReq{Id: 1896234472328663040}
	// resp, err := s.Run(req)
	// log.Printf("resp:%v\n err:%v\n", resp, err)

}
