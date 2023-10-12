package flow

import (
	"referral-system/entity"
	"referral-system/repository"
	"time"
)

type contributorFlow struct {
	repoContributor repository.IContributorRepository
	repoReferral    repository.IReferralRepository
	repoUser        repository.IUserRepository
}

func NewContributorFlow(contributorRepo repository.IContributorRepository, referralRepo repository.IReferralRepository, userRepo repository.IUserRepository) *contributorFlow {
	return &contributorFlow{
		repoContributor: contributorRepo,
		repoReferral:    referralRepo,
		repoUser:        userRepo,
	}
}

func (f *contributorFlow) Referral(contributor entity.Contributor) (isReady bool, expired bool, err error) {
	isReady, err = f.repoContributor.CheckContributor(contributor.Email)
	if err != nil {
		return
	} else if isReady {
		return
	}

	referral, err := f.repoReferral.GetReferral(contributor.ReferralCode)
	if err != nil {
		return
	}

	currentTime := time.Now()

	if currentTime.After(referral.ExpiredDate) || currentTime.Equal(referral.ExpiredDate) {
		return false, true, nil
	}

	err = f.repoContributor.InsertContributor(contributor)
	if err != nil {
		return
	}

	err = f.repoUser.UpdateReferralHit(referral.UserID)
	if err != nil {
		return
	}

	return
}
