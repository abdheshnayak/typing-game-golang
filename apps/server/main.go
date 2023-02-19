package main

import (
	"github.com/abdheshnayak/typing-game/apps/server/internal/framework"
	"go.uber.org/fx"
)

func main() {
	fx.New(framework.Module).Run()
}
