package main

import (
	"context"

	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/service"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp = new(order.PlaceOrderResp)
	if len(req.GetOrderItems()) == 0 || req.GetUserId() <= 0 || len(req.GetEmail()) == 0 || req.GetAddress() == nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	resp.BaseResp = errno.HundleRespAndErr(resp.BaseResp, err)

	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	resp = new(order.UpdateOrderResp)

	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	resp.BaseResp = errno.HundleRespAndErr(resp.BaseResp, err)

	return resp, err
}
