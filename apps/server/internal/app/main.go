package app

import (
	"github.com/abdheshnayak/typing-game/apps/server/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("app", domain.Module)
