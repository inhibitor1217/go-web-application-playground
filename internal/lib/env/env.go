package env

import (
	"net/url"
	"os"
)

type App struct {
	Name  string
	Stage AppStage
	Build string
}

type PublicHttp struct {
	BaseUrl *url.URL
	Port    string
}

type Swagger struct {
	BaseUrl *url.URL
	Port    string
}

type Env struct {
	App        App
	PublicHttp PublicHttp
	Swagger    Swagger
}

func FromEnvVars() (*Env, error) {
	appName := os.Getenv("APP_NAME")
	appStage, err := parseAppStage(os.Getenv("APP_STAGE"))
	if err != nil {
		return nil, err
	}
	appBuild := os.Getenv("APP_BUILD")

	publicHttpBaseUrl, err := url.Parse(os.Getenv("PUBLIC_HTTP_BASE_URL"))
	if err != nil {
		return nil, err
	}
	publicHttpPort := os.Getenv("PUBLIC_HTTP_PORT")

	swaggerBaseUrl, err := url.Parse(os.Getenv("SWAGGER_BASE_URL"))
	if err != nil {
		return nil, err
	}
	swaggerPort := os.Getenv("SWAGGER_PORT")

	return &Env{
		App: App{
			Name:  appName,
			Stage: appStage,
			Build: appBuild,
		},
		PublicHttp: PublicHttp{
			BaseUrl: publicHttpBaseUrl,
			Port:    publicHttpPort,
		},
		Swagger: Swagger{
			BaseUrl: swaggerBaseUrl,
			Port:    swaggerPort,
		},
	}, nil
}

func (e *Env) IsDevelopment() bool {
	return e.App.Stage == AppStageDev || e.App.Stage == AppStageTest
}

func (e *Env) IsProduction() bool {
	return e.App.Stage == AppStageProd
}
