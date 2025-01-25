package main

import (
	"context"
	"regexp"

	"github.com/Group-lifelong-youth-training/mygomall/app/user/biz/service"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = new(user.RegisterResp)
	pattern := "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$"
	re, _ := regexp.Compile(pattern)
	if len(req.GetEmail()) == 0 || len(req.GetPassword()) == 0 || len(req.GetConfirmPassword()) == 0 || req.GetConfirmPassword() != req.GetPassword() || !re.Match([]byte(req.GetEmail())) {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp, err = service.NewRegisterService(ctx).Run(req)
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

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = new(user.LoginResp)
	resp, err = service.NewLoginService(ctx).Run(req)
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
