package repository

import (
	"database/sql"
	"referral-system/entity"
)

type IContributorRepository interface {
	CheckContributor(email string) (isReady bool, err error)
	InsertContributor(contributor entity.Contributor) (err error)
}

type IUserRepository interface {
	Login(user entity.User) (entity.User, error)
	Register(user entity.User) (res sql.Result, err error)
	Update(user entity.User) (err error)
	UpdateReferralHit(ID int) (err error)
}

type IReferralRepository interface {
	InsertReferral(referral entity.Referral) (err error)
	GetReferral(referralCode string) (referral entity.Referral, err error)
}
