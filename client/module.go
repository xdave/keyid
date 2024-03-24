package client

import "go.uber.org/fx"

var Module = fx.Module("client",
	fx.Provide(
		NewRekordboxOptionsResolver,
		NewRekordboxHistory,
		NewRekordboxClient,
	),
)
