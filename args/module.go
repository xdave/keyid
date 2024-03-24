package args

import "go.uber.org/fx"

var CurrentArgs = NewArgs()

var Module = fx.Module("args",
	fx.Provide(func() *Args { return CurrentArgs }),
)
