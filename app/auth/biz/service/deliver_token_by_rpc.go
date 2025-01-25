package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/google/uuid"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	resp = new(auth.DeliveryResp)
	ok := utils.VerifyHMAC(strconv.FormatInt(req.GetUserId(), 10), req.GetTimestamp(), string(req.GetSignature()))
	if !ok {
		err = errno.UnauthorizedDeliverRequestErr
		return
	}
	// userid已经存在就获取，不存在就创建
	token, err := model.GetWithFunc(s.ctx, model.U2TPrefix+strconv.FormatUint(uint64(req.UserId), 10),
		func(ctx context.Context, key string) (string, error) {
			token := uuid.New().String()
			model.Write(ctx, model.T2UPrefix+token, strconv.FormatInt(req.UserId, 10))
			return token, nil
		})

	if err != nil {
		return
	}
	fmt.Printf("\n%v", token)
	return &auth.DeliveryResp{Token: token}, nil
}
