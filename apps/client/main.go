package main

import (
	"github.com/abdheshnayak/typing-game/apps/client/internal/framework"
	"go.uber.org/fx"
)

func main() {
	fx.New(framework.Module,fx.NopLogger).Run()
}
