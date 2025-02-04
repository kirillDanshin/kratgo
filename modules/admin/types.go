package admin

import (
	"io"

	"github.com/savsgio/kratgo/modules/cache"
	"github.com/savsgio/kratgo/modules/config"
	"github.com/savsgio/kratgo/modules/invalidator"

	"github.com/savsgio/atreugo/v8"
	logger "github.com/savsgio/go-logger"
)

// Config ...
type Config struct {
	FileConfig  config.Admin
	Cache       *cache.Cache
	Invalidator Invalidator

	HTTPScheme string

	LogLevel  string
	LogOutput io.Writer
}

// Admin ...
type Admin struct {
	fileConfig config.Admin

	server      Server
	cache       *cache.Cache
	invalidator Invalidator

	httpScheme string

	log *logger.Logger
}

// ###### INTERFACES ######

// Invalidator ...
type Invalidator interface {
	Start()
	Add(e invalidator.Entry) error
}

// Server ...
type Server interface {
	ListenAndServe() error
	Path(httpMethod string, url string, viewFn atreugo.View)
	SetLogOutput(output io.Writer)
}
