package env

import (
	"os"
)

type App struct {
	Name  string
	Stage AppStage
}

type Env struct {
	App App
}

func FromEnvVars() (*Env, error) {
	appName := os.Getenv("APP_NAME")
	appStage, err := parseAppStage(os.Getenv("APP_STAGE"))
	if err != nil {
		return nil, err
	}

	return &Env{
		App: App{
			Name:  appName,
			Stage: appStage,
		},
	}, nil
}
