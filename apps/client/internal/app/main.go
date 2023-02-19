package app

import (
	"github.com/abdheshnayak/typing-game/apps/client/internal/domain"
	"github.com/abdheshnayak/typing-game/apps/client/internal/ui"
	"go.uber.org/fx"
)

var Module = fx.Module("app",

	domain.Module,
	ui.Module,
)
