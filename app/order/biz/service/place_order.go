package service

import (
	"context"
	"encoding/json"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/app/order/infra/mq"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	order "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/order"
	"github.com/wagslane/go-rabbitmq"
	"gorm.io/gorm"
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
	resp = new(order.PlaceOrderResp)
	// TODO 查库存够不够

	// 本地通知表事务
	// 扣减库存
	// 清理购物车
	// 创建订单

	orderId, err := utils.GenerateID()
	if err != nil {
		return
	}

	o := &model.Order{
		Base: model.Base{
			ID: orderId,
		},
		OrderState: model.OrderStatePlaced,
		UserId:     req.UserId,
		Consignee: model.Consignee{
			Email:         req.Email,
			Country:       req.GetAddress().Country,
			State:         req.GetAddress().State,
			City:          req.GetAddress().City,
			StreetAddress: req.GetAddress().StreetAddress,
			ZipCode:       req.GetAddress().ZipCode,
		},
	}

	var itemList []*model.OrderItem

	orderMessageBody := model.OrderMessageBody{
		UserId: req.UserId,
	}
	for _, v := range req.OrderItems {
		itemList = append(itemList, &model.OrderItem{
			Base: model.Base{
				ID: v.Item.ProductId,
			},
			OrderIdRefer: o.Base.ID,
			Quantity:     v.Item.Quantity,
			Cost:         v.Cost,
		})
		orderMessageBody.Items = append(orderMessageBody.Items, model.Product{
			ProductId: v.Item.ProductId,
			Quantity:  v.Item.Quantity,
		})
	}

	messageId, err := utils.GenerateID()

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(orderMessageBody)
	if err != nil {
		return
	}
	msg := &model.LocalMessage{
		Base: model.Base{
			ID: messageId,
		},
		Topic:       "order_created_topic",
		MessageBody: string(jsonData), // 转 JSON
		Status:      0,                // 初始状态：待发送
		RetryCount:  0,
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {

		if err := model.CreateOrder(tx, s.ctx, o); err != nil {
			return err
		}

		if err := model.CreateItemLists(tx, s.ctx, itemList); err != nil {
			return err
		}

		if err := model.CreateLocalMessage(tx, s.ctx, msg); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	err = mq.OrderEventPublisher.Publish(
		[]byte(msg.MessageBody),
		[]string{msg.Topic},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("order_event_exchange"),
	)

	if err != nil {
		return
	}

	err = model.MarkMessageSended(mysql.DB, s.ctx, msg)

	return
}
