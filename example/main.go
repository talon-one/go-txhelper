package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/talon-one/go-txhelper"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// create a users table
	_, err = pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	txh, err := txhelper.New(pool, nil)
	if err != nil {
		panic(err)
	}
	// we need to close the helper at the end of our business logic
	defer func() {
		if err := txh.Commit(ctx); err != nil {
			panic(err)
		}
	}()

	um := NewUserMapper(txh)

	id := 1
	wantedName := "Joe"

	// Get the UserName of the specified user
	currentName, err := um.GetUserName(ctx, id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			panic(err)
		}
		// if there are no rows for this id create a user
		if err := um.CreateUser(ctx, id, wantedName); err != nil {
			panic(err)
		}
		fmt.Printf("Created User (id=%d, name=%s)\n", id, wantedName)
		return
	}
	if currentName != wantedName {
		if err := um.UpdateUserName(ctx, id, wantedName); err != nil {
			panic(err)
		}
		fmt.Printf("Updated User (id=%d, name=%s)\n", id, wantedName)
		return
	}
	fmt.Printf("nothing to do (id=%d, name=%s)\n", id, currentName)
}

type UserMapper struct {
	txh *txhelper.TxHelper
}

func NewUserMapper(txh *txhelper.TxHelper) *UserMapper {
	return &UserMapper{
		txh: txh,
	}
}

func (um *UserMapper) CreateUser(ctx context.Context, id int, name string) error {
	executor, err := um.txh.GetExecutor(ctx, txhelper.InsertOperation)
	if err != nil {
		return err
	}

	_, err = executor.Exec(ctx, "INSERT INTO users VALUES ($1, $2)", id, name)
	return err
}

func (um *UserMapper) GetUserName(ctx context.Context, id int) (string, error) {
	executor, err := um.txh.GetExecutor(ctx, txhelper.ReadOperation)
	if err != nil {
		return "", err
	}

	row := executor.QueryRow(ctx, "SELECT name FROM users WHERE id = $1", id)
	var name string
	if err := row.Scan(&name); err != nil {
		return "", err
	}

	return name, nil
}

func (um *UserMapper) UpdateUserName(ctx context.Context, id int, newName string) error {
	executor, err := um.txh.GetExecutor(ctx, txhelper.UpdateOperation)
	if err != nil {
		return err
	}

	_, err = executor.Exec(ctx, "UPDATE users SET name = $1 WHERE id = $2", newName, id)
	return err
}
