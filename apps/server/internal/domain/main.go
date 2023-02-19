package domain

import (
	"github.com/abdheshnayak/typing-game/grpc-interfaces/rpc/user"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module("domain",
	fx.Provide(fxRPCServer),
	fx.Invoke(func(server *grpc.Server, userServer user.UserServer) {
		user.RegisterUserServer(server, userServer)
	}),
)
