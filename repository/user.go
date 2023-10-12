package repository

import (
	"database/sql"
	"referral-system/entity"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Login(user entity.User) (entity.User, error) {
	query := `SELECT * FROM users WHERE email=? and password=?`

	err := r.db.Get(&user, query, user.Email, user.Password)
	return user, err
}

func (r *userRepo) Register(user entity.User) (res sql.Result, err error) {
	query := `INSERT INTO users(email, password, username, name, referral_hit, user_type) VALUES(:email, :password, :username, :name, 0, 1)`

	res, err = r.db.NamedExec(query, user)

	if err != nil {
		return
	}

	return
}

func (r *userRepo) Update(user entity.User) (err error) {
	query := `UPDATE users SET email=:email, password=:password, name=:name WHERE email=:email`

	_, err = r.db.NamedExec(query, user)

	if err != nil {
		return
	}

	return
}

func (r *userRepo) UpdateReferralHit(ID int) (err error) {
	query := `UPDATE users SET referral_hit = referral_hit + 1 WHERE id = ?`

	_, err = r.db.Exec(query, ID)
	return
}
