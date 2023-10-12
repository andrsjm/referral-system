package mock

import (
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockReferralRepo struct {
	mock.Mock
}

func (m *MockReferralRepo) InsertReferral(referral entity.Referral) (err error) {
	args := m.Called(referral)
	return args.Error(0)
}

func (m *MockReferralRepo) GetReferral(referralCode string) (referral entity.Referral, err error) {
	args := m.Called(referralCode)
	return args.Get(0).(entity.Referral), args.Error(1)
}
