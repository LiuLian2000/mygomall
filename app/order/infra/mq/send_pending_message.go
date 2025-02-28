package mq

import (
	"context"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/wagslane/go-rabbitmq"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

var tracer = otel.Tracer("order-local-service")

type MessageSender struct {
	DB        *gorm.DB
	Publisher *rabbitmq.Publisher
}

// NewMessageSender 初始化
func NewMessageSender(db *gorm.DB, publisher *rabbitmq.Publisher) *MessageSender {
	return &MessageSender{DB: db, Publisher: publisher}
}

// Start 定时扫描待发送消息
func (s *MessageSender) Start() {
	ticker := time.NewTicker(3 * time.Second) // 每 5 秒执行一次
	defer ticker.Stop()

	for range ticker.C {
		s.sendPendingMessages()
	}
}

// sendPendingMessages 发送待处理的消息
func (s *MessageSender) sendPendingMessages() {
	ctx, span := tracer.Start(context.Background(), "pollLocalMessages")
	defer span.End()
	pendingMessages, err := model.QueryPendingMessages(s.DB, ctx)

	if err != nil {
		klog.Error(err)
		return
	}

	for _, msg := range pendingMessages {

		err = s.Publisher.Publish(
			[]byte(msg.MessageBody),
			[]string{msg.Topic},
			rabbitmq.WithPublishOptionsContentType("application/json"),
			rabbitmq.WithPublishOptionsExchange("order_event_exchange"),
		)

		if err == nil {
			err = model.MarkMessageSended(s.DB, ctx, msg) // 标记为已发送
			if err != nil {                               // 更改失败无所谓，消费端有幂等
				return
			}
		} else {
			klog.Error(err)
		}
	}
}
