package domain

import (
	"context"

	"github.com/abdheshnayak/typing-game/grpc-interfaces/rpc/user"
)

type domainI struct {
	userClient user.UserClient
}

func (d *domainI) GetUserById(ctx context.Context, userId string) (*user.GetUserOut, error) {
	return d.userClient.GetUserById(ctx, &user.GetUserByIdIn{
		Id: userId,
	})

}

func (d *domainI) UserDebugLog(msg string) {
	ctx := context.TODO()
	go func() {
		d.userClient.UserDebugLog(ctx, &user.LogsIn{
			Msg: msg,
		})
	}()
}

func fxRPCClient(userClient user.UserClient) Domain {
	return &domainI{
		userClient: userClient,
	}
}
