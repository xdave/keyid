package printer

import "go.uber.org/fx"

var Module = fx.Module("printer",
	fx.Provide(ProvidePrinter),
)
