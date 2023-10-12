package mock

import (
	"database/sql"
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Login(user entity.User) (users entity.User, err error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepo) Register(user entity.User) (res sql.Result, err error) {
	args := m.Called(user)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *MockUserRepo) Update(user entity.User) (err error) {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) UpdateReferralHit(ID int) (err error) {
	args := m.Called(ID)
	return args.Error(0)
}
