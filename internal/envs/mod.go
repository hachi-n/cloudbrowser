package envs

import (
	"github.com/hachi-n/cloudbrowser/internal/utils"
	"os"
	"strings"
)

var appEnv *string

func AppEnv() string {
	if appEnv != nil {
		return *appEnv
	}

	_appEnv := strings.ToLower(os.Getenv("APP_ENV"))
	allowAppEnvs := []string{"development", "staging", "production"}
	if !utils.SliceContains(allowAppEnvs, _appEnv) {
		return ""
	}

	appEnv = &_appEnv
	return _appEnv
}
