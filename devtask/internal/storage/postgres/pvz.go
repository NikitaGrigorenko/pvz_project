//go:generate mockgen -source ./pvz.go -destination=./mocks/pvz.go -package=mock_pvz
package postgres

import (
	"context"
	"devtask/internal/model"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBops interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetPool(_ context.Context) *pgxpool.Pool
}

type PVZRepo struct {
	db DBops
}

func NewPVZs(database DBops) *PVZRepo {
	return &PVZRepo{db: database}
}

func (r *PVZRepo) AddPVZ(ctx context.Context, pvz *model.PVZ) (id int64, err error) {
	err = r.db.ExecQueryRow(ctx, `INSERT INTO pvz(name,address,contact) VALUES ($1,$2, $3) RETURNING id;`, pvz.Name, pvz.Address, pvz.Contact).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *PVZRepo) Update(ctx context.Context, pvz *model.PVZ, id int64) (idRet int64, err error) {
	cT, err := r.db.Exec(ctx, `UPDATE pvz SET name=$1, address=$2, contact=$3 WHERE id=$4;`, pvz.Name, pvz.Address, pvz.Contact, id)
	if err != nil {
		return 0, err
	}
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

func (r *PVZRepo) GetPVZ(ctx context.Context, id int64) (*model.PVZ, error) {
	var a model.PVZ
	err := r.db.Get(ctx, &a, `SELECT id,name,address,contact FROM pvz where id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (r *PVZRepo) Delete(ctx context.Context, id int64) (err error) {
	cT, err := r.db.Exec(ctx, `DELETE FROM pvz WHERE id = ($1);`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ErrObjectNotFound
		}
		return err
	}
	if cT.RowsAffected() == 0 {
		return model.ErrNoRowsInResultSet
	}
	return nil
}

func (r *PVZRepo) List(ctx context.Context) ([]model.PVZ, error) {
	var a []model.PVZ
	err := r.db.Select(ctx, &a, `SELECT id,name,address,contact FROM pvz`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}

	if a == nil {
		return nil, model.ErrObjectNotFound
	}
	return a, nil
}
