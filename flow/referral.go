package flow

import (
	"referral-system/entity"
	"referral-system/repository"
	"referral-system/util"
	"time"
)

type referralFlow struct {
	repoReferral repository.IReferralRepository
}

func NewReferralFlow(referralRepo repository.IReferralRepository) *referralFlow {
	return &referralFlow{
		repoReferral: referralRepo,
	}
}

func (f *referralFlow) GenerateReferralCode(user entity.User) (referral entity.Referral, err error) {
	referralCode, err := util.GenerateCode(user.UserName)
	if err != nil {
		return
	}

	referral = entity.Referral{
		ReferralCode: referralCode,
		ExpiredDate:  time.Now().Add(7 * 24 * time.Hour),
		UserID:       int(user.ID),
	}

	return
}
