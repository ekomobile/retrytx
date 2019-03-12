package retrytx

import (
	"context"
	"database/sql"
)

// RetryTx retries worker while it returns error.
// Each retry starts with new transaction.
func RetryTx(
	ctx context.Context,
	db *sql.DB,
	worker func(context.Context, *sql.Tx) error,
	opts ...Option,
) (err error) {
	txWrapper := func() (err error) {
		var tx *sql.Tx
		tx, err = db.BeginTx(ctx, nil)
		if err != nil {
			return
		}

		err = worker(ctx, tx)
		if err != nil {
			rollbackTx(tx)
			return
		}

		err = tx.Commit()
		if err != nil {
			rollbackTx(tx)
		}

		return
	}

	err = retry(ctx, txWrapper, opts...)

	return
}

func rollbackTx(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		// todo inject some logger
		// logrus.Errorf("could not roll back: %v", err)
	}
}
