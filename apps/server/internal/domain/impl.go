package domain

import (
	"context"

	"github.com/abdheshnayak/typing-game/grpc-interfaces/rpc/user"
)

type domainI struct {
	user.UnimplementedUserServer
}

func (d *domainI) GetUserByEmail(_ context.Context, input *user.GetUserByEmailIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserById(_ context.Context, input *user.GetUserByIdIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserByName(_ context.Context, input *user.GetUserByNameIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserBySession(_ context.Context, input *user.SessionIdIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) RegisterUser(_ context.Context, input *user.RegisterUserIn) (*user.GetUserOut, error) {
	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) UserLogin(_ context.Context, input *user.LoginIn) (*user.SessionOut, error) {
	return &user.SessionOut{
		Session: "hi",
		User: &user.GetUserOut{
			Username: "sample",
			Email:    "sample@gmail.com",
			Id:       "user_sample",
			Status:   "active",
		},
	}, nil
}

func (d *domainI) UserLogOut(_ context.Context, input *user.SessionIdIn) (*user.BoolOut, error) {
	return &user.BoolOut{
		IsSuccess: false,
	}, nil
}

func fxRPCServer() user.UserServer {
	return &domainI{}
}
