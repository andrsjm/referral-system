package entity

type Contributor struct {
	ID           int    `json:"id" db:"id"`
	ReferralCode string `json:"referral_code" db:"referral_code"`
	Email        string `json:"email" db:"email"`
}
