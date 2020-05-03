package rest

import "github.com/zees-dev/snippetbox/pkg/config"

type App struct {
	HTMLDir   string
	StaticDir string
}

func CreateApp(cfg *config.Server) *App {
	return &App{HTMLDir: cfg.HTMLDir, StaticDir: cfg.StaticDir}
}
