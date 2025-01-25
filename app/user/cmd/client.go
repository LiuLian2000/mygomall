package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/user/infra/rpc"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
)

func main() {
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// if err != nil {
	// 	panic(err)
	// }

	rpc.InitClient()
	ctx := context.Background()

	userId := int64(333)

	signature := utils.GenerateHMAC(strconv.FormatInt(userId, 10), time.Now().Unix())
	println(userId)
	println(signature)
	println("\n")
	req := &auth.DeliverTokenReq{UserId: userId, Timestamp: time.Now().Unix(), Signature: signature}
	fmt.Printf("%v", req)
	authResult, err := rpc.AuthClient.DeliverTokenByRPC(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", authResult)
	fmt.Println("认证测试完成")
	req2 := &auth.VerifyTokenReq{
		Token: authResult.GetToken(),
	}

	verifyResult, err := rpc.AuthClient.VerifyTokenByRPC(ctx, req2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", verifyResult)

	// fmt.Println("测试认证失败")

	// req3 := &auth.VerifyTokenReq{
	// 	Token: "1111111111111",
	// }

	// verifyResult3, err := rpc.AuthClient.VerifyTokenByRPC(ctx, req3)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%v\n", verifyResult3)

	// 测试hmac

	// userId := int64(444)

	// req := &auth.DeliverTokenReq{UserId: userId, Timestamp: time.Now().Unix(), Signature: "22222222"}
	// fmt.Printf("%v", req)
	// authResult, err := rpc.AuthClient.DeliverTokenByRPC(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", authResult)
}
