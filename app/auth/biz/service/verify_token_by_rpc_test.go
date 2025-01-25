package service

import (
	"context"
	"testing"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	dal.Init()

	type testCase struct {
		name         string
		req          *auth.VerifyTokenReq
		expectedErr  error
		expectedResp *auth.VerifyResp
	}

	testCases := []testCase{
		{
			name: "token not exist",
			req: &auth.VerifyTokenReq{
				Token: "111",
			},
			expectedErr: nil,
			expectedResp: &auth.VerifyResp{
				Res: false,
			},
		},
		{
			name: "token exist",
			req: &auth.VerifyTokenReq{
				Token: "30dc8853-081e-4d3c-9ce8-0f922274661f",
			},
			expectedErr: nil,
			expectedResp: &auth.VerifyResp{
				UserId: 222,
				Res:    true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := NewVerifyTokenByRPCService(context.Background()).Run(tc.req)
			assert.Equal(t, tc.expectedResp, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
