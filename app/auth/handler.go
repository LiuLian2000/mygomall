package main

import (
	"context"
	"fmt"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/service"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp = new(auth.DeliveryResp)

	fmt.Printf("%v", req)
	if len(req.GetSignature()) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	if resp.BaseResp != nil && resp.BaseResp.StatusCode != errno.SuccessCode {
		return
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	fmt.Printf("\n%v", resp)
	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp = new(auth.VerifyResp)
	if len(req.GetToken()) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	if resp.BaseResp != nil && resp.BaseResp.StatusCode != errno.SuccessCode {
		return
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)

	return resp, err
}
