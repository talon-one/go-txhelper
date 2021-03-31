package txhelper

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/stretchr/testify/require"
)

func TestTxHelper(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	require.NoError(t, err)
	defer pool.Close()

	t.Run("first and consequent read operations should return the pool", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		executor1, err := helper.GetExecutor(context.Background(), ReadOperation)
		require.NoError(t, err)
		_, ok := executor1.(*pgxpool.Pool)
		require.True(t, ok)

		executor2, err := helper.GetExecutor(context.Background(), ReadOperation)
		require.NoError(t, err)
		_, ok = executor2.(*pgxpool.Pool)
		require.True(t, ok)

		// should not be needed
		require.NoError(t, helper.Commit(context.Background()))
	})

	t.Run("first insert operation and consequent operations should return a transaction", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		executor1, err := helper.GetExecutor(context.Background(), InsertOperation)
		require.NoError(t, err)
		_, ok := executor1.(*pgxpool.Tx)
		require.True(t, ok)

		executor2, err := helper.GetExecutor(context.Background(), ReadOperation)
		require.NoError(t, err)
		_, ok = executor2.(*pgxpool.Tx)
		require.True(t, ok)

		executor3, err := helper.GetExecutor(context.Background(), UpdateOperation)
		require.NoError(t, err)
		_, ok = executor3.(*pgxpool.Tx)
		require.True(t, ok)

		require.NoError(t, helper.Commit(context.Background()))
	})

	t.Run("commit a transaction", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		_, err = helper.GetExecutor(context.Background(), InsertOperation)
		require.NoError(t, err)
		require.NoError(t, helper.Commit(context.Background()))
	})

	t.Run("commit a non transaction", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		_, err = helper.GetExecutor(context.Background(), ReadOperation)
		require.NoError(t, err)
		require.NoError(t, helper.Commit(context.Background()))
	})

	t.Run("rollback a transaction", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		_, err = helper.GetExecutor(context.Background(), InsertOperation)
		require.NoError(t, err)
		require.NoError(t, helper.Rollback(context.Background()))
	})

	t.Run("rollback a non transaction", func(t *testing.T) {
		helper, err := New(pool, nil)
		require.NoError(t, err)
		_, err = helper.GetExecutor(context.Background(), ReadOperation)
		require.NoError(t, err)
		require.NoError(t, helper.Rollback(context.Background()))
	})
}
