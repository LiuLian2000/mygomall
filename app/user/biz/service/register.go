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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	resp = new(user.RegisterResp)
	// // 查用户是否已经存在
	// _, queryErr := model.GetByEmail(mysql.DB, s.ctx, req.GetEmail())
	// if queryErr != nil {
	// 	//错误不是没查到（查询失败）
	// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 		err = queryErr
	// 		return
	// 	}
	// } else {
	// 	err = errno.UserAlreadyExistErr
	// 	return
	// }
	if req.GetPassword() != req.GetConfirmPassword() {
		err = errno.ConfirmPasswordMismatchErr
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	err = model.Create(mysql.DB, s.ctx, newUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err = errno.UserAlreadyExistErr
		}
		return
	}

	//TODO 获取token
	signature := utils.GenerateHMAC(strconv.FormatInt(newUser.ID, 10), time.Now().Unix())
	authResult, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: newUser.ID, Timestamp: time.Now().Unix(), Signature: signature})
	if err != nil {
		err = errno.RpcErr
		return
	}
	if authResult.BaseResp.StatusCode != errno.SuccessCode {
		resp.BaseResp = authResult.BaseResp
		return
	}

	token := authResult.Token
	return &user.RegisterResp{UserId: newUser.ID, Token: token}, nil
}
