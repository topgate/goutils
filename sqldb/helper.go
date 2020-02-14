package sqldb

import (
	"context"
	"database/sql"
)

// RunInTransaction 内部でトランザクションを開始し、受け取った関数内でエラーが発生した場合ロールバック、成功したらコミットする
func RunInTransaction(db *sql.DB, runFunc func(tx *sql.Tx) error) error {
	return runInTransaction(func() (*sql.Tx, error) {
		return db.Begin()
	}, runFunc)
}

// RunInTransactionTx 内部でトランザクションを開始し、受け取った関数内でエラーが発生した場合ロールバック、成功したらコミットする
// この関数は内部でdb.Beginの代わりに、db.BeginTxが呼び出される
func RunInTransactionTx(ctx context.Context, db *sql.DB, opts *sql.TxOptions, runFunc func(tx *sql.Tx) error) error {
	return runInTransaction(func() (*sql.Tx, error) {
		return db.BeginTx(ctx, opts)
	}, runFunc)
}

func runInTransaction(beginFunc func() (*sql.Tx, error), runFunc func(tx *sql.Tx) error) error {
	tx, err := beginFunc()
	if err != nil {
		return err
	}
	defer func() {
		err := recover()
		if err != nil {
			_ = tx.Rollback()
			panic(err)
		}
	}()
	if err := runFunc(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
