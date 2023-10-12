package repository

import (
	"referral-system/entity"

	"github.com/jmoiron/sqlx"
)

type referralRepo struct {
	db *sqlx.DB
}

func NewReferralRepository(db *sqlx.DB) IReferralRepository {
	return &referralRepo{
		db: db,
	}
}

func (r *referralRepo) InsertReferral(referral entity.Referral) (err error) {
	query := `INSERT INTO referrals(referral_code, user_id, expired_date) VALUES(:referral_code, :user_id, :expired_date)`

	_, err = r.db.NamedExec(query, referral)

	if err != nil {
		return
	}

	return
}

func (r *referralRepo) GetReferral(referralCode string) (referral entity.Referral, err error) {
	query := `SELECT * FROM referrals WHERE referral_code=?`

	err = r.db.Get(&referral, query, referralCode)

	return
}
