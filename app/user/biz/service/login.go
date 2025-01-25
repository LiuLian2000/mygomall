package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/user/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/user/biz/model"
	"github.com/Group-lifelong-youth-training/mygomall/app/user/infra/rpc"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/errno"
	"github.com/Group-lifelong-youth-training/mygomall/pkg/utils"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/auth"
	"github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	resp = new(user.LoginResp)
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.UserNotExistErr
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			err = errno.WrongPasswordErr
		}
		return
	}
	//获取token
	signature := utils.GenerateHMAC(strconv.FormatInt(userRow.ID, 10), time.Now().Unix())
	authResult, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: userRow.ID, Timestamp: time.Now().Unix(), Signature: signature})
	if err != nil {
		err = errno.RpcErr
		return
	}
	if authResult.BaseResp.StatusCode != errno.SuccessCode {
		resp.BaseResp = authResult.BaseResp
		return
	}

	token := authResult.Token

	return &user.LoginResp{UserId: userRow.ID, Token: token}, nil
}
