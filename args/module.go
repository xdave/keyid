package args

import "go.uber.org/fx"

var Module = fx.Module("args",
	fx.Provide(NewArgs),
)
