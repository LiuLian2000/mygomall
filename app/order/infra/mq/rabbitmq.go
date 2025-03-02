package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/wagslane/go-rabbitmq"
)

var (
	OrderEventPublisher *rabbitmq.Publisher
	conn                *rabbitmq.Conn
)

func Init() {
	//TODO 加配置
	conn, err := rabbitmq.NewConn(
		"amqp://root:devops666@localhost:5672/admin_vhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		klog.Fatal(err)
	}

	OrderEventPublisher, err = rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("order_event_exchange"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeKind("topic"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		klog.Fatal(err)
	}

}

// TODO main里要加上关闭mq连接
func Shutdown() {
	conn.Close()
	OrderEventPublisher.Close()
}
