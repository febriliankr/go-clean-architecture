package repo

import (
	"github.com/febriliankr/go-clean-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (r *UserDB) GetUserList(req entity.GetUserListRequest) (entity.GetUserListResponse, error) {
	var res entity.GetUserListResponse

	query := queryGetAllUsers

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Queryx(query, offset, req.Limit)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp entity.User
		if err := rows.StructScan(&temp); err != nil {
			return res, err
		}
		res.Data.Rows = append(res.Data.Rows, temp)
	}

	return res, nil

}
