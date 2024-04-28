package db

import (
	"context"
	"devtask/internal/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
)

func NewDb(ctx context.Context, fileName string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, generateDsn(fileName))
	if err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}

func generateDsn(fileName string) string {
	configData, err := config.Read(fileName)
	if err != nil {
		return ""
	}

	port, _ := strconv.ParseInt(configData.Port, 10, 64)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configData.Host, port, configData.User, configData.DbInfo.Password, configData.Name)
}
