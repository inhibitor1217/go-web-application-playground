package env

import (
	"os"
)

type App struct {
	Name  string
	Stage AppStage
}

type Http struct {
	Port string
}

type Env struct {
	App  App
	Http Http
}

func FromEnvVars() (*Env, error) {
	appName := os.Getenv("APP_NAME")
	appStage, err := parseAppStage(os.Getenv("APP_STAGE"))
	if err != nil {
		return nil, err
	}
	httpPort := os.Getenv("HTTP_PORT")

	return &Env{
		App: App{
			Name:  appName,
			Stage: appStage,
		},
		Http: Http{
			Port: httpPort,
		},
	}, nil
}

func (e *Env) IsDevelopment() bool {
	return e.App.Stage == AppStageDev || e.App.Stage == AppStageTest
}

func (e *Env) IsProduction() bool {
	return e.App.Stage == AppStageProd
}
