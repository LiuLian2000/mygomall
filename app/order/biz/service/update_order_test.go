package service

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/dal"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/cart"
	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/stretchr/testify/assert"
)

func TestUpdateOrder_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	msg := "111"
	hmactimestamp := time.Now().Unix()
	hmacedmsg := utils.GenerateHMAC(msg, hmactimestamp)

	ctx = metainfo.WithValue(ctx, "hmacedMsg", hmacedmsg)
	ctx = metainfo.WithValue(ctx, "checkoutMsg", msg)
	ctx = metainfo.WithValue(ctx, "checkoutHmacTimestamp", strconv.FormatInt(hmactimestamp, 10))

	s := NewUpdateOrderService(ctx)

	// init req and assert value

	type testCase struct {
		name         string
		req          *order.UpdateOrderReq
		expectedErr  error
		expectedResp *order.UpdateOrderResp
	}

	// todo: edit your unit test

	testCases := []testCase{
		{
			name: "changeState",
			req: &order.UpdateOrderReq{
				ChangedOrder: &order.Order{
					OrderId:    888,
					UserId:     111,
					OrderState: 1,
				},
			},
			expectedErr:  nil,
			expectedResp: &order.UpdateOrderResp{},
		},
		{
			name: "changecity",
			req: &order.UpdateOrderReq{
				ChangedOrder: &order.Order{
					OrderId: 888,
					UserId:  111,
					Address: &order.Address{
						City: "city2",
					},
				},
			},
			expectedErr:  nil,
			expectedResp: &order.UpdateOrderResp{},
		},
		{
			name: "changeitem",
			req: &order.UpdateOrderReq{
				ChangedOrder: &order.Order{
					OrderId: 888,
					UserId:  111,
					OrderItems: []*order.OrderItem{
						{
							Item: &cart.CartItem{
								ProductId: 1,
								Quantity:  3999999,
							},
							Cost: 3999,
						},
						{
							Item: &cart.CartItem{
								ProductId: 3,
								Quantity:  7799999,
							},
							Cost: 55999,
						},
					},
				},
			},
			expectedErr:  nil,
			expectedResp: &order.UpdateOrderResp{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := s.Run(tc.req)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedResp, resp)

		})
	}
}
