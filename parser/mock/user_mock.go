package mock

import (
	"net/http"
	"referral-system/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserParser struct {
	mock.Mock
}

func (m *MockUserParser) ParseUserEntity(r *http.Request) (entity.User, error) {
	args := m.Called(r)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserParser) ParseBlogID(r *http.Request) int {
	args := m.Called(r)
	return args.Int(0)
}

func (m *MockUserParser) ParseUserFromCoockies(r *http.Request) entity.User {
	args := m.Called(r)
	return args.Get(0).(entity.User)
}
