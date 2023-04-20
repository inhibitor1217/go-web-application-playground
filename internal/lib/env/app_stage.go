package env

import "fmt"

type AppStage string

const (
	AppStageDev AppStage = "development"
)

func parseAppStage(s string) (AppStage, error) {
	switch s {
	case "development":
		return AppStageDev, nil
	default:
		return "", fmt.Errorf("invalid app stage: %s", s)
	}
}
