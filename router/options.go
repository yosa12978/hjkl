package router

import (
	"os"

	"github.com/yosa12978/hjkl/logging"
)

type optionFunc func(*options)

type options struct {
	logger logging.Logger
}

func defaultOptions() options {
	return options{
		logger: logging.NewJsonLogger(os.Stdout),
	}
}

func newOptions(opts ...optionFunc) options {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithLogger(logger logging.Logger) optionFunc {
	return func(opt *options) {
		opt.logger = logger
	}
}
