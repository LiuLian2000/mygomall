package service

import (
	"context"

	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// 本地通知表事务
	// 扣减库存
	// 清理购物车
	// 创建订单

	return
}
