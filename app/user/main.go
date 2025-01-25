package main

import (
	"context"
	"net"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/user/biz/dal"
	"github.com/Group-lifelong-youth-training/mygomall/app/user/conf"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/mtl"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/serversuite"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	serviceName := conf.GetConf().Kitex.Service

	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background())
	dal.Init()
	opts := kitexInit()

	// nodeId, _ := strconv.ParseInt(os.Getenv("NODE_ID"), 10, 64)
	// utils.InitializeNode(nodeId)
	utils.InitializeNode(1)

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: conf.GetConf().Kitex.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))

	return
}
