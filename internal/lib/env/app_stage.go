package env

import "fmt"

type AppStage string

const (
	AppStageDev  AppStage = "development"
	AppStageProd AppStage = "production"
	AppStageTest AppStage = "test"
)

func (s AppStage) String() string {
	return string(s)
}

func parseAppStage(s string) (AppStage, error) {
	switch s {
	case "development":
		return AppStageDev, nil
	case "production":
		return AppStageProd, nil
	case "test":
		return AppStageTest, nil
	default:
		return "", fmt.Errorf("invalid app stage: %s", s)
	}
}
