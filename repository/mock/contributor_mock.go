package mock

import (
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockContributorRepo struct {
	mock.Mock
}

func (m *MockContributorRepo) CheckContributor(email string) (isReady bool, err error) {
	args := m.Called(email)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockContributorRepo) InsertContributor(contributor entity.Contributor) (err error) {
	args := m.Called(contributor)
	return args.Error(0)
}
