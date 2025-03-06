package service

import (
	"context"

	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/cart"
	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	resp = new(order.ListOrderResp)

	var orders []model.Order
	// 查未删除
	if req.ListDeleted == 0 {
		orders, err = model.ListOrder(mysql.DB, s.ctx, req.UserId)
		if err != nil {
			klog.Errorf("model.ListOrder.err:%v", err)
			return nil, err
		}

	} else {
		// 查回收站
		orders, err = model.ListDeletedOrder(mysql.DB, s.ctx, req.UserId)
		if err != nil {
			klog.Errorf("model.ListOrder.err:%v", err)
			return nil, err
		}
	}

	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, v := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: *v.Cost,
				Item: &cart.CartItem{
					ProductId: v.ID,
					Quantity:  *v.Quantity,
				},
			})
		}
		o := &order.Order{
			OrderId:   v.ID,
			UserId:    v.UserId,
			CreatedAt: int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Address.Country,
				City:          v.Address.City,
				StreetAddress: v.Address.StreetAddress,
				ZipCode:       v.Address.ZipCode,
			},
			OrderItems: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}

	return
}
