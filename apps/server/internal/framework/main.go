package framework

import (
	"github.com/abdheshnayak/typing-game/apps/server/internal/app"
	"github.com/abdheshnayak/typing-game/pkg/grpcfx"
	"go.uber.org/fx"
)

var Module = fx.Module("framework",
	grpcfx.FxGrpcServer,
	app.Module,
)
