package env

import (
	"fmt"
	"net/url"
	"os"
)

type App struct {
	Domain string
	Name   string
	Stage  AppStage
	Build  string
}

type Auth struct {
	JwtSecret string
}

type PSQL struct {
	Host     string
	User     string
	Password string
	Database string
	SSL      bool
}

func (p PSQL) DatasourceName() string {
	sslmode := "disable"
	if p.SSL {
		sslmode = "verify-full"
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host,
		p.User,
		p.Password,
		p.Database,
		sslmode,
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
	Auth       Auth
	PSQL       PSQL
	PublicHttp PublicHttp
	Swagger    Swagger
}

func FromEnvVars() (*Env, error) {
	appDomain := os.Getenv("APP_DOMAIN")
	appName := os.Getenv("APP_NAME")
	appStage, err := parseAppStage(os.Getenv("APP_STAGE"))
	if err != nil {
		return nil, err
	}
	appBuild := os.Getenv("APP_BUILD")

	authJwtSecret := os.Getenv("AUTH_JWT_SECRET")

	psqlHost := os.Getenv("PSQL_HOST")
	psqlUser := os.Getenv("PSQL_USER")
	psqlPassword := os.Getenv("PSQL_PASSWORD")
	psqlDatabase := os.Getenv("PSQL_DATABASE")
	psqlSSL := os.Getenv("PSQL_SSL") == "true"

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
			Domain: appDomain,
			Name:   appName,
			Stage:  appStage,
			Build:  appBuild,
		},
		Auth: Auth{
			JwtSecret: authJwtSecret,
		},
		PSQL: PSQL{
			Host:     psqlHost,
			User:     psqlUser,
			Password: psqlPassword,
			Database: psqlDatabase,
			SSL:      psqlSSL,
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
