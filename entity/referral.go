package entity

import "time"

type Referral struct {
	ID           int       `json:"id" db:"id"`
	ReferralCode string    `json:"referral_code" db:"referral_code"`
	ExpiredDate  time.Time `json:"expired_date" db:"expired_date"`
	UserID       int       `json:"user_id" db:"user_id"`
}
