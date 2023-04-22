package healthcheck

type HealthcheckView struct {
	AppName  string `json:"app_name" binding:"required"`
	AppStage string `json:"app_stage" binding:"required"`
	AppBuild string `json:"app_build" binding:"required"`
}
