package mock

import (
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockReferralFlow struct {
	mock.Mock
}

func (m *MockReferralFlow) GenerateReferralCode(user entity.User) (referral entity.Referral, err error) {
	args := m.Called(user)
	return args.Get(0).(entity.Referral), args.Error(1)
}
