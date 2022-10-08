package main

import (
	"MorselShogiew/Users-service-rest/application"
	"MorselShogiew/Users-service-rest/config"
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/logger/opt"
	"MorselShogiew/Users-service-rest/provider"
	"MorselShogiew/Users-service-rest/repos"
	"MorselShogiew/Users-service-rest/service/api"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func main() {
	conf := config.LoadConfig()
	conf.InstanceID = uuid.New()
	opts := makeLoggerOpts(conf)
	l := logger.New(opts)
	p := provider.New(conf, l)

	repositories := repos.New(p, l)

	UserService := api.New(l, repositories)

	app := application.New(conf, l, UserService)
	app.Start()
}

func makeLoggerOpts(c *config.Config) *opt.LoggerOpts {
	return &opt.LoggerOpts{
		Opts: &opt.GeneralOpts{
			InstanceID: c.InstanceID,
			Env:        c.Environment,
			AppName:    c.ApplicationName,
			Level:      c.Logger.Level,
		},
		StdLoggerOpts: &opt.StdLoggerOpts{
			LogFile:  c.Logger.LoggerStd.LogFile,
			Stdout:   c.Logger.LoggerStd.Stdout,
			Disabled: c.Logger.LoggerStd.Disabled,
		},
	}
}
