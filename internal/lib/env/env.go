package env

import (
	"fmt"
	"net/url"
	"os"
)

type App struct {
	Name  string
	Stage AppStage
	Build string
}

type PSQL struct {
	Host     string
	User     string
	Password string
	Database string
}

func (p PSQL) DatasourceName() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=require",
		p.Host,
		p.User,
		p.Password,
		p.Database,
	)
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
	PSQL       PSQL
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

	psqlHost := os.Getenv("PSQL_HOST")
	psqlUser := os.Getenv("PSQL_USER")
	psqlPassword := os.Getenv("PSQL_PASSWORD")
	psqlDatabase := os.Getenv("PSQL_DATABASE")

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
		PSQL: PSQL{
			Host:     psqlHost,
			User:     psqlUser,
			Password: psqlPassword,
			Database: psqlDatabase,
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
