//go:build integration

package tests

import (
	"context"
	"devtask/internal/pkg/db"
	"fmt"
	"testing"
)

type TDB struct {
	DB db.Database
}

func NewFromEnv() (*TDB, error) {
	database, err := db.NewDb(context.Background(), "test_config.json")
	if err != nil {
		return nil, err
	}
	return &TDB{DB: *database}, nil
}

func (d *TDB) SetUp(t *testing.T, tableName string) error {
	t.Helper()
	err := d.truncateTable(context.Background(), tableName)
	return err
}

func (d *TDB) TearDown() {
}

func (d *TDB) truncateTable(ctx context.Context, tableName string) error {
	q := fmt.Sprintf("TRUNCATE table %s RESTART IDENTITY", tableName)
	if _, err := d.DB.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}
