package retrytx

import (
	"time"
)

type (
	Options struct {
		WithConstantBackOff     bool
		ConstantBackOffInterval time.Duration
		WithRetries             bool
		MaxRetries              uint64
	}

	Option func(*Options)
)

func WithMaxRetries(retries uint64) Option {
	if retries < 0 {
		retries = 0
	}

	return func(opts *Options) {
		opts.WithRetries = true
		opts.MaxRetries = retries
	}
}

func WithConstantBackOff(interval time.Duration) Option {
	return func(opts *Options) {
		opts.WithConstantBackOff = true
		opts.ConstantBackOffInterval = interval
	}
}

func prepareOptions(opts ...Option) (options *Options) {
	options = &Options{}

	for _, o := range opts {
		o(options)
	}

	return
}
