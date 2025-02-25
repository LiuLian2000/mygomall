package mw

import (
	"context"

	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func AuthMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		//获取 metadata
		token, ok := metainfo.GetValue(ctx, "userId")

		authResult, err := AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})

		if err != nil {
			return err
		} else if authResult.BaseResp.StatusCode != errno.SuccessCode {
			return errno.NewErrNo(int64(authResult.BaseResp.StatusCode), authResult.BaseResp.StatusMessage)
		} else if authResult.Res != ok {
			return errno.UnauthorizedDeliverRequestErr
		}

		//继续执行业务逻辑
		err = next(ctx, request, response)

		return err
	}
}
