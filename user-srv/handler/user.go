package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"lcb123/user-srv/model"
	user "lcb123/user-srv/proto/user"
)

type User struct{}

func (e *User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.Response) error {
	log.Log("Received User.Call request")
	rsp.Msg = "Hello "
	return nil
}
func (e *User) Logout(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received User.Call request")
	rsp.Msg = "Hello "
	return nil
}
func (e *User) UserDetail(ctx context.Context, req *user.UserRequest, rsp *user.UserInfo) error {

	log.Info("Received Service.GetOne request")

	admin := model.AdminUser{}
	user, err := admin.GetByID(req.Uid)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	rsp.Uid = req.Uid
	rsp.Username = user.UserName
	rsp.Avatar = "user.Avatar"
	return nil
}
