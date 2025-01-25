package service

import (
	"context"
	"strconv"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.

	resp = new(auth.VerifyResp)
	userIdStr, ok, err := model.Get(s.ctx, model.T2UPrefix+req.GetToken())
	if err != nil {
		klog.Error(err)
		return
	}

	if !ok {
		klog.Error(err)
		resp.Res = false
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		klog.Error(err)
		return
	}
	return &auth.VerifyResp{Res: ok, UserId: userId}, nil
}
