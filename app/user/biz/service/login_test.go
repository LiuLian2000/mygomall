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

func TestLogin_Run(t *testing.T) {
	rpc.InitClient()
	dal.Init()

	// init req and assert value

	type testCase struct {
		name         string
		req          *user.LoginReq
		expectedErr  error
		expectedResp *user.LoginResp
	}

	// 定义多组测试数据
	testCases := []testCase{
		{
			name: "success",
			req: &user.LoginReq{
				Email:    "123",
				Password: "123",
			},
			expectedErr: nil,
			expectedResp: &user.LoginResp{
				UserId: 4,
				Token:  "c74c9f5e-d384-4c1d-ad45-a9bc7334166c",
			},
		},
		{
			name: "user not existed",
			req: &user.LoginReq{
				Email:    "567354345",
				Password: "123",
			},
			expectedErr:  errno.UserNotExistErr,
			expectedResp: &user.LoginResp{},
		},
		{
			name: "wrong password",
			req: &user.LoginReq{
				Email:    "123",
				Password: "555",
			},
			expectedErr:  errno.WrongPasswordErr,
			expectedResp: &user.LoginResp{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := NewLoginService(context.Background()).Run(tc.req)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedResp, resp)
		})
	}

}
