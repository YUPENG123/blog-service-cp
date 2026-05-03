package global

import (
	"github.com/YUPENG123/blog-service-cp/pkg/logger"
	"github.com/YUPENG123/blog-service-cp/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
