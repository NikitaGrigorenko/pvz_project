package db

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	cluster *pgxpool.Pool
}

const key = "transaction"

func newDatabase(cluster *pgxpool.Pool) *Database {
	return &Database{cluster: cluster}
}

func (db Database) GetPool(_ context.Context) *pgxpool.Pool {
	return db.cluster
}

func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (cT pgconn.CommandTag, err error) {
	// TODO: take transaction from the context instead of creating the new one
	err = db.InTransaction(ctx, pgx.ReadOnly, func(ctxTX context.Context) error {
		cT, err = db.cluster.Exec(ctx, query, args...)
		return err
	})
	return cT, err
}

func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) (row pgx.Row) {
	// TODO: take transaction from the context instead of creating the new one
	err := db.InTransaction(ctx, pgx.ReadWrite, func(ctxTX context.Context) error {
		row = db.cluster.QueryRow(ctx, query, args...)
		return nil
	})
	if err != nil {
		return nil
	}
	return row
}

func (db Database) InTransaction(ctx context.Context, accessMode pgx.TxAccessMode, f func(ctxTX context.Context) error) error {
	tx, err := db.GetPool(ctx).BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: accessMode,
	})
	if err != nil {
		return err
	}

	if err := f(context.WithValue(ctx, key, tx)); err != nil {
		errRollback := tx.Rollback(ctx)
		return errors.Join(err, errRollback)
	}

	if err := tx.Commit(ctx); err != nil {
		errRollback := tx.Rollback(ctx)
		return errors.Join(err, errRollback)
	}

	return nil
}
