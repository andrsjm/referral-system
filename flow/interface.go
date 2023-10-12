package flow

import "referral-system/entity"

type IUserFlow interface {
	Login(user entity.User) (entity.User, error)
	Register(user entity.User) (referralCode string, err error)
	Update(user entity.User) (err error)
	GenerateNewCode(user entity.User) (referralLink string, err error)
}

type IContributorFlow interface {
	Referral(contributor entity.Contributor) (isReady bool, expired bool, err error)
}

type IReferralFlow interface {
	GenerateReferralCode(user entity.User) (referral entity.Referral, err error)
}
