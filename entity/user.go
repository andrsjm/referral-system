package entity

type User struct {
	ID          int    `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	UserName    string `json:"username" db:"username"`
	Name        string `json:"name" db:"name"`
	ReferralHit int    `json:"referral_hit" db:"referral_hit"`
	UserType    int    `json:"user_type" db:"user_type"`
}
