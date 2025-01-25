package service

import (
	"context"
	"testing"

	"github.com/Group-lifelong-youth-training/mygomall/app/user/biz/dal"
	"github.com/Group-lifelong-youth-training/mygomall/app/user/infra/rpc"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Run(t *testing.T) {
	rpc.InitClient()
	dal.Init()

	// init req and assert value

	type testCase struct {
		name         string
		req          *user.RegisterReq
		expectedErr  error
		expectedResp *user.RegisterResp
	}

	// 定义多组测试数据
	testCases := []testCase{
		{
			name: "success",
			req: &user.RegisterReq{
				Email:           "4588465",
				Password:        "123",
				ConfirmPassword: "123",
			},
			expectedErr:  nil,
			expectedResp: &user.RegisterResp{},
		},
		{
			name: "mismatch confirmpassword",
			req: &user.RegisterReq{
				Email:           "234",
				Password:        "234",
				ConfirmPassword: "245",
			},
			expectedErr:  errno.ConfirmPasswordMismatchErr,
			expectedResp: &user.RegisterResp{},
		},
		{
			name: "already exist",
			req: &user.RegisterReq{
				Email:           "123",
				Password:        "234",
				ConfirmPassword: "234",
			},
			expectedErr:  errno.UserAlreadyExistErr,
			expectedResp: &user.RegisterResp{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := NewRegisterService(context.Background()).Run(tc.req)
			if tc.name == "success" {
				assert.NotNil(t, resp.GetUserId())
				assert.NotNil(t, resp.GetToken())
			}
			assert.Equal(t, tc.expectedErr, err)
		})
	}

}
