package txhelper

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//go:generate go run gen/readme/main.go

// Operation defines the operation that will be performed on the database.
type Operation int

const (
	// ReadOperation should be used for read only activities, such as selecting data.
	ReadOperation Operation = 1 << iota
	// InsertOperation should be used for the INSERT operation.
	InsertOperation
	// UpdateOperation should be used for the UPDATE operation.
	UpdateOperation
	// DeleteOperation should be used for the DELETE operation.
	DeleteOperation
)

// Has checks if the Operation flag has the specified Operation.
func (flag Operation) Has(b Operation) bool {
	return b&flag != 0
}

// TxHelper is a helper that can be used inside mappers. Its purpose is to assist the mappers with transaction
// handling.
type TxHelper struct {
	pool    *pgxpool.Pool
	txn     pgx.Tx
	mu      sync.Mutex
	options *Options
}

// Executor is a type that implements the Exec, Query and QueryRow of pgx, it can be used to hide the types
// pgxpool.Pool, pgx.Conn,...
type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Options can be used to finetune txhelper.
type Options struct {
	Hooks Hooks
}

// Hooks can be used to perform custom logic before or after certain events.
type Hooks struct {
	BeforeStartTransaction func(*TxHelper, *pgxpool.Pool, context.Context) error
	AfterStartTransaction  func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error
	BeforeCommit           func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error
	AfterCommit            func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error
	BeforeRollback         func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error
	AfterRollback          func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error
}

// New creates a new TxHelper that can be used inside mappers. Its purpose is to assist the mappers with transaction
// handling.
func New(pool *pgxpool.Pool, options *Options) (*TxHelper, error) {
	tx := &TxHelper{
		pool:    pool,
		options: options,
		txn:     nil,
	}
	tx.setOptions()
	return tx, nil
}

func (tx *TxHelper) setOptions() {
	if tx.options == nil {
		tx.options = &Options{}
	}
	if tx.options.Hooks.BeforeStartTransaction == nil {
		tx.options.Hooks.BeforeStartTransaction = func(*TxHelper, *pgxpool.Pool, context.Context) error {
			return nil
		}
	}
	if tx.options.Hooks.AfterStartTransaction == nil {
		tx.options.Hooks.AfterStartTransaction = func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error {
			return nil
		}
	}
	if tx.options.Hooks.BeforeCommit == nil {
		tx.options.Hooks.BeforeCommit = func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error {
			return nil
		}
	}
	if tx.options.Hooks.AfterCommit == nil {
		tx.options.Hooks.AfterCommit = func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error {
			return nil
		}
	}
	if tx.options.Hooks.BeforeRollback == nil {
		tx.options.Hooks.BeforeRollback = func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error {
			return nil
		}
	}
	if tx.options.Hooks.AfterRollback == nil {
		tx.options.Hooks.AfterRollback = func(*TxHelper, *pgxpool.Pool, pgx.Tx, context.Context) error {
			return nil
		}
	}
}

// GetExecutor gets the executor for the specified operation.
// The executor can be used to query the database.
// GetExecutor will start and return a transaction if there is no existing transaction AND if the operation is either a
// InsertOperation, UpdateOperation or DeleteOperation.
// In case of a non existing transaction and a ReadOperation GetExecutor will NOT start a transaction.
func (tx *TxHelper) GetExecutor(ctx context.Context, operations Operation) (Executor, error) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	// we already started a transaction
	if tx.txn != nil {
		return tx.txn, nil
	}
	if operations.Has(InsertOperation) || operations.Has(UpdateOperation) || operations.Has(DeleteOperation) {
		if err := tx.options.Hooks.BeforeStartTransaction(tx, tx.pool, ctx); err != nil {
			return nil, err
		}
		var err error
		tx.txn, err = tx.pool.Begin(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to start transaction: %w", err)
		}
		if err := tx.options.Hooks.AfterStartTransaction(tx, tx.pool, tx.txn, ctx); err != nil {
			return nil, err
		}
		return tx.txn, nil
	}
	return tx.pool, nil
}

//nolint:dupl //false positive with Rollback() function
// Commit commits the transaction (if there is one).
func (tx *TxHelper) Commit(ctx context.Context) error {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.txn == nil {
		return nil
	}
	if err := tx.options.Hooks.BeforeCommit(tx, tx.pool, tx.txn, ctx); err != nil {
		return err
	}
	if err := tx.txn.Commit(ctx); err != nil {
		return err
	}
	err := tx.options.Hooks.AfterCommit(tx, tx.pool, tx.txn, ctx)
	tx.txn = nil
	return err
}

//nolint:dupl //false positive with Commit() function
// Rollback rollbacks the transaction (if there is one), remember that you only need to Rollback if there was no error
// before.
func (tx *TxHelper) Rollback(ctx context.Context) error {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.txn == nil {
		return nil
	}
	if err := tx.options.Hooks.BeforeRollback(tx, tx.pool, tx.txn, ctx); err != nil {
		return err
	}

	if err := tx.txn.Rollback(ctx); err != nil {
		return err
	}

	err := tx.options.Hooks.AfterRollback(tx, tx.pool, tx.txn, ctx)
	tx.txn = nil
	return err
}
