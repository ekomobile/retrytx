package retrytx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// RetryTxx retries worker while it returns error.
// Each retry starts with new transaction.
func RetryTxx(
	ctx context.Context,
	db *sqlx.DB,
	worker func(context.Context, *sqlx.Tx) error,
	opts ...Option,
) (err error) {
	txWrapper := func() (err error) {
		var tx *sqlx.Tx
		tx, err = db.BeginTxx(ctx, nil)
		if err != nil {
			return
		}

		err = worker(ctx, tx)
		if err != nil {
			rollbackTxx(tx)
			return
		}

		err = tx.Commit()
		if err != nil {
			rollbackTxx(tx)
		}

		return
	}

	err = retry(ctx, txWrapper, opts...)

	return
}

func rollbackTxx(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		// todo inject some logger
		// logrus.Errorf("could not roll back: %v", err)
	}
}
