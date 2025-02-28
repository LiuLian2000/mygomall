package service

import (
	"context"

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
		hmacedmsg := metainfo.GetValue(s.ctx, "hmacedMsg")
		msg := metainfo.GetValue(s.ctx, "checkoutMsg")
		hmactimestamp := metainfo.GetValue(s.ctx, "checkoutHmacTimestamp")
		ok := utils.VerifyHMAC(hmacmsg, hmactimestamp, hmacedmsg)
		if !ok {
			err = errno.UnauthorizedUpdateOrderStatusRequestErr
			return
		}
	}

	// consignee需要修改的地方
	if req.GetChangedOrder().Address.

	return
}
