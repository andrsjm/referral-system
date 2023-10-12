package flow

import (
	"fmt"
	"referral-system/entity"
	"referral-system/repository"
)

type userFlow struct {
	repoUser     repository.IUserRepository
	repoReferral repository.IReferralRepository
	flowReferral IReferralFlow
}

func NewUserFlow(userRepo repository.IUserRepository, referralRepo repository.IReferralRepository, flowReferral IReferralFlow) *userFlow {
	return &userFlow{
		repoUser:     userRepo,
		repoReferral: referralRepo,
		flowReferral: flowReferral,
	}
}

func (f *userFlow) Login(user entity.User) (entity.User, error) {
	users, err := f.repoUser.Login(user)

	return users, err
}

func (f *userFlow) Register(user entity.User) (referralLink string, err error) {
	res, err := f.repoUser.Register(user)
	if err != nil {
		return
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return
	}

	user.ID = int(lastInsertID)

	referralLink, err = f.GenerateNewCode(user)
	if err != nil {
		return
	}

	return
}

func (f *userFlow) Update(user entity.User) (err error) {
	err = f.repoUser.Update(user)

	return
}

func (f *userFlow) GenerateNewCode(user entity.User) (referralLink string, err error) {
	referral, err := f.flowReferral.GenerateReferralCode(user)
	if err != nil {
		return
	}

	err = f.repoReferral.InsertReferral(referral)
	if err != nil {
		return
	}

	referralLink = fmt.Sprintf("localhost:8000/referral/%s", referral.ReferralCode)

	return
}
