package config

import (
	"github.com/gorilla/sessions"
)

var sConfig *SessionConfig

type SessionConfig struct {
	SessionStore *sessions.CookieStore
}

func InitSessionConfig(isInProduction bool) {
	conf := GetConfig()
	sConfig = &SessionConfig{}
	sConfig.SessionStore = sessions.NewCookieStore([]byte(conf.GetString("session.secret_key")))
	sConfig.SessionStore.Options.HttpOnly = true
	if isInProduction {
		sConfig.SessionStore.Options.Secure = true
		sConfig.SessionStore.Options.MaxAge = 3600
	} else {
		sConfig.SessionStore.Options.Secure = false
		sConfig.SessionStore.Options.MaxAge = 1000
	}
}

func GetSessionConfig() *SessionConfig {
	return sConfig
}
