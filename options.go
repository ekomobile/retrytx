package retrytx

import (
	"github.com/cenkalti/backoff"
)

type (
	Options struct {
		BackOff backoff.BackOff
	}

	Option func(*Options)
)

func WithBackOff(b backoff.BackOff) Option {
	return func(opts *Options) {
		opts.BackOff = b
	}
}

func prepareOptions(opts ...Option) (options *Options) {
	options = &Options{}

	for _, o := range opts {
		o(options)
	}

	return
}
