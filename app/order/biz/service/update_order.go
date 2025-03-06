package service

import (
	"context"
	"strconv"

	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
	"github.com/bytedance/gopkg/cloud/metainfo"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// Finish your business logic.
	//TODO 检查修改状态是否是支付服务发来的 hmac
	resp = new(order.UpdateOrderResp)

	if req.GetChangedOrder().OrderState != 0 {
		hmacedmsg, ok1 := metainfo.GetValue(s.ctx, "hmacedMsg")
		msg, ok2 := metainfo.GetValue(s.ctx, "checkoutMsg")
		hmactimestamp, ok3 := metainfo.GetValue(s.ctx, "checkoutHmacTimestamp")
		if !ok1 || !ok2 || !ok3 {
			err = errno.UnauthorizedUpdateOrderStatusRequestErr
			return
		}

		intHmactimestamp, parseErr := strconv.ParseInt(hmactimestamp, 10, 64)
		if parseErr != nil {
			err = parseErr
			return
		}

		ok := utils.VerifyHMAC(msg, intHmactimestamp, hmacedmsg)
		if !ok {
			err = errno.UnauthorizedUpdateOrderStatusRequestErr
			return
		}
	}

	// orderitem 商品需要修改的地方 改价
	var changedOrderItem []*model.OrderItem

	reqOrderItems := req.GetChangedOrder().OrderItems
	orderId := req.GetChangedOrder().OrderId

	for _, item := range reqOrderItems {
		changedOrderItem = append(changedOrderItem, &model.OrderItem{
			Base: model.Base{
				ID: item.Item.ProductId,
			},
			OrderIdRefer: orderId,
			Quantity:     &item.Item.Quantity,
			Cost:         &item.Cost,
		})
	}

	err = model.UpdateItemLists(mysql.DB, s.ctx, changedOrderItem)
	if err != nil {
		return
	}
	// 除了orderitem 需要修改的地方

	reqChangedOrder := req.GetChangedOrder()

	changedOrder := &model.Order{
		Base: model.Base{
			ID: reqChangedOrder.OrderId,
		},
		UserId:     reqChangedOrder.UserId,
		OrderState: reqChangedOrder.OrderState,
	}

	if req.GetChangedOrder().Address != nil {
		changedOrder.Address = model.Address{
			StreetAddress: reqChangedOrder.Address.StreetAddress,
			City:          reqChangedOrder.Address.City,
			State:         reqChangedOrder.Address.State,
			Country:       reqChangedOrder.Address.Country,
			ZipCode:       reqChangedOrder.Address.ZipCode,
		}
	}

	err = model.UpdateOrder(mysql.DB, s.ctx, changedOrder)

	return
}
