package mock

import (
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockContributorFlow struct {
	mock.Mock
}

func (m *MockContributorFlow) Referral(contributor entity.Contributor) (isReady bool, expired bool, err error) {
	args := m.Called(contributor)
	return args.Get(0).(bool), args.Get(1).(bool), args.Error(2)
}
