package retrytx

import (
	"context"

	"github.com/cenkalti/backoff"
)

func retry(ctx context.Context, worker func() error, opts ...Option) (err error) {
	options := prepareOptions(opts...)

	return backoff.Retry(worker, backoff.WithContext(options.BackOff, ctx))
}
