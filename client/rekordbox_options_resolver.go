package client

import (
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/fx"
)

type RekordboxOptionsResolver struct{}

type RekordboxOptionsResolverParams struct {
	fx.In
}

type RekordboxOptionsResolverResult struct {
	fx.Out
	Resolver *RekordboxOptionsResolver
}

func NewRekordboxOptionsResolver(params RekordboxOptionsResolverParams) RekordboxOptionsResolverResult {
	return RekordboxOptionsResolverResult{
		Resolver: &RekordboxOptionsResolver{},
	}
}

func (r *RekordboxOptionsResolver) Resolve() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "darwin":
		{
			return filepath.Join(homeDir, "Library", "Application Support", "Pioneer", "rekordboxAgent", "storage", "options.json")
		}
	case "windows":
		{
			// Might not work, not tested on Windows
			return filepath.Join(homeDir, "AppData", "Local", "Pioneer", "rekordboxAgent", "storage", "options.json")
		}
	}
	panic("Cannot determine Rekordbox db options path, unsupported OS (mac & windows only)")
}
