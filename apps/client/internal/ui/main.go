package ui

import (
	"context"
	"fmt"

	"github.com/abdheshnayak/typing-game/apps/client/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("ui",
	fx.Invoke(func(d domain.Domain) {
		ctx := context.TODO()
		resp, err := d.GetUserById(ctx, "hello")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(resp)
	}),
	// HomeModule,
)
