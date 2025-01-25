package service

import (
	"context"
	"testing"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	dal.Init()

	type testCase struct {
		name         string
		req          *auth.DeliverTokenReq
		expectedErr  error
		expectedResp *auth.DeliveryResp
	}

	var userId1 int64 = 668

	timeStamp1 := time.Now().Unix()

	var userId2 int64 = 669
	userId2S := "669"
	timeStamp2 := time.Now().Unix()

	var userId3 int64 = 222
	userId3S := "222"
	timeStamp3 := time.Now().Unix()

	testCases := []testCase{
		{
			name: "Unauthorized",
			req: &auth.DeliverTokenReq{
				UserId:    userId1,
				Timestamp: timeStamp1,
				Signature: "sdfwesrt",
			},
			expectedErr:  errno.UnauthorizedDeliverRequestErr,
			expectedResp: &auth.DeliveryResp{},
		},
		{
			name: "Success",
			req: &auth.DeliverTokenReq{
				UserId:    userId2,
				Timestamp: timeStamp2,
				Signature: utils.GenerateHMAC(userId2S, timeStamp2),
			},
			expectedErr: nil,
			expectedResp: &auth.DeliveryResp{
				Token: "ssss",
			},
		},
		{
			name: "Already Exist",
			req: &auth.DeliverTokenReq{
				UserId:    userId3,
				Timestamp: timeStamp3,
				Signature: utils.GenerateHMAC(userId3S, timeStamp3),
			},
			expectedErr: nil,
			expectedResp: &auth.DeliveryResp{
				Token: "30dc8853-081e-4d3c-9ce8-0f922274661f",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := NewDeliverTokenByRPCService(context.Background()).Run(tc.req)
			assert.Equal(t, tc.expectedErr, err)
			if tc.name == "Success" {
				assert.NotNil(t, resp.Token)
			} else {
				assert.Equal(t, tc.expectedResp, resp)
			}

		})
	}

}
