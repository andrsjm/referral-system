package repository

import (
	"referral-system/entity"

	"github.com/jmoiron/sqlx"
)

type contributorRepo struct {
	db *sqlx.DB
}

func NewContributorRepository(db *sqlx.DB) IContributorRepository {
	return &contributorRepo{
		db: db,
	}
}

func (r *contributorRepo) CheckContributor(email string) (isReady bool, err error) {
	query := `SELECT EXISTS(SELECT * FROM contributors WHERE email = ?)`

	err = r.db.Get(&isReady, query, email)

	return
}

func (r *contributorRepo) InsertContributor(contributor entity.Contributor) (err error) {
	query := `INSERT INTO contributors(referral_code, email) VALUES(:referral_code, :email)`

	_, err = r.db.NamedExec(query, contributor)

	return
}
