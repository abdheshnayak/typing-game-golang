package domain

import (
	"context"

	"github.com/abdheshnayak/typing-game/grpc-interfaces/rpc/user"
)

type Domain interface {
	GetUserById(ctx context.Context, userId string) (*user.GetUserOut, error)

	UserDebugLog(msg string)
}
