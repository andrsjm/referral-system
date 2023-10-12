package mock

import (
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserFlow struct {
	mock.Mock
}

func (m *MockUserFlow) Login(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserFlow) Register(user entity.User) (referralString string, err error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserFlow) Update(user entity.User) (err error) {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserFlow) GenerateNewCode(user entity.User) (referralLink string, err error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
