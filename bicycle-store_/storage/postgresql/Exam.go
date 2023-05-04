package postgresql

import (
	"app/api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type R_Repo struct {
	db *pgxpool.Pool
}

func NewCodeRepo(db *pgxpool.Pool) *R_Repo {
	return &R_Repo{db: db}
}
func (r R_Repo) Create(ctx context.Context, req *models.PromoCreate) (int, error) {
	var id int

	query := `INSERT INTO
    promo(promo_name,is_percent,discount,order_limit_price)
    values($1,$2,$3,$4) RETURNING promo_id`

	fmt.Println(query)

	err := r.db.QueryRow(ctx, query,
		req.PromoName,
		req.IsPercent,
		req.Discount,
		req.Limit_Price,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r R_Repo) GetByID(ctx context.Context, req *models.PromoPrimaryKey) (*models.Promo, error) {
	var (
		query string
		promo models.Promo
	)
	query = `select promo_id , promo_name , is_percent , discount , order_limit_price from promo WHERE promo_id = $1`

	err := r.db.QueryRow(ctx, query, req.Promo_id).Scan(
		&promo.Promo_id,
		&promo.PromoName,
		&promo.IsPercent,
		&promo.Discount,
		&promo.Limit_Price,
	)

	if err != nil {
		return nil, err
	}

	return &promo, nil

}

func (r R_Repo) GetList(ctx context.Context, req *models.Query) (resp []models.Promo, err error) {
	query := `select promo_id , promo_name , is_percent , discount , order_limit_price from promo OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(ctx, query, req.Offset, req.Limit)

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.Promo
		err = rows.Scan(
			&a.Promo_id,
			&a.PromoName,
			&a.IsPercent,
			&a.Discount,
			&a.Limit_Price,
		)

		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r R_Repo) Delete(ctx context.Context, req *models.PromoPrimaryKey) (int64, error) {
	query := `DELETE FROM promo WHERE promo_id = $1`

	result, err := r.db.Exec(ctx, query, req.Promo_id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil

}
