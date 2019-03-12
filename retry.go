package retrytx

import (
	"context"

	"github.com/cenkalti/backoff"
)

func retry(ctx context.Context, worker func() error, opts ...Option) (err error) {
	options := prepareOptions(opts...)

	var b backoff.BackOff

	if options.WithConstantBackOff {
		b = backoff.NewConstantBackOff(options.ConstantBackOffInterval)

	} else {
		b = backoff.NewExponentialBackOff()
	}

	if options.WithRetries {
		b = backoff.WithMaxRetries(b, options.MaxRetries)
	}

	err = backoff.Retry(worker, backoff.WithContext(b, ctx))

	return
}
